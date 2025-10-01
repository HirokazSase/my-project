# アーキテクチャ設計書

## 概要

本システムは、ドメイン駆動設計（DDD）とクリーンアーキテクチャの原則に基づいて設計された経費精算システムです。
保守性、拡張性、テスタビリティを重視し、ビジネスロジックとインフラストラクチャの関心事を明確に分離しています。

## アーキテクチャ構造

### レイヤー構成

```
┌─────────────────────────────────────────────────┐
│                   cmd/api                        │  ← エントリーポイント
│                (main.go)                         │
└─────────────────────────────────────────────────┘
                         │
┌─────────────────────────────────────────────────┐
│            Infrastructure Layer                  │  ← 外部インターフェース
│  • Web (HTTP API, Router, Handler)              │
│  • Persistence (Repository実装)                  │
└─────────────────────────────────────────────────┘
                         │
┌─────────────────────────────────────────────────┐
│             Application Layer                    │  ← アプリケーションサービス
│  • UseCase (ビジネスフロー調整)                 │
│  • DTO (データ転送オブジェクト)                  │
└─────────────────────────────────────────────────┘
                         │
┌─────────────────────────────────────────────────┐
│               Domain Layer                       │  ← ビジネスロジックの中核
│  • Entity (ビジネスオブジェクト)                │
│  • Value Object (不変の値オブジェクト)           │
│  • Repository Interface (データアクセス抽象化)   │
│  • Domain Service (ドメインサービス)             │
└─────────────────────────────────────────────────┘
```

### 依存関係の方向

```
cmd/api
   ↓
Infrastructure ────→ Application ────→ Domain
                        ↑                 ↑
                        └─────────────────┘
```

**重要な原則**: 依存関係は内側に向かってのみ流れます（依存関係逆転の原則）

## レイヤー詳細設計

### 1. Domain Layer（ドメイン層）

**責務**: ビジネスロジックとルールの実装

#### Entity
- **User**: ユーザーエンティティ
  - ユーザー情報の管理
  - プロフィール更新ロジック
  - バリデーションルール

- **Category**: カテゴリエンティティ
  - カテゴリ情報の管理
  - カテゴリ更新ロジック
  - 色設定とバリデーション

- **Expense**: 経費エンティティ
  - 経費情報の管理
  - ステータス遷移ロジック（draft → submitted → approved/rejected）
  - 編集可能性の判定
  - ビジネスルールの実装

#### Value Object
- **UserID, CategoryID, ExpenseID**: 各エンティティの識別子
  - UUID形式の検証
  - 等価性の実装
  - 不変性の保証

- **Money**: 金額を表現する値オブジェクト
  - 通貨を含む金額の表現
  - 計算メソッド（加算、減算、乗算）
  - 不正な値の防止

#### Repository Interface
- データアクセスの抽象化
- ドメイン層はインフラストラクチャ層の実装に依存しない
- テスタビリティの向上

#### ビジネスルール例
```go
// 経費の申請可能性チェック
func (e *Expense) CanSubmit() bool {
    return e.status == ExpenseStatusDraft
}

// 経費の編集可能性チェック  
func (e *Expense) CanEdit() bool {
    return e.status == ExpenseStatusDraft
}

// 金額の妥当性チェック
func NewMoney(amount float64, currency string) (*Money, error) {
    if amount < 0 {
        return nil, errors.NewDomainError(InvalidExpenseAmount, "金額は負の値にできません")
    }
    // ...
}
```

### 2. Application Layer（アプリケーション層）

**責務**: ドメインオブジェクトの協調とアプリケーションフローの調整

#### UseCase
- **UserUseCase**: ユーザー管理の調整
  - メールアドレス重複チェック
  - ユーザー作成・更新・削除フロー

- **CategoryUseCase**: カテゴリ管理の調整
  - カテゴリ名重複チェック
  - 使用中カテゴリの削除防止

- **ExpenseUseCase**: 経費管理の調整
  - 経費作成時の関連エンティティ存在確認
  - ステータス遷移の調整
  - 複雑な検索処理

#### DTO (Data Transfer Object)
- 外部インターフェースとのデータ交換
- バリデーション属性の定義
- レスポンス形式の統一

#### トランザクション境界
```go
func (uc *ExpenseUseCase) CreateExpense(ctx context.Context, userID string, req *dto.CreateExpenseRequest) (*dto.ExpenseResponse, error) {
    // 1. ユーザー存在確認
    user, err := uc.userRepo.FindByID(ctx, uid)
    
    // 2. カテゴリ存在確認
    category, err := uc.categoryRepo.FindByID(ctx, cid)
    
    // 3. ドメインオブジェクト作成
    expense, err := entity.NewExpense(...)
    
    // 4. 永続化
    if err := uc.expenseRepo.Save(ctx, expense); err != nil {
        return nil, err
    }
    
    return response, nil
}
```

### 3. Infrastructure Layer（インフラストラクチャ層）

**責務**: 外部システムとの連携、技術的詳細の実装

#### Web Sub-layer
- **Router**: HTTPルーティングの設定
- **Handler**: HTTPリクエスト・レスポンスの処理
- **Middleware**: CORS、ロギング、エラーハンドリング

#### Persistence Sub-layer
- **MemoryRepository**: メモリ内データストア実装
  - 開発・テスト用の軽量実装
  - 本番環境では DB Repository に置き換え可能
  - Repository Interface の実装

#### エラーハンドリング戦略
```go
func handleError(c *gin.Context, err error) {
    switch e := err.(type) {
    case *errors.DomainError:
        handleDomainError(c, e)
    case *errors.ApplicationError:
        handleApplicationError(c, e)
    default:
        c.JSON(http.StatusInternalServerError, ErrorResponse{...})
    }
}
```

### 4. Entry Point（エントリーポイント）

**責務**: アプリケーションの初期化と起動

#### 依存関係の注入
```go
func main() {
    // Repository層の初期化
    userRepo := persistence.NewMemoryUserRepository()
    categoryRepo := persistence.NewMemoryCategoryRepository()
    expenseRepo := persistence.NewMemoryExpenseRepository()

    // UseCase層の初期化
    userUseCase := usecase.NewUserUseCase(userRepo)
    categoryUseCase := usecase.NewCategoryUseCase(categoryRepo, expenseRepo)
    expenseUseCase := usecase.NewExpenseUseCase(expenseRepo, userRepo, categoryRepo)

    // Handler層の初期化
    userHandler := handler.NewUserHandler(userUseCase)
    categoryHandler := handler.NewCategoryHandler(categoryUseCase)
    expenseHandler := handler.NewExpenseHandler(expenseUseCase)

    // Router設定とサーバー起動
    router := web.SetupRouter(userHandler, categoryHandler, expenseHandler)
    srv := &http.Server{Addr: ":8080", Handler: router}
    srv.ListenAndServe()
}
```

## 設計パターンと原則

### 1. ドメイン駆動設計（DDD）

#### Entity Pattern
```go
type Expense struct {
    id          *valueobject.ExpenseID
    userID      *valueobject.UserID
    categoryID  *valueobject.CategoryID
    amount      *valueobject.Money
    title       string
    status      ExpenseStatus
    // ...
}

func (e *Expense) Submit() error {
    if e.status != ExpenseStatusDraft {
        return errors.NewDomainError("EXPENSE_SUBMIT_NOT_ALLOWED", "...")
    }
    e.status = ExpenseStatusSubmitted
    e.updatedAt = time.Now()
    return nil
}
```

#### Value Object Pattern
```go
type Money struct {
    amount   float64
    currency string
}

func (m *Money) Add(other *Money) (*Money, error) {
    if m.currency != other.currency {
        return nil, errors.NewDomainError("INVALID_CURRENCY", "...")
    }
    return NewMoney(m.amount+other.amount, m.currency)
}
```

#### Repository Pattern
```go
type ExpenseRepository interface {
    Save(ctx context.Context, expense *entity.Expense) error
    FindByID(ctx context.Context, id *valueobject.ExpenseID) (*entity.Expense, error)
    FindByUserID(ctx context.Context, userID *valueobject.UserID) ([]*entity.Expense, error)
    // ...
}
```

### 2. クリーンアーキテクチャ原則

#### 依存関係逆転の原則
- 内側のレイヤーは外側のレイヤーを知らない
- インターフェースによる抽象化
- 実装の詳細は外側に追いやる

#### 単一責任の原則
- 各レイヤーは明確な責務を持つ
- 変更の理由は一つ

#### オープン・クローズドの原則
- 拡張に対して開いている
- 修正に対して閉じている

### 3. SOLID原則の適用

#### Interface Segregation Principle
```go
// 肥大化したインターフェースを避ける
type UserReader interface {
    FindByID(ctx context.Context, id *valueobject.UserID) (*entity.User, error)
    FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type UserWriter interface {
    Save(ctx context.Context, user *entity.User) error
    Update(ctx context.Context, user *entity.User) error
}

// 必要に応じて組み合わせ
type UserRepository interface {
    UserReader
    UserWriter
}
```

## テスト戦略

### 1. 単体テスト

#### Domain Layer
```go
func TestExpense_Submit(t *testing.T) {
    expense, _ := entity.NewExpense(...)
    
    err := expense.Submit()
    
    assert.NoError(t, err)
    assert.Equal(t, entity.ExpenseStatusSubmitted, expense.Status())
}
```

#### Application Layer
```go
func TestExpenseUseCase_CreateExpense(t *testing.T) {
    // Repository のモックを使用
    userRepo := &mockUserRepository{}
    categoryRepo := &mockCategoryRepository{}
    expenseRepo := &mockExpenseRepository{}
    
    useCase := usecase.NewExpenseUseCase(expenseRepo, userRepo, categoryRepo)
    
    result, err := useCase.CreateExpense(ctx, userID, req)
    
    assert.NoError(t, err)
    assert.NotNil(t, result)
}
```

### 2. 統合テスト

#### End-to-End API Test
```go
func TestExpenseWorkflow(t *testing.T) {
    server := setupTestServer()
    
    // 1. ユーザー作成
    user := createTestUser(server)
    
    // 2. カテゴリ作成
    category := createTestCategory(server)
    
    // 3. 経費作成
    expense := createTestExpense(server, user.ID, category.ID)
    
    // 4. 経費申請
    submittedExpense := submitExpense(server, expense.ID)
    
    // 5. 経費承認
    approvedExpense := approveExpense(server, expense.ID)
    
    assert.Equal(t, "approved", approvedExpense.Status)
}
```

## 拡張性と保守性

### 1. 新機能の追加

#### レシート画像アップロード機能の追加例
```go
// 1. Domain Layer: Value Object追加
type AttachmentID struct {
    value string
}

// 2. Domain Layer: Entity拡張
type Expense struct {
    // 既存フィールド...
    attachments []*valueobject.AttachmentID
}

func (e *Expense) AddAttachment(id *valueobject.AttachmentID) error {
    // ビジネスルール実装
}

// 3. Application Layer: UseCase拡張
func (uc *ExpenseUseCase) UploadReceipt(ctx context.Context, expenseID string, file []byte) error {
    // ファイル処理とドメインオブジェクト更新
}

// 4. Infrastructure Layer: Handler追加
func (h *ExpenseHandler) UploadReceipt(c *gin.Context) {
    // HTTPファイルアップロード処理
}
```

### 2. データストアの変更

#### PostgreSQLへの移行例
```go
// 新しいRepository実装を追加
type PostgreSQLUserRepository struct {
    db *sql.DB
}

func (r *PostgreSQLUserRepository) FindByID(ctx context.Context, id *valueobject.UserID) (*entity.User, error) {
    // PostgreSQL実装
}

// main.goで注入先を変更するだけ
func main() {
    // userRepo := persistence.NewMemoryUserRepository()  // 旧実装
    userRepo := persistence.NewPostgreSQLUserRepository(db)  // 新実装
    
    // 以降は変更なし
    userUseCase := usecase.NewUserUseCase(userRepo)
    // ...
}
```

### 3. 新しいUI層の追加

#### GraphQL APIの追加例
```go
// GraphQLハンドラーを追加
type GraphQLHandler struct {
    userUseCase     *usecase.UserUseCase
    categoryUseCase *usecase.CategoryUseCase
    expenseUseCase  *usecase.ExpenseUseCase
}

// 既存のUseCaseを再利用
func (h *GraphQLHandler) GetUser(ctx context.Context, id string) (*dto.UserResponse, error) {
    return h.userUseCase.GetUser(ctx, id)
}
```

## パフォーマンス考慮事項

### 1. N+1問題の回避

```go
// 悪い例: N+1問題発生
func (uc *ExpenseUseCase) GetExpensesByUser(ctx context.Context, userID string) ([]*dto.ExpenseResponse, error) {
    expenses, _ := uc.expenseRepo.FindByUserID(ctx, uid)
    
    for _, expense := range expenses {
        // 各経費に対してカテゴリを個別取得（N+1問題）
        category, _ := uc.categoryRepo.FindByID(ctx, expense.CategoryID())
    }
}

// 良い例: バッチ取得
func (uc *ExpenseUseCase) GetExpensesByUser(ctx context.Context, userID string) ([]*dto.ExpenseResponse, error) {
    expenses, _ := uc.expenseRepo.FindByUserID(ctx, uid)
    
    categoryIDs := extractCategoryIDs(expenses)
    categories, _ := uc.categoryRepo.FindByIDs(ctx, categoryIDs) // バッチ取得
    
    return buildResponse(expenses, categories), nil
}
```

### 2. キャッシュ戦略

```go
// Repository層でキャッシュ実装
type CachedCategoryRepository struct {
    repo  repository.CategoryRepository
    cache map[string]*entity.Category
    mu    sync.RWMutex
}

func (r *CachedCategoryRepository) FindByID(ctx context.Context, id *valueobject.CategoryID) (*entity.Category, error) {
    r.mu.RLock()
    if cached, exists := r.cache[id.String()]; exists {
        r.mu.RUnlock()
        return cached, nil
    }
    r.mu.RUnlock()
    
    category, err := r.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    r.mu.Lock()
    r.cache[id.String()] = category
    r.mu.Unlock()
    
    return category, nil
}
```

## セキュリティ考慮事項

### 1. 入力値検証

```go
// ドメイン層でのバリデーション
func validateExpenseAmount(amount float64) error {
    if amount < 0 {
        return errors.NewDomainError(InvalidExpenseAmount, "金額は負の値にできません")
    }
    if amount > math.MaxFloat64/100 {
        return errors.NewDomainError(InvalidExpenseAmount, "金額が大きすぎます")
    }
    return nil
}

// アプリケーション層でのバリデーション
type CreateExpenseRequest struct {
    Amount      float64   `json:"amount" binding:"required,min=0"`
    Title       string    `json:"title" binding:"required,max=100"`
    Description string    `json:"description" binding:"max=500"`
}
```

### 2. 認証・認可の追加準備

```go
// 将来的な認証機能の追加例
type AuthContext struct {
    UserID *valueobject.UserID
    Roles  []string
}

func (uc *ExpenseUseCase) GetExpense(ctx context.Context, auth *AuthContext, expenseID string) (*dto.ExpenseResponse, error) {
    expense, err := uc.expenseRepo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    // 認可チェック
    if !auth.UserID.Equals(expense.UserID()) && !auth.HasRole("admin") {
        return nil, errors.NewApplicationError("FORBIDDEN", "アクセス権限がありません")
    }
    
    return response, nil
}
```

## まとめ

本アーキテクチャは以下の特徴を持ちます：

1. **明確な責務分離**: 各レイヤーが明確な役割を持つ
2. **高い保守性**: 変更の影響範囲が限定される
3. **優れたテスタビリティ**: 各層を独立してテスト可能
4. **拡張性**: 新機能追加やインフラ変更が容易
5. **ビジネスロジックの保護**: ドメイン層が技術的詳細から独立

このアーキテクチャにより、長期的に保守可能で拡張性の高いシステムを構築できています。
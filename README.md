# 経費精算システム (Expense Management System)

転職用ポートフォリオとして作成した、Go言語によるシンプルな経費精算システムです。  
ドメイン駆動設計（DDD）とクリーンアーキテクチャの考え方を採用し、保守性と拡張性を重視して設計されています。

## 🏗️ アーキテクチャ

本システムは以下のレイヤー構造で設計されています：

```
cmd/                     # アプリケーションエントリーポイント
├── api/                 # REST APIサーバー
internal/
├── domain/              # ドメイン層
│   ├── entity/          # エンティティ
│   ├── valueobject/     # 値オブジェクト
│   ├── repository/      # リポジトリインターフェース
│   └── service/         # ドメインサービス
├── application/         # アプリケーション層
│   ├── usecase/         # ユースケース
│   └── dto/             # データ転送オブジェクト
└── infrastructure/      # インフラストラクチャ層
    ├── persistence/     # データ永続化
    └── web/             # Web層（HTTP API）
pkg/
└── errors/              # 共通エラー定義
test/                    # 統合テスト
```

## 🚀 機能

### 1. ユーザー管理
- ユーザー登録・更新・削除・取得
- メールアドレスの重複チェック

### 2. カテゴリ管理  
- 経費カテゴリの作成・更新・削除・取得
- カテゴリ名の重複チェック
- カラーコード設定

### 3. 経費管理
- 経費の作成・更新・削除・取得
- 経費ステータス管理（下書き→申請→承認/却下）
- ユーザー別・ステータス別経費一覧取得
- 日付範囲での絞り込み

### 4. 経費ワークフロー
- 下書き状態での編集
- 申請による承認フローへの移行
- 承認者による承認・却下

## 🛠️ 技術スタック

- **言語**: Go 1.21
- **Webフレームワーク**: Gin
- **テスト**: testify
- **UUID生成**: google/uuid
- **アーキテクチャ**: DDD + Clean Architecture
- **データ永続化**: メモリストア（ポートフォリオ用）

## 📋 API仕様

### ユーザー API

| メソッド | エンドポイント | 説明 |
|---------|---------------|------|
| POST    | /api/v1/users | ユーザー作成 |
| GET     | /api/v1/users | 全ユーザー取得 |
| GET     | /api/v1/users/{id} | ユーザー取得 |
| PUT     | /api/v1/users/{id} | ユーザー更新 |
| DELETE  | /api/v1/users/{id} | ユーザー削除 |

### カテゴリ API

| メソッド | エンドポイント | 説明 |
|---------|---------------|------|
| POST    | /api/v1/categories | カテゴリ作成 |
| GET     | /api/v1/categories | 全カテゴリ取得 |
| GET     | /api/v1/categories/{id} | カテゴリ取得 |
| PUT     | /api/v1/categories/{id} | カテゴリ更新 |
| DELETE  | /api/v1/categories/{id} | カテゴリ削除 |

### 経費 API

| メソッド | エンドポイント | 説明 |
|---------|---------------|------|
| POST    | /api/v1/users/{user_id}/expenses | 経費作成 |
| GET     | /api/v1/users/{user_id}/expenses | ユーザー経費一覧取得 |
| GET     | /api/v1/expenses/{id} | 経費取得 |
| PUT     | /api/v1/expenses/{id} | 経費更新 |
| DELETE  | /api/v1/expenses/{id} | 経費削除 |
| POST    | /api/v1/expenses/{id}/submit | 経費申請 |
| POST    | /api/v1/expenses/{id}/approve | 経費承認 |
| POST    | /api/v1/expenses/{id}/reject | 経費却下 |

### ヘルスチェック

| メソッド | エンドポイント | 説明 |
|---------|---------------|------|
| GET     | /health | ヘルスチェック |

## 🔧 セットアップ・実行方法

### 1. 環境準備

```bash
# Go 1.21+ が必要です
go version

# プロジェクトをクローン
git clone <repository-url>
cd expense-management-system
```

### 2. 依存関係のインストール

```bash
go mod tidy
```

### 3. アプリケーションのビルド

```bash
go build -o bin/expense-api ./cmd/api
```

### 4. サーバー起動

```bash
# 直接実行
go run ./cmd/api

# または、ビルドしたバイナリを実行
./bin/expense-api
```

サーバーはデフォルトで `http://localhost:8080` で起動します。

### 5. ヘルスチェック

```bash
curl http://localhost:8080/health
```

## 🧪 テスト実行

### 単体テスト

```bash
# 全テスト実行
go test ./...

# カバレッジ付きテスト
go test ./... -coverprofile=coverage.out

# カバレッジレポート表示
go tool cover -html=coverage.out
```

### 特定パッケージのテスト

```bash
# ドメインレイヤーのテスト
go test ./internal/domain/... -v

# アプリケーションレイヤーのテスト
go test ./internal/application/... -v

# 統合テスト
go test ./test -v
```

## 📖 使用例

### 1. ユーザー作成

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "田中太郎",
    "email": "tanaka@example.com"
  }'
```

### 2. カテゴリ作成

```bash
curl -X POST http://localhost:8080/api/v1/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "交通費",
    "description": "電車・バス・タクシーなどの交通費",
    "color": "#FF6B6B"
  }'
```

### 3. 経費作成

```bash
curl -X POST http://localhost:8080/api/v1/users/{user_id}/expenses \
  -H "Content-Type: application/json" \
  -d '{
    "category_id": "{category_id}",
    "amount": 1500,
    "currency": "JPY",
    "title": "渋谷駅からオフィス",
    "description": "営業訪問のための交通費",
    "date": "2023-10-01T09:00:00Z"
  }'
```

### 4. 経費申請

```bash
curl -X POST http://localhost:8080/api/v1/expenses/{expense_id}/submit
```

## 🔍 設計のポイント

### 1. ドメイン駆動設計（DDD）

- **エンティティ**: ビジネスの中核となる識別可能なオブジェクト
- **値オブジェクト**: 不変で等価性を持つオブジェクト（Money, UserIDなど）
- **ドメインサービス**: エンティティに属さないビジネスロジック
- **リポジトリパターン**: データ永続化の抽象化

### 2. クリーンアーキテクチャ

- **依存関係の逆転**: 外側の層が内側の層に依存
- **レイヤー分離**: ビジネスロジックと技術的詳細の分離
- **インターフェース活用**: 実装の詳細を隠蔽

### 3. エラーハンドリング

- **カスタムエラー型**: ドメインエラーとアプリケーションエラーの分離
- **エラーコード**: 構造化されたエラー情報
- **適切なHTTPステータス**: RESTful なエラーレスポンス

### 4. テスト戦略

- **単体テスト**: 各コンポーネントの独立したテスト
- **統合テスト**: APIエンドポイントの動作確認
- **テストダブル**: モックを使った依存関係の分離

## 🏆 ポートフォリオとしての特徴

### 1. 実践的な設計パターン
- DDD（ドメイン駆動設計）の実践
- クリーンアーキテクチャによる保守性の向上
- SOLID原則の適用

### 2. 高品質なコード
- 包括的なテストスイート
- 適切なエラーハンドリング
- 読みやすく保守しやすいコード構造

### 3. 現実的なビジネス要件
- 経費精算という実在するビジネス要件
- 承認ワークフローの実装
- バリデーションとビジネスルールの実装

### 4. 拡張性
- 新機能追加が容易な構造
- データベース実装の切り替えが容易
- 新しいUIフロントエンドの追加が容易

## 📝 今後の拡張可能性

- データベース（PostgreSQL、MySQL等）への移行
- 認証・認可機能の追加
- ファイルアップロード機能（レシート画像等）
- 管理者機能の追加
- フロントエンド（React、Vue.js等）の追加
- Docker化とK8s対応
- CI/CDパイプラインの構築

## 📄 ライセンス

MIT License

## 👤 作成者

転職用ポートフォリオプロジェクト

---

このプロジェクトは、Go言語によるクリーンアーキテクチャとDDDの実践例として作成されました。  
実際の業務で使用できるレベルの設計と実装を心がけています。
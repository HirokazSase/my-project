# 経費管理システム - React Frontend

## 🎯 プロジェクト概要

**DDD（ドメイン駆動設計）** と **MVVM（Model-View-ViewModel）** パターンを採用したモダンなReact TypeScript経費管理フロントエンドアプリケーションです。

### ✨ 特徴

- 🏗️ **Clean Architecture**: 層分離によるメンテナブルな設計
- 🎯 **DDD実装**: ドメインモデル中心の設計
- 📱 **MVVM パターン**: ViewModelによるビジネスロジック分離
- 🔒 **TypeScript**: 型安全なコード
- 🎨 **レスポンシブUI**: モバイル対応デザイン
- 🧪 **テストカバレッジ**: 包括的なユニットテスト

## 🚀 クイックスタート

### 前提条件

- Node.js 16.x 以上
- npm 7.x 以上
- Go Backend API (port 8080)

### インストール

```bash
# 依存関係のインストール
npm install

# 開発サーバー起動
npm start

# テスト実行
npm test

# プロダクションビルド
npm run build
```

### 環境設定

```bash
# .env.local ファイルを作成（オプション）
REACT_APP_API_URL=http://localhost:8080
```

## 📁 プロジェクト構造

```
frontend/
├── 📂 public/                     # 静的ファイル
│   ├── index.html                 # エントリーポイント
│   └── manifest.json             # PWA設定
├── 📂 src/
│   ├── 📂 domain/                 # ドメイン層
│   │   ├── 📂 models/             # ドメインモデル
│   │   │   ├── User.ts           # ユーザーエンティティ
│   │   │   ├── Category.ts       # カテゴリエンティティ  
│   │   │   └── Expense.ts        # 経費エンティティ
│   │   └── 📂 repositories/       # リポジトリインターフェース
│   │       ├── UserRepository.ts
│   │       ├── CategoryRepository.ts
│   │       └── ExpenseRepository.ts
│   ├── 📂 application/            # アプリケーション層
│   │   └── 📂 viewModels/         # ViewModelクラス
│   │       ├── UserViewModel.ts   # ユーザー操作ロジック
│   │       ├── CategoryViewModel.ts # カテゴリ操作ロジック
│   │       ├── ExpenseViewModel.ts # 経費操作ロジック
│   │       └── 📂 __tests__/      # ViewModelテスト
│   ├── 📂 infrastructure/         # インフラストラクチャ層
│   │   └── 📂 api/               # API通信層
│   │       ├── ApiClient.ts      # HTTPクライアント
│   │       ├── UserApiService.ts # ユーザーAPI
│   │       ├── CategoryApiService.ts # カテゴリAPI
│   │       └── ExpenseApiService.ts # 経費API
│   ├── 📂 presentation/           # プレゼンテーション層
│   │   ├── 📂 components/         # Reactコンポーネント
│   │   │   ├── 📂 Common/         # 共通コンポーネント
│   │   │   │   ├── LoadingSpinner.tsx
│   │   │   │   └── ErrorMessage.tsx
│   │   │   ├── 📂 Layout/         # レイアウトコンポーネント
│   │   │   │   ├── Header.tsx
│   │   │   │   └── Layout.tsx
│   │   │   ├── 📂 User/           # ユーザー管理
│   │   │   │   ├── UserList.tsx
│   │   │   │   └── UserForm.tsx
│   │   │   ├── 📂 Category/       # カテゴリ管理
│   │   │   │   ├── CategoryList.tsx
│   │   │   │   └── CategoryForm.tsx
│   │   │   └── 📂 Expense/        # 経費管理
│   │   │       ├── ExpenseList.tsx
│   │   │       ├── ExpenseCard.tsx
│   │   │       └── ExpenseForm.tsx
│   │   └── 📂 pages/              # ページコンポーネント
│   │       └── ExpensePage.tsx
│   ├── App.tsx                    # メインアプリケーション
│   ├── App.css                   # グローバルスタイル
│   └── index.tsx                 # ReactDOMエントリーポイント
├── package.json                   # 依存関係設定
├── tsconfig.json                 # TypeScript設定
└── README.md                     # このファイル
```

## 🎨 画面構成・レイアウト

### 📊 メインダッシュボード

```
┌─────────────────────────────────────────────────────┐
│                    🏢 経費管理システム                  │
├─────────────────────────────────────────────────────┤
│ [経費管理] [カテゴリ] [ユーザー]     👤 田中太郎 ▼    │
├─────────────────────────────────────────────────────┤
│                                                     │
│  📋 経費管理                      [+ 新しい経費]      │
│                                                     │
│  [承認待ち] [承認済み] [却下済み] ◄── ステータスタブ     │
│                                                     │
│  📊 サマリー                                         │
│  ├ 総件数: 15件                                      │
│  └ 総金額: ¥125,400                                  │
│                                                     │
│  📅 2024年10月 (¥85,200 / 8件)                      │
│  ┌─────────────────────────────────────────────┐  │
│  │ 🍽️ 昼食代              ¥1,200   2024/10/01  │  │
│  │    営業先での昼食代              [承認][却下]     │  │
│  │                         [編集][削除]         │  │
│  └─────────────────────────────────────────────┘  │
│  ┌─────────────────────────────────────────────┐  │
│  │ 🚗 交通費              ¥580     2024/10/01  │  │  
│  │    新宿→渋谷 電車代              [承認][却下]     │  │
│  │                         [編集][削除]         │  │
│  └─────────────────────────────────────────────┘  │
│                                                     │
└─────────────────────────────────────────────────────┘
```

### 🏷️ カテゴリ管理画面

```
┌─────────────────────────────────────────────────────┐
│  📁 カテゴリ管理                  [+ 新しいカテゴリ]    │
│                                                     │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐ │
│  │ 🔵 交通費    │  │ 🔴 食費      │  │ 🟢 事務用品   │ │
│  │             │  │             │  │             │ │
│  │ 電車、バス、  │  │ 昼食、会議時  │  │ 文房具、PC   │ │
│  │ タクシー代   │  │ の食事代     │  │ 用品等      │ │
│  │             │  │             │  │             │ │
│  │ 2024/01/15  │  │ 2024/01/15  │  │ 2024/01/15  │ │
│  │ [編集][削除]  │  │ [編集][削除]  │  │ [編集][削除]  │ │
│  └─────────────┘  └─────────────┘  └─────────────┘ │
│                                                     │
└─────────────────────────────────────────────────────┘
```

### 👥 ユーザー管理画面

```
┌─────────────────────────────────────────────────────┐
│  👥 ユーザー管理                   [+ 新しいユーザー]   │
│                                                     │
│  ┌─────────────────────────────────────────────┐    │
│  │ 👤 田中太郎 ★                               │    │
│  │    tanaka@example.com            現在のユーザー   │    │
│  │                                             │    │
│  │    作成日: 2024/01/15                        │    │
│  │                           [編集] [削除]        │    │
│  └─────────────────────────────────────────────┘    │
│                                                     │
│  ┌─────────────────────────────────────────────┐    │
│  │ 👤 佐藤花子                                  │    │
│  │    sato@example.com                         │    │
│  │                                             │    │
│  │    作成日: 2024/01/16                        │    │
│  │                    [選択] [編集] [削除]        │    │
│  └─────────────────────────────────────────────┘    │
│                                                     │
└─────────────────────────────────────────────────────┘
```

## 🏗️ アーキテクチャ詳細

### DDD (Domain-Driven Design) 実装

#### 1. ドメイン層 (Domain Layer)

```typescript
// User Entity (ドメインエンティティ)
export class User {
  constructor(
    public readonly id: string,
    public readonly name: string,
    public readonly email: string,
    public readonly createdAt: Date,
    public readonly updatedAt: Date
  ) {
    this.validate(); // ドメインルール検証
  }
  
  private validate(): void {
    // ビジネスルールをここで実装
  }
}

// Repository Interface (ドメインサービス)
export interface UserRepository {
  getAll(): Promise<User[]>;
  getById(id: string): Promise<User | null>;
  create(data: CreateUserData): Promise<User>;
  update(id: string, data: UpdateUserData): Promise<User>;
  delete(id: string): Promise<void>;
}
```

#### 2. アプリケーション層 (Application Layer)

```typescript
// ViewModel (アプリケーションサービス)
export class UserViewModel {
  constructor(private userRepository: UserRepository) {}
  
  async createUser(name: string, email: string): Promise<User> {
    // バリデーション + ビジネスロジック
    const errors = this.validateUser(name, email);
    if (errors.length > 0) {
      throw new Error(errors.join(', '));
    }
    
    return await this.userRepository.create({ name, email });
  }
}
```

#### 3. インフラストラクチャ層 (Infrastructure Layer)

```typescript
// API Service (外部サービス実装)
export class UserApiService implements UserRepository {
  constructor(private apiClient: ApiClient) {}
  
  async getAll(): Promise<User[]> {
    const response = await this.apiClient.get<UserApiData[]>('/users');
    return response.data?.map(data => this.mapToUser(data)) || [];
  }
}
```

#### 4. プレゼンテーション層 (Presentation Layer)

```typescript
// React Component
export const UserList: React.FC<UserListProps> = ({
  users, onUserEdit, onUserDelete 
}) => {
  // UIロジックのみ。ビジネスロジックはViewModelに委譲
  return (
    <div>
      {users.map(user => (
        <UserCard 
          key={user.id} 
          user={user}
          onEdit={() => onUserEdit(user)}
          onDelete={() => onUserDelete(user.id)}
        />
      ))}
    </div>
  );
};
```

### MVVM パターン実装

```
┌─────────────┐    ┌──────────────┐    ┌─────────────┐
│    View     │◄──►│  ViewModel   │◄──►│    Model    │
│ (React UI)  │    │ (Business    │    │ (Domain     │
│             │    │  Logic)      │    │  Entities)  │
└─────────────┘    └──────────────┘    └─────────────┘
       ▲                   ▲                   ▲
       │                   │                   │
   UI Events          State Management    Data Access
   User Input         Validation         Repository
   Display Logic      Error Handling     API Calls
```

## 🎯 機能一覧

### 👥 ユーザー管理

- ✅ ユーザー一覧表示
- ✅ 新規ユーザー作成
- ✅ ユーザー情報編集
- ✅ ユーザー削除
- ✅ 現在ユーザー切り替え
- ✅ メールアドレス重複チェック

### 🏷️ カテゴリ管理

- ✅ カテゴリ一覧表示（グリッドレイアウト）
- ✅ 新規カテゴリ作成
- ✅ カテゴリ編集（名前・説明・カラー）
- ✅ カテゴリ削除
- ✅ カラーピッカー（10色プリセット）
- ✅ バリデーション（文字数制限等）

### 💰 経費管理

- ✅ 経費一覧表示（月別グループ化）
- ✅ ステータス別フィルター（承認待ち・承認済み・却下済み）
- ✅ 新規経費登録
- ✅ 経費編集
- ✅ 経費削除
- ✅ 承認・却下処理
- ✅ 金額・件数サマリー表示
- ✅ 日付・カテゴリ別ソート

### 🔧 共通機能

- ✅ レスポンシブデザイン
- ✅ ローディング状態表示
- ✅ エラーハンドリング
- ✅ 確認ダイアログ
- ✅ フォームバリデーション
- ✅ 型安全なAPI通信

## 🧪 テスト戦略

### テスト種別

```bash
# 全テスト実行
npm test

# カバレッジレポート生成
npm test -- --coverage

# 特定テスト実行
npm test UserViewModel.test.ts

# ウォッチモード
npm test -- --watch
```

### テストファイル構成

```
src/
├── domain/models/__tests__/
│   ├── User.test.ts           # ドメインモデルテスト
│   ├── Category.test.ts
│   └── Expense.test.ts
├── application/viewModels/__tests__/
│   ├── UserViewModel.test.ts  # ViewModelテスト
│   ├── CategoryViewModel.test.ts
│   └── ExpenseViewModel.test.ts
├── infrastructure/api/__tests__/
│   └── ExpenseApiService.test.ts # APIサービステスト
└── presentation/components/**/__tests__/
    ├── ExpenseCard.test.tsx   # Componentテスト
    ├── ExpenseForm.test.tsx
    ├── CategoryList.test.tsx
    └── CategoryForm.test.tsx
```

### テストカバレッジ目標

- 🎯 **ドメインモデル**: 100%
- 🎯 **ViewModel**: 95%以上  
- 🎯 **APIサービス**: 90%以上
- 🎯 **Reactコンポーネント**: 85%以上

## 🚀 デプロイ・ビルド

### 開発環境

```bash
# 開発サーバー起動
npm start
# → http://localhost:3000

# バックエンド接続確認
curl http://localhost:8080/health
```

### プロダクションビルド

```bash
# 最適化ビルド実行
npm run build

# ビルド成果物確認
ls -la build/

# 静的サーバーでテスト
npx serve -s build
```

### 環境変数

```bash
# .env.local
REACT_APP_API_URL=https://api.expense-app.com
REACT_APP_VERSION=$npm_package_version
GENERATE_SOURCEMAP=false
```

## 🔧 開発ツール・設定

### TypeScript設定 (tsconfig.json)

```json
{
  "compilerOptions": {
    "target": "es5",
    "lib": ["dom", "dom.iterable", "es6"],
    "allowJs": true,
    "skipLibCheck": true,
    "esModuleInterop": true,
    "allowSyntheticDefaultImports": true,
    "strict": true,
    "forceConsistentCasingInFileNames": true,
    "noFallthroughCasesInSwitch": true,
    "module": "esnext",
    "moduleResolution": "node",
    "resolveJsonModule": true,
    "isolatedModules": true,
    "noEmit": true,
    "jsx": "react-jsx"
  },
  "include": ["src"]
}
```

### ESLint設定

```json
{
  "extends": [
    "react-app",
    "react-app/jest"
  ]
}
```

## 📊 パフォーマンス最適化

### バンドルサイズ最適化

```bash
# バンドル分析
npm run build
npx webpack-bundle-analyzer build/static/js/*.js
```

### コード分割

```typescript
// 遅延ローディング実装例
const ExpensePage = React.lazy(() => import('./presentation/pages/ExpensePage'));

function App() {
  return (
    <Suspense fallback={<LoadingSpinner />}>
      <ExpensePage />
    </Suspense>
  );
}
```

## 🐛 トラブルシューティング

### よくある問題

#### 1. API接続エラー

```bash
# バックエンドサーバー確認
curl http://localhost:8080/health

# プロキシ設定確認
# package.json の "proxy": "http://localhost:8080" を確認
```

#### 2. TypeScriptコンパイルエラー

```bash
# 型チェック実行
npx tsc --noEmit

# キャッシュクリア
rm -rf node_modules/.cache
npm start
```

#### 3. テスト実行エラー  

```bash
# Jest キャッシュクリア
npm test -- --clearCache

# テスト環境セットアップ確認
cat src/setupTests.ts
```

## 🔗 関連リンク

- 📚 **React公式ドキュメント**: https://react.dev
- 📘 **TypeScript Handbook**: https://www.typescriptlang.org/docs/
- 🧪 **Jest Testing Framework**: https://jestjs.io/docs/getting-started
- 🎨 **Create React App**: https://create-react-app.dev

## 🤝 開発チーム・コントリビューション

### 開発者

- **Frontend Developer**: AI Assistant
- **Architecture**: DDD + MVVM + Clean Architecture
- **Tech Lead**: Advanced React + TypeScript Implementation

### コントリビューション方法

1. 🍴 フォーク
2. 🌿 ブランチ作成 (`git checkout -b feature/amazing-feature`)
3. 💾 コミット (`git commit -m 'Add amazing feature'`)
4. 📤 プッシュ (`git push origin feature/amazing-feature`)
5. 🔄 プルリクエスト作成

### コーディング規約

- **命名規則**: camelCase (変数・関数), PascalCase (コンポーネント・クラス)
- **ファイル名**: PascalCase (コンポーネント), camelCase (ユーティリティ)
- **インポート順序**: 外部ライブラリ → 内部モジュール → 相対パス
- **コメント**: JSDoc形式で関数・クラスの説明を記載

## 📄 ライセンス

このプロジェクトはMITライセンスの下で公開されています。

---

## 🎉 まとめ

この経費管理システムフロントエンドは、**モダンなReact開発のベストプラクティス**を実装したポートフォリオ品質のアプリケーションです。

### 🎯 学習ポイント

- **DDD**: ドメイン中心設計の実践
- **MVVM**: 責務分離によるメンテナブルなコード  
- **Clean Architecture**: 依存関係の制御と柔軟性
- **TypeScript**: 型安全性による品質向上
- **Testing**: 包括的なテスト戦略
- **UI/UX**: ユーザー中心のインターフェース設計

**🚀 就職活動での活用価値**: エンタープライズレベルの設計パターンと実装技術を証明できる実用的なプロジェクトです。
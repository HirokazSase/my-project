# API仕様書

## 概要

経費精算システムのREST API仕様です。

## 基本情報

- **Base URL**: `http://localhost:8080/api/v1`
- **Content-Type**: `application/json`
- **レスポンス形式**: JSON

## エラーレスポンス

全てのエンドポイントで共通のエラーレスポンス形式を使用します：

```json
{
  "error": "ERROR_CODE",
  "message": "エラーメッセージ",
  "details": "詳細情報（オプション）"
}
```

## エンドポイント

### ヘルスチェック

#### GET /health

サーバーの稼働状況を確認します。

**レスポンス**
```json
{
  "status": "ok",
  "message": "Expense Management System is running"
}
```

---

## ユーザー管理

### POST /users

新しいユーザーを作成します。

**リクエストボディ**
```json
{
  "name": "田中太郎",
  "email": "tanaka@example.com"
}
```

**レスポンス（201 Created）**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "田中太郎",
  "email": "tanaka@example.com",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T10:00:00Z"
}
```

**エラー**
- `400 Bad Request`: バリデーションエラー
- `409 Conflict`: メールアドレス重複

### GET /users

全てのユーザーを取得します。

**レスポンス（200 OK）**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "田中太郎",
    "email": "tanaka@example.com",
    "created_at": "2023-10-01T10:00:00Z",
    "updated_at": "2023-10-01T10:00:00Z"
  }
]
```

### GET /users/{id}

指定されたIDのユーザーを取得します。

**パス パラメータ**
- `id` (string): ユーザーID（UUID形式）

**レスポンス（200 OK）**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "田中太郎",
  "email": "tanaka@example.com",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T10:00:00Z"
}
```

**エラー**
- `400 Bad Request`: 無効なUUID形式
- `404 Not Found`: ユーザーが見つからない

### PUT /users/{id}

指定されたIDのユーザーを更新します。

**パス パラメータ**
- `id` (string): ユーザーID（UUID形式）

**リクエストボディ**
```json
{
  "name": "田中花子",
  "email": "hanako@example.com"
}
```

**レスポンス（200 OK）**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "田中花子",
  "email": "hanako@example.com",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T11:00:00Z"
}
```

**エラー**
- `400 Bad Request`: バリデーションエラー
- `404 Not Found`: ユーザーが見つからない
- `409 Conflict`: メールアドレス重複

### DELETE /users/{id}

指定されたIDのユーザーを削除します。

**パス パラメータ**
- `id` (string): ユーザーID（UUID形式）

**レスポンス（204 No Content）**

**エラー**
- `400 Bad Request`: 無効なUUID形式
- `404 Not Found`: ユーザーが見つからない

---

## カテゴリ管理

### POST /categories

新しいカテゴリを作成します。

**リクエストボディ**
```json
{
  "name": "交通費",
  "description": "電車・バス・タクシーなどの交通費",
  "color": "#FF6B6B"
}
```

**レスポンス（201 Created）**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440001",
  "name": "交通費",
  "description": "電車・バス・タクシーなどの交通費",
  "color": "#FF6B6B",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T10:00:00Z"
}
```

**エラー**
- `400 Bad Request`: バリデーションエラー
- `409 Conflict`: カテゴリ名重複

### GET /categories

全てのカテゴリを取得します。

**レスポンス（200 OK）**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "name": "交通費",
    "description": "電車・バス・タクシーなどの交通費",
    "color": "#FF6B6B",
    "created_at": "2023-10-01T10:00:00Z",
    "updated_at": "2023-10-01T10:00:00Z"
  }
]
```

### GET /categories/{id}

指定されたIDのカテゴリを取得します。

**パス パラメータ**
- `id` (string): カテゴリID（UUID形式）

**レスポンス（200 OK）**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440001",
  "name": "交通費",
  "description": "電車・バス・タクシーなどの交通費",
  "color": "#FF6B6B",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T10:00:00Z"
}
```

**エラー**
- `400 Bad Request`: 無効なUUID形式
- `404 Not Found`: カテゴリが見つからない

### PUT /categories/{id}

指定されたIDのカテゴリを更新します。

**パス パラメータ**
- `id` (string): カテゴリID（UUID形式）

**リクエストボディ**
```json
{
  "name": "交通費（更新）",
  "description": "更新された説明",
  "color": "#00FF00"
}
```

**レスポンス（200 OK）**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440001",
  "name": "交通費（更新）",
  "description": "更新された説明",
  "color": "#00FF00",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T11:00:00Z"
}
```

**エラー**
- `400 Bad Request`: バリデーションエラー
- `404 Not Found`: カテゴリが見つからない
- `409 Conflict`: カテゴリ名重複

### DELETE /categories/{id}

指定されたIDのカテゴリを削除します。

**パス パラメータ**
- `id` (string): カテゴリID（UUID形式）

**レスポンス（204 No Content）**

**エラー**
- `400 Bad Request`: 無効なUUID形式
- `404 Not Found`: カテゴリが見つからない
- `409 Conflict`: カテゴリが使用中のため削除不可

---

## 経費管理

### POST /users/{id}/expenses

指定されたユーザーの新しい経費を作成します。

**パス パラメータ**
- `id` (string): ユーザーID（UUID形式）

**リクエストボディ**
```json
{
  "category_id": "550e8400-e29b-41d4-a716-446655440001",
  "amount": 1500,
  "currency": "JPY",
  "title": "渋谷駅からオフィス",
  "description": "営業訪問のための交通費",
  "date": "2023-10-01T00:00:00Z"
}
```

**レスポンス（201 Created）**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440002",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "category_id": "550e8400-e29b-41d4-a716-446655440001",
  "category": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "name": "交通費",
    "description": "電車・バス・タクシーなどの交通費",
    "color": "#FF6B6B",
    "created_at": "2023-10-01T10:00:00Z",
    "updated_at": "2023-10-01T10:00:00Z"
  },
  "amount": 1500,
  "currency": "JPY",
  "title": "渋谷駅からオフィス",
  "description": "営業訪問のための交通費",
  "date": "2023-10-01T00:00:00Z",
  "status": "draft",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T10:00:00Z"
}
```

**エラー**
- `400 Bad Request`: バリデーションエラー
- `404 Not Found`: ユーザーまたはカテゴリが見つからない

### GET /users/{id}/expenses

指定されたユーザーの経費一覧を取得します。

**パス パラメータ**
- `id` (string): ユーザーID（UUID形式）

**クエリ パラメータ**
- `status` (string, optional): 経費ステータスでフィルタ（draft, submitted, approved, rejected）

**レスポンス（200 OK）**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440002",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "category_id": "550e8400-e29b-41d4-a716-446655440001",
    "category": {
      "id": "550e8400-e29b-41d4-a716-446655440001",
      "name": "交通費",
      "description": "電車・バス・タクシーなどの交通費",
      "color": "#FF6B6B",
      "created_at": "2023-10-01T10:00:00Z",
      "updated_at": "2023-10-01T10:00:00Z"
    },
    "amount": 1500,
    "currency": "JPY",
    "title": "渋谷駅からオフィス",
    "description": "営業訪問のための交通費",
    "date": "2023-10-01T00:00:00Z",
    "status": "draft",
    "created_at": "2023-10-01T10:00:00Z",
    "updated_at": "2023-10-01T10:00:00Z"
  }
]
```

**エラー**
- `400 Bad Request`: 無効なUUID形式または無効なステータス
- `404 Not Found`: ユーザーが見つからない

### GET /expenses/{id}

指定されたIDの経費を取得します。

**パス パラメータ**
- `id` (string): 経費ID（UUID形式）

**レスポンス（200 OK）**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440002",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "category_id": "550e8400-e29b-41d4-a716-446655440001",
  "category": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "name": "交通費",
    "description": "電車・バス・タクシーなどの交通費",
    "color": "#FF6B6B",
    "created_at": "2023-10-01T10:00:00Z",
    "updated_at": "2023-10-01T10:00:00Z"
  },
  "amount": 1500,
  "currency": "JPY",
  "title": "渋谷駅からオフィス",
  "description": "営業訪問のための交通費",
  "date": "2023-10-01T00:00:00Z",
  "status": "draft",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T10:00:00Z"
}
```

**エラー**
- `400 Bad Request`: 無効なUUID形式
- `404 Not Found`: 経費が見つからない

### PUT /expenses/{id}

指定されたIDの経費を更新します（下書き状態のみ）。

**パス パラメータ**
- `id` (string): 経費ID（UUID形式）

**リクエストボディ**
```json
{
  "category_id": "550e8400-e29b-41d4-a716-446655440001",
  "amount": 2000,
  "currency": "JPY",
  "title": "更新されたタイトル",
  "description": "更新された説明",
  "date": "2023-10-02T00:00:00Z"
}
```

**レスポンス（200 OK）**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440002",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "category_id": "550e8400-e29b-41d4-a716-446655440001",
  "category": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "name": "交通費",
    "description": "電車・バス・タクシーなどの交通費",
    "color": "#FF6B6B",
    "created_at": "2023-10-01T10:00:00Z",
    "updated_at": "2023-10-01T10:00:00Z"
  },
  "amount": 2000,
  "currency": "JPY",
  "title": "更新されたタイトル",
  "description": "更新された説明",
  "date": "2023-10-02T00:00:00Z",
  "status": "draft",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T11:00:00Z"
}
```

**エラー**
- `400 Bad Request`: バリデーションエラーまたは更新不可能な状態
- `404 Not Found`: 経費またはカテゴリが見つからない

### DELETE /expenses/{id}

指定されたIDの経費を削除します。

**パス パラメータ**
- `id` (string): 経費ID（UUID形式）

**レスポンス（204 No Content）**

**エラー**
- `400 Bad Request`: 無効なUUID形式
- `404 Not Found`: 経費が見つからない

### POST /expenses/{id}/submit

指定されたIDの経費を申請します（下書き → 申請済み）。

**パス パラメータ**
- `id` (string): 経費ID（UUID形式）

**レスポンス（200 OK）**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440002",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "category_id": "550e8400-e29b-41d4-a716-446655440001",
  "category": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "name": "交通費",
    "description": "電車・バス・タクシーなどの交通費",
    "color": "#FF6B6B",
    "created_at": "2023-10-01T10:00:00Z",
    "updated_at": "2023-10-01T10:00:00Z"
  },
  "amount": 1500,
  "currency": "JPY",
  "title": "渋谷駅からオフィス",
  "description": "営業訪問のための交通費",
  "date": "2023-10-01T00:00:00Z",
  "status": "submitted",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T11:00:00Z"
}
```

**エラー**
- `400 Bad Request`: 無効なUUID形式または申請不可能な状態
- `404 Not Found`: 経費が見つからない

### POST /expenses/{id}/approve

指定されたIDの経費を承認します（申請済み → 承認済み）。

**パス パラメータ**
- `id` (string): 経費ID（UUID形式）

**レスポンス（200 OK）**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440002",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "category_id": "550e8400-e29b-41d4-a716-446655440001",
  "category": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "name": "交通費",
    "description": "電車・バス・タクシーなどの交通費",
    "color": "#FF6B6B",
    "created_at": "2023-10-01T10:00:00Z",
    "updated_at": "2023-10-01T10:00:00Z"
  },
  "amount": 1500,
  "currency": "JPY",
  "title": "渋谷駅からオフィス",
  "description": "営業訪問のための交通費",
  "date": "2023-10-01T00:00:00Z",
  "status": "approved",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T11:30:00Z"
}
```

**エラー**
- `400 Bad Request`: 無効なUUID形式または承認不可能な状態
- `404 Not Found`: 経費が見つからない

### POST /expenses/{id}/reject

指定されたIDの経費を却下します（申請済み → 却下）。

**パス パラメータ**
- `id` (string): 経費ID（UUID形式）

**レスポンス（200 OK）**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440002",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "category_id": "550e8400-e29b-41d4-a716-446655440001",
  "category": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "name": "交通費",
    "description": "電車・バス・タクシーなどの交通費",
    "color": "#FF6B6B",
    "created_at": "2023-10-01T10:00:00Z",
    "updated_at": "2023-10-01T10:00:00Z"
  },
  "amount": 1500,
  "currency": "JPY",
  "title": "渋谷駅からオフィス",
  "description": "営業訪問のための交通費",
  "date": "2023-10-01T00:00:00Z",
  "status": "rejected",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T11:30:00Z"
}
```

**エラー**
- `400 Bad Request`: 無効なUUID形式または却下不可能な状態
- `404 Not Found`: 経費が見つからない

## ステータス遷移

経費のステータスは以下のように遷移します：

```
draft → submitted → approved/rejected
  ↑                     ↓
  └─── （編集可能） ←──────┘
```

- `draft`: 下書き状態（編集・削除・申請が可能）
- `submitted`: 申請済み状態（承認・却下が可能）
- `approved`: 承認済み状態（最終状態）
- `rejected`: 却下状態（最終状態）

## バリデーション

### ユーザー
- `name`: 必須、1-100文字
- `email`: 必須、有効なメールアドレス、255文字以内、ユニーク

### カテゴリ
- `name`: 必須、1-50文字、ユニーク
- `description`: 0-200文字
- `color`: 有効な16進数カラーコード（#RRGGBB）

### 経費
- `category_id`: 必須、有効なUUID、存在するカテゴリ
- `amount`: 必須、正の数値
- `currency`: 省略可（デフォルト: JPY）
- `title`: 必須、1-100文字
- `description`: 0-500文字
- `date`: 必須、未来日付不可、1年以上前不可
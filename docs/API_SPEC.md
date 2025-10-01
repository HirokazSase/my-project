# API仕様書

## 概要

経費精算システムのREST API仕様書です。

- Base URL: `http://localhost:8080`
- Content-Type: `application/json`
- 日時フォーマット: RFC3339 (`2023-10-01T09:00:00Z`)

## 共通レスポンス

### 成功レスポンス
- 200 OK: 取得成功
- 201 Created: 作成成功
- 204 No Content: 削除成功

### エラーレスポンス
```json
{
  "error": "ERROR_CODE",
  "message": "エラーメッセージ",
  "details": "詳細情報（任意）"
}
```

- 400 Bad Request: リクエスト形式エラー
- 404 Not Found: リソースが見つからない
- 409 Conflict: 重複エラー
- 500 Internal Server Error: サーバーエラー

## ユーザー管理 API

### ユーザー作成
```
POST /api/v1/users
```

**リクエスト**
```json
{
  "name": "田中太郎",
  "email": "tanaka@example.com"
}
```

**レスポンス (201 Created)**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "田中太郎",
  "email": "tanaka@example.com",
  "created_at": "2023-10-01T09:00:00Z",
  "updated_at": "2023-10-01T09:00:00Z"
}
```

### ユーザー一覧取得
```
GET /api/v1/users
```

**レスポンス (200 OK)**
```json
[
  {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "田中太郎",
    "email": "tanaka@example.com",
    "created_at": "2023-10-01T09:00:00Z",
    "updated_at": "2023-10-01T09:00:00Z"
  }
]
```

### ユーザー取得
```
GET /api/v1/users/{id}
```

**レスポンス (200 OK)**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "田中太郎",
  "email": "tanaka@example.com",
  "created_at": "2023-10-01T09:00:00Z",
  "updated_at": "2023-10-01T09:00:00Z"
}
```

### ユーザー更新
```
PUT /api/v1/users/{id}
```

**リクエスト**
```json
{
  "name": "田中二郎",
  "email": "tanaka2@example.com"
}
```

**レスポンス (200 OK)**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "田中二郎",
  "email": "tanaka2@example.com",
  "created_at": "2023-10-01T09:00:00Z",
  "updated_at": "2023-10-01T10:00:00Z"
}
```

### ユーザー削除
```
DELETE /api/v1/users/{id}
```

**レスポンス (204 No Content)**

## カテゴリ管理 API

### カテゴリ作成
```
POST /api/v1/categories
```

**リクエスト**
```json
{
  "name": "交通費",
  "description": "電車・バス・タクシーなどの交通費",
  "color": "#FF6B6B"
}
```

**レスポンス (201 Created)**
```json
{
  "id": "456e7890-e89b-12d3-a456-426614174000",
  "name": "交通費",
  "description": "電車・バス・タクシーなどの交通費",
  "color": "#FF6B6B",
  "created_at": "2023-10-01T09:00:00Z",
  "updated_at": "2023-10-01T09:00:00Z"
}
```

### カテゴリ一覧取得
```
GET /api/v1/categories
```

**レスポンス (200 OK)**
```json
[
  {
    "id": "456e7890-e89b-12d3-a456-426614174000",
    "name": "交通費",
    "description": "電車・バス・タクシーなどの交通費",
    "color": "#FF6B6B",
    "created_at": "2023-10-01T09:00:00Z",
    "updated_at": "2023-10-01T09:00:00Z"
  }
]
```

### カテゴリ取得
```
GET /api/v1/categories/{id}
```

**レスポンス (200 OK)**
```json
{
  "id": "456e7890-e89b-12d3-a456-426614174000",
  "name": "交通費",
  "description": "電車・バス・タクシーなどの交通費",
  "color": "#FF6B6B",
  "created_at": "2023-10-01T09:00:00Z",
  "updated_at": "2023-10-01T09:00:00Z"
}
```

### カテゴリ更新
```
PUT /api/v1/categories/{id}
```

**リクエスト**
```json
{
  "name": "交通費（更新）",
  "description": "更新された交通費カテゴリ",
  "color": "#00FF00"
}
```

**レスポンス (200 OK)**
```json
{
  "id": "456e7890-e89b-12d3-a456-426614174000",
  "name": "交通費（更新）",
  "description": "更新された交通費カテゴリ",
  "color": "#00FF00",
  "created_at": "2023-10-01T09:00:00Z",
  "updated_at": "2023-10-01T10:00:00Z"
}
```

### カテゴリ削除
```
DELETE /api/v1/categories/{id}
```

**レスポンス (204 No Content)**

**注意**: このカテゴリを使用している経費が存在する場合は削除できません (409 Conflict)

## 経費管理 API

### 経費作成
```
POST /api/v1/users/{user_id}/expenses
```

**リクエスト**
```json
{
  "category_id": "456e7890-e89b-12d3-a456-426614174000",
  "amount": 1500,
  "currency": "JPY",
  "title": "渋谷駅からオフィス",
  "description": "営業訪問のための交通費",
  "date": "2023-10-01T09:00:00Z"
}
```

**レスポンス (201 Created)**
```json
{
  "id": "789e0123-e89b-12d3-a456-426614174000",
  "user_id": "123e4567-e89b-12d3-a456-426614174000",
  "category_id": "456e7890-e89b-12d3-a456-426614174000",
  "category": {
    "id": "456e7890-e89b-12d3-a456-426614174000",
    "name": "交通費",
    "description": "電車・バス・タクシーなどの交通費",
    "color": "#FF6B6B",
    "created_at": "2023-10-01T09:00:00Z",
    "updated_at": "2023-10-01T09:00:00Z"
  },
  "amount": 1500,
  "currency": "JPY",
  "title": "渋谷駅からオフィス",
  "description": "営業訪問のための交通費",
  "date": "2023-10-01T09:00:00Z",
  "status": "draft",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T10:00:00Z"
}
```

### ユーザー経費一覧取得
```
GET /api/v1/users/{user_id}/expenses
```

**クエリパラメータ**
- `status` (任意): 経費ステータス (`draft`, `submitted`, `approved`, `rejected`)

**レスポンス (200 OK)**
```json
[
  {
    "id": "789e0123-e89b-12d3-a456-426614174000",
    "user_id": "123e4567-e89b-12d3-a456-426614174000",
    "category_id": "456e7890-e89b-12d3-a456-426614174000",
    "category": {
      "id": "456e7890-e89b-12d3-a456-426614174000",
      "name": "交通費",
      "description": "電車・バス・タクシーなどの交通費",
      "color": "#FF6B6B",
      "created_at": "2023-10-01T09:00:00Z",
      "updated_at": "2023-10-01T09:00:00Z"
    },
    "amount": 1500,
    "currency": "JPY",
    "title": "渋谷駅からオフィス",
    "description": "営業訪問のための交通費",
    "date": "2023-10-01T09:00:00Z",
    "status": "draft",
    "created_at": "2023-10-01T10:00:00Z",
    "updated_at": "2023-10-01T10:00:00Z"
  }
]
```

### 経費取得
```
GET /api/v1/expenses/{id}
```

**レスポンス (200 OK)**
```json
{
  "id": "789e0123-e89b-12d3-a456-426614174000",
  "user_id": "123e4567-e89b-12d3-a456-426614174000",
  "category_id": "456e7890-e89b-12d3-a456-426614174000",
  "category": {
    "id": "456e7890-e89b-12d3-a456-426614174000",
    "name": "交通費",
    "description": "電車・バス・タクシーなどの交通費",
    "color": "#FF6B6B",
    "created_at": "2023-10-01T09:00:00Z",
    "updated_at": "2023-10-01T09:00:00Z"
  },
  "amount": 1500,
  "currency": "JPY",
  "title": "渋谷駅からオフィス",
  "description": "営業訪問のための交通費",
  "date": "2023-10-01T09:00:00Z",
  "status": "draft",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T10:00:00Z"
}
```

### 経費更新
```
PUT /api/v1/expenses/{id}
```

**注意**: 下書き状態（`draft`）の経費のみ更新可能

**リクエスト**
```json
{
  "category_id": "456e7890-e89b-12d3-a456-426614174000",
  "amount": 2000,
  "currency": "JPY",
  "title": "新宿駅からオフィス",
  "description": "更新された交通費",
  "date": "2023-10-01T09:00:00Z"
}
```

**レスポンス (200 OK)**
```json
{
  "id": "789e0123-e89b-12d3-a456-426614174000",
  "user_id": "123e4567-e89b-12d3-a456-426614174000",
  "category_id": "456e7890-e89b-12d3-a456-426614174000",
  "category": {
    "id": "456e7890-e89b-12d3-a456-426614174000",
    "name": "交通費",
    "description": "電車・バス・タクシーなどの交通費",
    "color": "#FF6B6B",
    "created_at": "2023-10-01T09:00:00Z",
    "updated_at": "2023-10-01T09:00:00Z"
  },
  "amount": 2000,
  "currency": "JPY",
  "title": "新宿駅からオフィス",
  "description": "更新された交通費",
  "date": "2023-10-01T09:00:00Z",
  "status": "draft",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T11:00:00Z"
}
```

### 経費削除
```
DELETE /api/v1/expenses/{id}
```

**レスポンス (204 No Content)**

## 経費ステータス管理 API

### 経費申請
```
POST /api/v1/expenses/{id}/submit
```

**注意**: 下書き状態（`draft`）の経費のみ申請可能

**レスポンス (200 OK)**
```json
{
  "id": "789e0123-e89b-12d3-a456-426614174000",
  "user_id": "123e4567-e89b-12d3-a456-426614174000",
  "category_id": "456e7890-e89b-12d3-a456-426614174000",
  "category": {
    "id": "456e7890-e89b-12d3-a456-426614174000",
    "name": "交通費",
    "description": "電車・バス・タクシーなどの交通費",
    "color": "#FF6B6B",
    "created_at": "2023-10-01T09:00:00Z",
    "updated_at": "2023-10-01T09:00:00Z"
  },
  "amount": 1500,
  "currency": "JPY",
  "title": "渋谷駅からオフィス",
  "description": "営業訪問のための交通費",
  "date": "2023-10-01T09:00:00Z",
  "status": "submitted",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T12:00:00Z"
}
```

### 経費承認
```
POST /api/v1/expenses/{id}/approve
```

**注意**: 申請済み状態（`submitted`）の経費のみ承認可能

**レスポンス (200 OK)**
```json
{
  "id": "789e0123-e89b-12d3-a456-426614174000",
  "user_id": "123e4567-e89b-12d3-a456-426614174000",
  "category_id": "456e7890-e89b-12d3-a456-426614174000",
  "category": {
    "id": "456e7890-e89b-12d3-a456-426614174000",
    "name": "交通費",
    "description": "電車・バス・タクシーなどの交通費",
    "color": "#FF6B6B",
    "created_at": "2023-10-01T09:00:00Z",
    "updated_at": "2023-10-01T09:00:00Z"
  },
  "amount": 1500,
  "currency": "JPY",
  "title": "渋谷駅からオフィス",
  "description": "営業訪問のための交通費",
  "date": "2023-10-01T09:00:00Z",
  "status": "approved",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T13:00:00Z"
}
```

### 経費却下
```
POST /api/v1/expenses/{id}/reject
```

**注意**: 申請済み状態（`submitted`）の経費のみ却下可能

**レスポンス (200 OK)**
```json
{
  "id": "789e0123-e89b-12d3-a456-426614174000",
  "user_id": "123e4567-e89b-12d3-a456-426614174000",
  "category_id": "456e7890-e89b-12d3-a456-426614174000",
  "category": {
    "id": "456e7890-e89b-12d3-a456-426614174000",
    "name": "交通費",
    "description": "電車・バス・タクシーなどの交通費",
    "color": "#FF6B6B",
    "created_at": "2023-10-01T09:00:00Z",
    "updated_at": "2023-10-01T09:00:00Z"
  },
  "amount": 1500,
  "currency": "JPY",
  "title": "渋谷駅からオフィス",
  "description": "営業訪問のための交通費",
  "date": "2023-10-01T09:00:00Z",
  "status": "rejected",
  "created_at": "2023-10-01T10:00:00Z",
  "updated_at": "2023-10-01T13:00:00Z"
}
```

## ヘルスチェック API

### ヘルスチェック
```
GET /health
```

**レスポンス (200 OK)**
```json
{
  "status": "ok",
  "message": "Expense Management System is running"
}
```

## ステータス遷移図

```
経費ステータス遷移:

draft (下書き)
  ↓ submit
submitted (申請済み)
  ↓ approve / reject
approved (承認済み) / rejected (却下)
```

## バリデーションルール

### ユーザー
- `name`: 必須、1-100文字
- `email`: 必須、有効なメールアドレス形式、255文字以内、重複不可

### カテゴリ
- `name`: 必須、1-50文字、重複不可
- `description`: 任意、200文字以内
- `color`: 任意、有効な16進数カラーコード (#RRGGBB)

### 経費
- `category_id`: 必須、有効なカテゴリID
- `amount`: 必須、0以上の数値
- `currency`: 任意、デフォルト "JPY"
- `title`: 必須、1-100文字
- `description`: 任意、500文字以内
- `date`: 必須、過去1年以内、未来日不可

## エラーコード一覧

| コード | 説明 |
|-------|------|
| INVALID_REQUEST | リクエスト形式エラー |
| VALIDATION_FAILED | バリデーションエラー |
| USER_NOT_FOUND | ユーザーが見つからない |
| CATEGORY_NOT_FOUND | カテゴリが見つからない |
| EXPENSE_NOT_FOUND | 経費が見つからない |
| EMAIL_ALREADY_EXISTS | メールアドレスが既に存在 |
| CATEGORY_NAME_ALREADY_EXISTS | カテゴリ名が既に存在 |
| CATEGORY_IN_USE | カテゴリが使用中のため削除不可 |
| EXPENSE_UPDATE_NOT_ALLOWED | 経費更新不可 |
| EXPENSE_SUBMIT_NOT_ALLOWED | 経費申請不可 |
| EXPENSE_APPROVE_NOT_ALLOWED | 経費承認不可 |
| EXPENSE_REJECT_NOT_ALLOWED | 経費却下不可 |
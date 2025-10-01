package dto

import "time"

// CreateExpenseRequest 経費作成リクエスト
type CreateExpenseRequest struct {
	CategoryID  string    `json:"category_id" binding:"required"`
	Amount      float64   `json:"amount" binding:"required,min=0"`
	Currency    string    `json:"currency"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" binding:"required"`
}

// UpdateExpenseRequest 経費更新リクエスト
type UpdateExpenseRequest struct {
	CategoryID  string    `json:"category_id" binding:"required"`
	Amount      float64   `json:"amount" binding:"required,min=0"`
	Currency    string    `json:"currency"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" binding:"required"`
}

// ExpenseResponse 経費レスポンス
type ExpenseResponse struct {
	ID          string                `json:"id"`
	UserID      string                `json:"user_id"`
	CategoryID  string                `json:"category_id"`
	Category    *CategoryResponse      `json:"category,omitempty"`
	Amount      float64               `json:"amount"`
	Currency    string                `json:"currency"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Date        time.Time             `json:"date"`
	Status      string                `json:"status"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
}

// ExpenseListRequest 経費一覧取得リクエスト
type ExpenseListRequest struct {
	UserID     string    `json:"user_id"`
	CategoryID string    `json:"category_id"`
	Status     string    `json:"status"`
	DateFrom   time.Time `json:"date_from"`
	DateTo     time.Time `json:"date_to"`
}

// ExpenseStatusChangeRequest 経費ステータス変更リクエスト
type ExpenseStatusChangeRequest struct {
	Status string `json:"status" binding:"required,oneof=submitted approved rejected"`
}
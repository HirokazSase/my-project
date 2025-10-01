package repository

import (
	"context"
	"expense-management-system/internal/domain/entity"
	"expense-management-system/internal/domain/valueobject"
	"time"
)

// ExpenseRepository 経費リポジトリインターフェース
type ExpenseRepository interface {
	// Save 経費を保存
	Save(ctx context.Context, expense *entity.Expense) error

	// FindByID IDで経費を検索
	FindByID(ctx context.Context, id *valueobject.ExpenseID) (*entity.Expense, error)

	// FindByUserID ユーザーIDで経費を検索
	FindByUserID(ctx context.Context, userID *valueobject.UserID) ([]*entity.Expense, error)

	// FindByUserIDAndStatus ユーザーIDとステータスで経費を検索
	FindByUserIDAndStatus(ctx context.Context, userID *valueobject.UserID, status entity.ExpenseStatus) ([]*entity.Expense, error)

	// FindByCategoryID カテゴリIDで経費を検索
	FindByCategoryID(ctx context.Context, categoryID *valueobject.CategoryID) ([]*entity.Expense, error)

	// FindByDateRange 日付範囲で経費を検索
	FindByDateRange(ctx context.Context, userID *valueobject.UserID, from, to time.Time) ([]*entity.Expense, error)

	// FindAll 全ての経費を取得
	FindAll(ctx context.Context) ([]*entity.Expense, error)

	// Update 経費を更新
	Update(ctx context.Context, expense *entity.Expense) error

	// Delete 経費を削除
	Delete(ctx context.Context, id *valueobject.ExpenseID) error

	// Exists 経費が存在するかチェック
	Exists(ctx context.Context, id *valueobject.ExpenseID) (bool, error)
}
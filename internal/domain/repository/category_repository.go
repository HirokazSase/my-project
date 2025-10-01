package repository

import (
	"context"
	"expense-management-system/internal/domain/entity"
	"expense-management-system/internal/domain/valueobject"
)

// CategoryRepository カテゴリリポジトリインターフェース
type CategoryRepository interface {
	// Save カテゴリを保存
	Save(ctx context.Context, category *entity.Category) error

	// FindByID IDでカテゴリを検索
	FindByID(ctx context.Context, id *valueobject.CategoryID) (*entity.Category, error)

	// FindByName 名前でカテゴリを検索
	FindByName(ctx context.Context, name string) (*entity.Category, error)

	// FindAll 全てのカテゴリを取得
	FindAll(ctx context.Context) ([]*entity.Category, error)

	// Update カテゴリを更新
	Update(ctx context.Context, category *entity.Category) error

	// Delete カテゴリを削除
	Delete(ctx context.Context, id *valueobject.CategoryID) error

	// Exists カテゴリが存在するかチェック
	Exists(ctx context.Context, id *valueobject.CategoryID) (bool, error)
}
package repository

import (
	"context"
	"expense-management-system/internal/domain/entity"
	"expense-management-system/internal/domain/valueobject"
)

// UserRepository ユーザーリポジトリインターフェース
type UserRepository interface {
	// Save ユーザーを保存
	Save(ctx context.Context, user *entity.User) error

	// FindByID IDでユーザーを検索
	FindByID(ctx context.Context, id *valueobject.UserID) (*entity.User, error)

	// FindByEmail メールアドレスでユーザーを検索
	FindByEmail(ctx context.Context, email string) (*entity.User, error)

	// FindAll 全てのユーザーを取得
	FindAll(ctx context.Context) ([]*entity.User, error)

	// Update ユーザーを更新
	Update(ctx context.Context, user *entity.User) error

	// Delete ユーザーを削除
	Delete(ctx context.Context, id *valueobject.UserID) error

	// Exists ユーザーが存在するかチェック
	Exists(ctx context.Context, id *valueobject.UserID) (bool, error)
}
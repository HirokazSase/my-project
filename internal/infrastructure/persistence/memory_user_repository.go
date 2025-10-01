package persistence

import (
	"context"
	"expense-management-system/internal/domain/entity"
	"expense-management-system/internal/domain/valueobject"
	"expense-management-system/pkg/errors"
	"sync"
)

// MemoryUserRepository メモリベースのユーザーリポジトリ実装
type MemoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*entity.User
}

// NewMemoryUserRepository MemoryUserRepositoryのコンストラクタ
func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[string]*entity.User),
	}
}

// Save ユーザーを保存
func (r *MemoryUserRepository) Save(ctx context.Context, user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID().String()] = user
	return nil
}

// FindByID IDでユーザーを検索
func (r *MemoryUserRepository) FindByID(ctx context.Context, id *valueobject.UserID) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id.String()]
	if !exists {
		return nil, errors.NewDomainError(errors.UserNotFound, "ユーザーが見つかりません")
	}

	return user, nil
}

// FindByEmail メールアドレスでユーザーを検索
func (r *MemoryUserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email() == email {
			return user, nil
		}
	}

	return nil, errors.NewDomainError(errors.UserNotFound, "ユーザーが見つかりません")
}

// FindAll 全てのユーザーを取得
func (r *MemoryUserRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*entity.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

// Update ユーザーを更新
func (r *MemoryUserRepository) Update(ctx context.Context, user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID().String()]; !exists {
		return errors.NewDomainError(errors.UserNotFound, "ユーザーが見つかりません")
	}

	r.users[user.ID().String()] = user
	return nil
}

// Delete ユーザーを削除
func (r *MemoryUserRepository) Delete(ctx context.Context, id *valueobject.UserID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id.String()]; !exists {
		return errors.NewDomainError(errors.UserNotFound, "ユーザーが見つかりません")
	}

	delete(r.users, id.String())
	return nil
}

// Exists ユーザーが存在するかチェック
func (r *MemoryUserRepository) Exists(ctx context.Context, id *valueobject.UserID) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.users[id.String()]
	return exists, nil
}
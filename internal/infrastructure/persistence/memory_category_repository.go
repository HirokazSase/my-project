package persistence

import (
	"context"
	"expense-management-system/internal/domain/entity"
	"expense-management-system/internal/domain/valueobject"
	"expense-management-system/pkg/errors"
	"sync"
)

// MemoryCategoryRepository メモリベースのカテゴリリポジトリ実装
type MemoryCategoryRepository struct {
	mu         sync.RWMutex
	categories map[string]*entity.Category
}

// NewMemoryCategoryRepository MemoryCategoryRepositoryのコンストラクタ
func NewMemoryCategoryRepository() *MemoryCategoryRepository {
	return &MemoryCategoryRepository{
		categories: make(map[string]*entity.Category),
	}
}

// Save カテゴリを保存
func (r *MemoryCategoryRepository) Save(ctx context.Context, category *entity.Category) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.categories[category.ID().String()] = category
	return nil
}

// FindByID IDでカテゴリを検索
func (r *MemoryCategoryRepository) FindByID(ctx context.Context, id *valueobject.CategoryID) (*entity.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	category, exists := r.categories[id.String()]
	if !exists {
		return nil, errors.NewDomainError(errors.CategoryNotFound, "カテゴリが見つかりません")
	}

	return category, nil
}

// FindByName 名前でカテゴリを検索
func (r *MemoryCategoryRepository) FindByName(ctx context.Context, name string) (*entity.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, category := range r.categories {
		if category.Name() == name {
			return category, nil
		}
	}

	return nil, errors.NewDomainError(errors.CategoryNotFound, "カテゴリが見つかりません")
}

// FindAll 全てのカテゴリを取得
func (r *MemoryCategoryRepository) FindAll(ctx context.Context) ([]*entity.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	categories := make([]*entity.Category, 0, len(r.categories))
	for _, category := range r.categories {
		categories = append(categories, category)
	}

	return categories, nil
}

// Update カテゴリを更新
func (r *MemoryCategoryRepository) Update(ctx context.Context, category *entity.Category) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.categories[category.ID().String()]; !exists {
		return errors.NewDomainError(errors.CategoryNotFound, "カテゴリが見つかりません")
	}

	r.categories[category.ID().String()] = category
	return nil
}

// Delete カテゴリを削除
func (r *MemoryCategoryRepository) Delete(ctx context.Context, id *valueobject.CategoryID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.categories[id.String()]; !exists {
		return errors.NewDomainError(errors.CategoryNotFound, "カテゴリが見つかりません")
	}

	delete(r.categories, id.String())
	return nil
}

// Exists カテゴリが存在するかチェック
func (r *MemoryCategoryRepository) Exists(ctx context.Context, id *valueobject.CategoryID) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.categories[id.String()]
	return exists, nil
}
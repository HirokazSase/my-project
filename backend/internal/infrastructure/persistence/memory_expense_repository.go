package persistence

import (
	"context"
	"expense-management-system/internal/domain/entity"
	"expense-management-system/internal/domain/valueobject"
	"expense-management-system/pkg/errors"
	"sync"
	"time"
)

// MemoryExpenseRepository メモリベースの経費リポジトリ実装
type MemoryExpenseRepository struct {
	mu       sync.RWMutex
	expenses map[string]*entity.Expense
}

// NewMemoryExpenseRepository MemoryExpenseRepositoryのコンストラクタ
func NewMemoryExpenseRepository() *MemoryExpenseRepository {
	return &MemoryExpenseRepository{
		expenses: make(map[string]*entity.Expense),
	}
}

// Save 経費を保存
func (r *MemoryExpenseRepository) Save(ctx context.Context, expense *entity.Expense) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.expenses[expense.ID().String()] = expense
	return nil
}

// FindByID IDで経費を検索
func (r *MemoryExpenseRepository) FindByID(ctx context.Context, id *valueobject.ExpenseID) (*entity.Expense, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	expense, exists := r.expenses[id.String()]
	if !exists {
		return nil, errors.NewDomainError(errors.ExpenseNotFound, "経費が見つかりません")
	}

	return expense, nil
}

// FindByUserID ユーザーIDで経費を検索
func (r *MemoryExpenseRepository) FindByUserID(ctx context.Context, userID *valueobject.UserID) ([]*entity.Expense, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	expenses := make([]*entity.Expense, 0)
	for _, expense := range r.expenses {
		if expense.UserID().Equals(userID) {
			expenses = append(expenses, expense)
		}
	}

	return expenses, nil
}

// FindByUserIDAndStatus ユーザーIDとステータスで経費を検索
func (r *MemoryExpenseRepository) FindByUserIDAndStatus(ctx context.Context, userID *valueobject.UserID, status entity.ExpenseStatus) ([]*entity.Expense, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	expenses := make([]*entity.Expense, 0)
	for _, expense := range r.expenses {
		if expense.UserID().Equals(userID) && expense.Status() == status {
			expenses = append(expenses, expense)
		}
	}

	return expenses, nil
}

// FindByCategoryID カテゴリIDで経費を検索
func (r *MemoryExpenseRepository) FindByCategoryID(ctx context.Context, categoryID *valueobject.CategoryID) ([]*entity.Expense, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	expenses := make([]*entity.Expense, 0)
	for _, expense := range r.expenses {
		if expense.CategoryID().Equals(categoryID) {
			expenses = append(expenses, expense)
		}
	}

	return expenses, nil
}

// FindByDateRange 日付範囲で経費を検索
func (r *MemoryExpenseRepository) FindByDateRange(ctx context.Context, userID *valueobject.UserID, from, to time.Time) ([]*entity.Expense, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	expenses := make([]*entity.Expense, 0)
	for _, expense := range r.expenses {
		if expense.UserID().Equals(userID) {
			expenseDate := expense.Date()
			if (expenseDate.Equal(from) || expenseDate.After(from)) &&
				(expenseDate.Equal(to) || expenseDate.Before(to)) {
				expenses = append(expenses, expense)
			}
		}
	}

	return expenses, nil
}

// FindAll 全ての経費を取得
func (r *MemoryExpenseRepository) FindAll(ctx context.Context) ([]*entity.Expense, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	expenses := make([]*entity.Expense, 0, len(r.expenses))
	for _, expense := range r.expenses {
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

// Update 経費を更新
func (r *MemoryExpenseRepository) Update(ctx context.Context, expense *entity.Expense) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.expenses[expense.ID().String()]; !exists {
		return errors.NewDomainError(errors.ExpenseNotFound, "経費が見つかりません")
	}

	r.expenses[expense.ID().String()] = expense
	return nil
}

// Delete 経費を削除
func (r *MemoryExpenseRepository) Delete(ctx context.Context, id *valueobject.ExpenseID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.expenses[id.String()]; !exists {
		return errors.NewDomainError(errors.ExpenseNotFound, "経費が見つかりません")
	}

	delete(r.expenses, id.String())
	return nil
}

// Exists 経費が存在するかチェック
func (r *MemoryExpenseRepository) Exists(ctx context.Context, id *valueobject.ExpenseID) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.expenses[id.String()]
	return exists, nil
}

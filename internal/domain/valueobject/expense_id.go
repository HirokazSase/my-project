package valueobject

import (
	"expense-management-system/pkg/errors"
	"strings"

	"github.com/google/uuid"
)

// ExpenseID 経費IDを表すValue Object
type ExpenseID struct {
	value string
}

// NewExpenseID 新しいExpenseIDを作成
func NewExpenseID(value string) (*ExpenseID, error) {
	if strings.TrimSpace(value) == "" {
		return nil, errors.NewDomainError(errors.InvalidUserID, "経費IDは空文字列にできません")
	}

	// UUIDの形式チェック
	if _, err := uuid.Parse(value); err != nil {
		return nil, errors.NewDomainError(errors.InvalidUserID, "経費IDは有効なUUID形式である必要があります")
	}

	return &ExpenseID{value: value}, nil
}

// GenerateExpenseID 新しいExpenseIDを生成
func GenerateExpenseID() *ExpenseID {
	return &ExpenseID{value: uuid.New().String()}
}

// Value 値を取得
func (e *ExpenseID) Value() string {
	return e.value
}

// Equals 等価性をチェック
func (e *ExpenseID) Equals(other *ExpenseID) bool {
	if other == nil {
		return false
	}
	return e.value == other.value
}

// String 文字列表現
func (e *ExpenseID) String() string {
	return e.value
}
package valueobject

import (
	"expense-management-system/pkg/errors"
	"strings"

	"github.com/google/uuid"
)

// CategoryID カテゴリIDを表すValue Object
type CategoryID struct {
	value string
}

// NewCategoryID 新しいCategoryIDを作成
func NewCategoryID(value string) (*CategoryID, error) {
	if strings.TrimSpace(value) == "" {
		return nil, errors.NewDomainError(errors.InvalidCategoryID, "カテゴリIDは空文字列にできません")
	}

	// UUIDの形式チェック
	if _, err := uuid.Parse(value); err != nil {
		return nil, errors.NewDomainError(errors.InvalidCategoryID, "カテゴリIDは有効なUUID形式である必要があります")
	}

	return &CategoryID{value: value}, nil
}

// GenerateCategoryID 新しいCategoryIDを生成
func GenerateCategoryID() *CategoryID {
	return &CategoryID{value: uuid.New().String()}
}

// Value 値を取得
func (c *CategoryID) Value() string {
	return c.value
}

// Equals 等価性をチェック
func (c *CategoryID) Equals(other *CategoryID) bool {
	if other == nil {
		return false
	}
	return c.value == other.value
}

// String 文字列表現
func (c *CategoryID) String() string {
	return c.value
}
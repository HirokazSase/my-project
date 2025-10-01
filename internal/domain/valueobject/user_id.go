package valueobject

import (
	"expense-management-system/pkg/errors"
	"strings"

	"github.com/google/uuid"
)

// UserID ユーザーIDを表すValue Object
type UserID struct {
	value string
}

// NewUserID 新しいUserIDを作成
func NewUserID(value string) (*UserID, error) {
	if strings.TrimSpace(value) == "" {
		return nil, errors.NewDomainError(errors.InvalidUserID, "ユーザーIDは空文字列にできません")
	}

	// UUIDの形式チェック
	if _, err := uuid.Parse(value); err != nil {
		return nil, errors.NewDomainError(errors.InvalidUserID, "ユーザーIDは有効なUUID形式である必要があります")
	}

	return &UserID{value: value}, nil
}

// GenerateUserID 新しいUserIDを生成
func GenerateUserID() *UserID {
	return &UserID{value: uuid.New().String()}
}

// Value 値を取得
func (u *UserID) Value() string {
	return u.value
}

// Equals 等価性をチェック
func (u *UserID) Equals(other *UserID) bool {
	if other == nil {
		return false
	}
	return u.value == other.value
}

// String 文字列表現
func (u *UserID) String() string {
	return u.value
}
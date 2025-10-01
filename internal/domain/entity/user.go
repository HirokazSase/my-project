package entity

import (
	"expense-management-system/internal/domain/valueobject"
	"expense-management-system/pkg/errors"
	"strings"
	"time"
)

// User ユーザーエンティティ
type User struct {
	id        *valueobject.UserID
	name      string
	email     string
	createdAt time.Time
	updatedAt time.Time
}

// NewUser 新しいUserを作成
func NewUser(name, email string) (*User, error) {
	if err := validateUserName(name); err != nil {
		return nil, err
	}

	if err := validateUserEmail(email); err != nil {
		return nil, err
	}

	now := time.Now()
	return &User{
		id:        valueobject.GenerateUserID(),
		name:      strings.TrimSpace(name),
		email:     strings.TrimSpace(email),
		createdAt: now,
		updatedAt: now,
	}, nil
}

// ReconstructUser 既存データからUserを再構築
func ReconstructUser(id *valueobject.UserID, name, email string, createdAt, updatedAt time.Time) (*User, error) {
	if id == nil {
		return nil, errors.NewDomainError(errors.InvalidUserID, "ユーザーIDが必要です")
	}

	if err := validateUserName(name); err != nil {
		return nil, err
	}

	if err := validateUserEmail(email); err != nil {
		return nil, err
	}

	return &User{
		id:        id,
		name:      name,
		email:     email,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

// ID IDを取得
func (u *User) ID() *valueobject.UserID {
	return u.id
}

// Name 名前を取得
func (u *User) Name() string {
	return u.name
}

// Email メールアドレスを取得
func (u *User) Email() string {
	return u.email
}

// CreatedAt 作成日時を取得
func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

// UpdatedAt 更新日時を取得
func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

// UpdateProfile プロフィールを更新
func (u *User) UpdateProfile(name, email string) error {
	if err := validateUserName(name); err != nil {
		return err
	}

	if err := validateUserEmail(email); err != nil {
		return err
	}

	u.name = strings.TrimSpace(name)
	u.email = strings.TrimSpace(email)
	u.updatedAt = time.Now()

	return nil
}

// validateUserName ユーザー名のバリデーション
func validateUserName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.NewDomainError("INVALID_USER_NAME", "ユーザー名は必須です")
	}

	if len(name) > 100 {
		return errors.NewDomainError("INVALID_USER_NAME", "ユーザー名は100文字以内である必要があります")
	}

	return nil
}

// validateUserEmail メールアドレスのバリデーション
func validateUserEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return errors.NewDomainError("INVALID_USER_EMAIL", "メールアドレスは必須です")
	}

	// 簡単なメールアドレス形式チェック
	if !strings.Contains(email, "@") || len(email) < 5 {
		return errors.NewDomainError("INVALID_USER_EMAIL", "有効なメールアドレスを入力してください")
	}

	if len(email) > 255 {
		return errors.NewDomainError("INVALID_USER_EMAIL", "メールアドレスは255文字以内である必要があります")
	}

	return nil
}
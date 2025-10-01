package entity

import (
	"expense-management-system/internal/domain/valueobject"
	"expense-management-system/pkg/errors"
	"strings"
	"time"
)

// Category カテゴリエンティティ
type Category struct {
	id          *valueobject.CategoryID
	name        string
	description string
	color       string
	createdAt   time.Time
	updatedAt   time.Time
}

// NewCategory 新しいCategoryを作成
func NewCategory(name, description, color string) (*Category, error) {
	if err := validateCategoryName(name); err != nil {
		return nil, err
	}

	if err := validateCategoryDescription(description); err != nil {
		return nil, err
	}

	if err := validateCategoryColor(color); err != nil {
		return nil, err
	}

	now := time.Now()
	return &Category{
		id:          valueobject.GenerateCategoryID(),
		name:        strings.TrimSpace(name),
		description: strings.TrimSpace(description),
		color:       strings.TrimSpace(color),
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

// ReconstructCategory 既存データからCategoryを再構築
func ReconstructCategory(id *valueobject.CategoryID, name, description, color string, createdAt, updatedAt time.Time) (*Category, error) {
	if id == nil {
		return nil, errors.NewDomainError(errors.InvalidCategoryID, "カテゴリIDが必要です")
	}

	if err := validateCategoryName(name); err != nil {
		return nil, err
	}

	if err := validateCategoryDescription(description); err != nil {
		return nil, err
	}

	if err := validateCategoryColor(color); err != nil {
		return nil, err
	}

	return &Category{
		id:          id,
		name:        name,
		description: description,
		color:       color,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}, nil
}

// ID IDを取得
func (c *Category) ID() *valueobject.CategoryID {
	return c.id
}

// Name 名前を取得
func (c *Category) Name() string {
	return c.name
}

// Description 説明を取得
func (c *Category) Description() string {
	return c.description
}

// Color 色を取得
func (c *Category) Color() string {
	return c.color
}

// CreatedAt 作成日時を取得
func (c *Category) CreatedAt() time.Time {
	return c.createdAt
}

// UpdatedAt 更新日時を取得
func (c *Category) UpdatedAt() time.Time {
	return c.updatedAt
}

// Update カテゴリ情報を更新
func (c *Category) Update(name, description, color string) error {
	if err := validateCategoryName(name); err != nil {
		return err
	}

	if err := validateCategoryDescription(description); err != nil {
		return err
	}

	if err := validateCategoryColor(color); err != nil {
		return err
	}

	c.name = strings.TrimSpace(name)
	c.description = strings.TrimSpace(description)
	c.color = strings.TrimSpace(color)
	c.updatedAt = time.Now()

	return nil
}

// validateCategoryName カテゴリ名のバリデーション
func validateCategoryName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.NewDomainError("INVALID_CATEGORY_NAME", "カテゴリ名は必須です")
	}

	if len(name) > 50 {
		return errors.NewDomainError("INVALID_CATEGORY_NAME", "カテゴリ名は50文字以内である必要があります")
	}

	return nil
}

// validateCategoryDescription カテゴリ説明のバリデーション
func validateCategoryDescription(description string) error {
	if len(description) > 200 {
		return errors.NewDomainError("INVALID_CATEGORY_DESCRIPTION", "カテゴリ説明は200文字以内である必要があります")
	}

	return nil
}

// validateCategoryColor カテゴリ色のバリデーション
func validateCategoryColor(color string) error {
	color = strings.TrimSpace(color)
	if color == "" {
		return nil // 色は任意
	}

	// 簡単な16進数カラーコードのチェック
	if len(color) == 7 && color[0] == '#' {
		for _, c := range color[1:] {
			if !((c >= '0' && c <= '9') || (c >= 'A' && c <= 'F') || (c >= 'a' && c <= 'f')) {
				return errors.NewDomainError("INVALID_CATEGORY_COLOR", "色は有効な16進数カラーコード（#RRGGBB）である必要があります")
			}
		}
		return nil
	}

	return errors.NewDomainError("INVALID_CATEGORY_COLOR", "色は有効な16進数カラーコード（#RRGGBB）である必要があります")
}
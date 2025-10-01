package usecase

import (
	"context"
	"expense-management-system/internal/application/dto"
	"expense-management-system/internal/domain/entity"
	"expense-management-system/internal/domain/repository"
	"expense-management-system/internal/domain/valueobject"
	"expense-management-system/pkg/errors"
)

// CategoryUseCase カテゴリユースケース
type CategoryUseCase struct {
	categoryRepo repository.CategoryRepository
	expenseRepo  repository.ExpenseRepository
}

// NewCategoryUseCase CategoryUseCaseのコンストラクタ
func NewCategoryUseCase(categoryRepo repository.CategoryRepository, expenseRepo repository.ExpenseRepository) *CategoryUseCase {
	return &CategoryUseCase{
		categoryRepo: categoryRepo,
		expenseRepo:  expenseRepo,
	}
}

// CreateCategory カテゴリを作成
func (uc *CategoryUseCase) CreateCategory(ctx context.Context, req *dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	// カテゴリ名の重複チェック
	existingCategory, err := uc.categoryRepo.FindByName(ctx, req.Name)
	if err == nil && existingCategory != nil {
		return nil, errors.NewApplicationError("CATEGORY_NAME_ALREADY_EXISTS", "このカテゴリ名は既に使用されています")
	}

	// 新しいカテゴリを作成
	category, err := entity.NewCategory(req.Name, req.Description, req.Color)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	// カテゴリを保存
	if err := uc.categoryRepo.Save(ctx, category); err != nil {
		return nil, errors.NewApplicationError("CATEGORY_CREATION_FAILED", "カテゴリの作成に失敗しました")
	}

	return &dto.CategoryResponse{
		ID:          category.ID().String(),
		Name:        category.Name(),
		Description: category.Description(),
		Color:       category.Color(),
		CreatedAt:   category.CreatedAt(),
		UpdatedAt:   category.UpdatedAt(),
	}, nil
}

// GetCategory カテゴリを取得
func (uc *CategoryUseCase) GetCategory(ctx context.Context, categoryID string) (*dto.CategoryResponse, error) {
	id, err := valueobject.NewCategoryID(categoryID)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	category, err := uc.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.NewApplicationError(errors.CategoryNotFound, "カテゴリが見つかりません")
	}

	return &dto.CategoryResponse{
		ID:          category.ID().String(),
		Name:        category.Name(),
		Description: category.Description(),
		Color:       category.Color(),
		CreatedAt:   category.CreatedAt(),
		UpdatedAt:   category.UpdatedAt(),
	}, nil
}

// UpdateCategory カテゴリを更新
func (uc *CategoryUseCase) UpdateCategory(ctx context.Context, categoryID string, req *dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {
	id, err := valueobject.NewCategoryID(categoryID)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	category, err := uc.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.NewApplicationError(errors.CategoryNotFound, "カテゴリが見つかりません")
	}

	// カテゴリ名の重複チェック（自分以外で同じ名前が存在するか）
	if req.Name != category.Name() {
		existingCategory, err := uc.categoryRepo.FindByName(ctx, req.Name)
		if err == nil && existingCategory != nil && !existingCategory.ID().Equals(category.ID()) {
			return nil, errors.NewApplicationError("CATEGORY_NAME_ALREADY_EXISTS", "このカテゴリ名は既に使用されています")
		}
	}

	// カテゴリ情報を更新
	if err := category.Update(req.Name, req.Description, req.Color); err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	// カテゴリを保存
	if err := uc.categoryRepo.Update(ctx, category); err != nil {
		return nil, errors.NewApplicationError("CATEGORY_UPDATE_FAILED", "カテゴリの更新に失敗しました")
	}

	return &dto.CategoryResponse{
		ID:          category.ID().String(),
		Name:        category.Name(),
		Description: category.Description(),
		Color:       category.Color(),
		CreatedAt:   category.CreatedAt(),
		UpdatedAt:   category.UpdatedAt(),
	}, nil
}

// DeleteCategory カテゴリを削除
func (uc *CategoryUseCase) DeleteCategory(ctx context.Context, categoryID string) error {
	id, err := valueobject.NewCategoryID(categoryID)
	if err != nil {
		return errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	exists, err := uc.categoryRepo.Exists(ctx, id)
	if err != nil {
		return errors.NewApplicationError("CATEGORY_DELETE_FAILED", "カテゴリの削除チェックに失敗しました")
	}

	if !exists {
		return errors.NewApplicationError(errors.CategoryNotFound, "カテゴリが見つかりません")
	}

	// このカテゴリを使用している経費があるかチェック
	expenses, err := uc.expenseRepo.FindByCategoryID(ctx, id)
	if err != nil {
		return errors.NewApplicationError("CATEGORY_DELETE_FAILED", "カテゴリの使用状況チェックに失敗しました")
	}

	if len(expenses) > 0 {
		return errors.NewApplicationError("CATEGORY_IN_USE", "このカテゴリは経費で使用されているため削除できません")
	}

	if err := uc.categoryRepo.Delete(ctx, id); err != nil {
		return errors.NewApplicationError("CATEGORY_DELETE_FAILED", "カテゴリの削除に失敗しました")
	}

	return nil
}

// GetAllCategories 全てのカテゴリを取得
func (uc *CategoryUseCase) GetAllCategories(ctx context.Context) ([]*dto.CategoryResponse, error) {
	categories, err := uc.categoryRepo.FindAll(ctx)
	if err != nil {
		return nil, errors.NewApplicationError("CATEGORY_FETCH_FAILED", "カテゴリ一覧の取得に失敗しました")
	}

	responses := make([]*dto.CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = &dto.CategoryResponse{
			ID:          category.ID().String(),
			Name:        category.Name(),
			Description: category.Description(),
			Color:       category.Color(),
			CreatedAt:   category.CreatedAt(),
			UpdatedAt:   category.UpdatedAt(),
		}
	}

	return responses, nil
}
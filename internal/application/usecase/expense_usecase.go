package usecase

import (
	"context"
	"expense-management-system/internal/application/dto"
	"expense-management-system/internal/domain/entity"
	"expense-management-system/internal/domain/repository"
	"expense-management-system/internal/domain/valueobject"
	"expense-management-system/pkg/errors"
)

// ExpenseUseCase 経費ユースケース
type ExpenseUseCase struct {
	expenseRepo  repository.ExpenseRepository
	userRepo     repository.UserRepository
	categoryRepo repository.CategoryRepository
}

// NewExpenseUseCase ExpenseUseCaseのコンストラクタ
func NewExpenseUseCase(
	expenseRepo repository.ExpenseRepository,
	userRepo repository.UserRepository,
	categoryRepo repository.CategoryRepository,
) *ExpenseUseCase {
	return &ExpenseUseCase{
		expenseRepo:  expenseRepo,
		userRepo:     userRepo,
		categoryRepo: categoryRepo,
	}
}

// CreateExpense 経費を作成
func (uc *ExpenseUseCase) CreateExpense(ctx context.Context, userID string, req *dto.CreateExpenseRequest) (*dto.ExpenseResponse, error) {
	// ユーザーIDの検証
	uid, err := valueobject.NewUserID(userID)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	// ユーザーの存在確認
	user, err := uc.userRepo.FindByID(ctx, uid)
	if err != nil {
		return nil, errors.NewApplicationError(errors.UserNotFound, "ユーザーが見つかりません")
	}

	// カテゴリIDの検証
	cid, err := valueobject.NewCategoryID(req.CategoryID)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	// カテゴリの存在確認
	category, err := uc.categoryRepo.FindByID(ctx, cid)
	if err != nil {
		return nil, errors.NewApplicationError(errors.CategoryNotFound, "カテゴリが見つかりません")
	}

	// 金額の作成
	currency := req.Currency
	if currency == "" {
		currency = "JPY"
	}
	amount, err := valueobject.NewMoney(req.Amount, currency)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	// 新しい経費を作成
	expense, err := entity.NewExpense(uid, cid, amount, req.Title, req.Description, req.Date)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	// 経費を保存
	if err := uc.expenseRepo.Save(ctx, expense); err != nil {
		return nil, errors.NewApplicationError(errors.ExpenseCreationFailed, "経費の作成に失敗しました")
	}

	return uc.buildExpenseResponse(expense, user, category), nil
}

// GetExpense 経費を取得
func (uc *ExpenseUseCase) GetExpense(ctx context.Context, expenseID string) (*dto.ExpenseResponse, error) {
	id, err := valueobject.NewExpenseID(expenseID)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	expense, err := uc.expenseRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ExpenseNotFound, "経費が見つかりません")
	}

	// ユーザーとカテゴリ情報を取得
	user, err := uc.userRepo.FindByID(ctx, expense.UserID())
	if err != nil {
		return nil, errors.NewApplicationError(errors.UserNotFound, "ユーザーが見つかりません")
	}

	category, err := uc.categoryRepo.FindByID(ctx, expense.CategoryID())
	if err != nil {
		return nil, errors.NewApplicationError(errors.CategoryNotFound, "カテゴリが見つかりません")
	}

	return uc.buildExpenseResponse(expense, user, category), nil
}

// UpdateExpense 経費を更新
func (uc *ExpenseUseCase) UpdateExpense(ctx context.Context, expenseID string, req *dto.UpdateExpenseRequest) (*dto.ExpenseResponse, error) {
	id, err := valueobject.NewExpenseID(expenseID)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	expense, err := uc.expenseRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ExpenseNotFound, "経費が見つかりません")
	}

	// カテゴリIDの検証
	cid, err := valueobject.NewCategoryID(req.CategoryID)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	// カテゴリの存在確認
	category, err := uc.categoryRepo.FindByID(ctx, cid)
	if err != nil {
		return nil, errors.NewApplicationError(errors.CategoryNotFound, "カテゴリが見つかりません")
	}

	// 金額の作成
	currency := req.Currency
	if currency == "" {
		currency = "JPY"
	}
	amount, err := valueobject.NewMoney(req.Amount, currency)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	// 経費情報を更新
	if err := expense.UpdateDetails(cid, amount, req.Title, req.Description, req.Date); err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	// 経費を保存
	if err := uc.expenseRepo.Update(ctx, expense); err != nil {
		return nil, errors.NewApplicationError(errors.ExpenseUpdateFailed, "経費の更新に失敗しました")
	}

	// ユーザー情報を取得
	user, err := uc.userRepo.FindByID(ctx, expense.UserID())
	if err != nil {
		return nil, errors.NewApplicationError(errors.UserNotFound, "ユーザーが見つかりません")
	}

	return uc.buildExpenseResponse(expense, user, category), nil
}

// DeleteExpense 経費を削除
func (uc *ExpenseUseCase) DeleteExpense(ctx context.Context, expenseID string) error {
	id, err := valueobject.NewExpenseID(expenseID)
	if err != nil {
		return errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	exists, err := uc.expenseRepo.Exists(ctx, id)
	if err != nil {
		return errors.NewApplicationError(errors.ExpenseDeletionFailed, "経費の削除チェックに失敗しました")
	}

	if !exists {
		return errors.NewApplicationError(errors.ExpenseNotFound, "経費が見つかりません")
	}

	if err := uc.expenseRepo.Delete(ctx, id); err != nil {
		return errors.NewApplicationError(errors.ExpenseDeletionFailed, "経費の削除に失敗しました")
	}

	return nil
}

// GetExpensesByUser ユーザーの経費一覧を取得
func (uc *ExpenseUseCase) GetExpensesByUser(ctx context.Context, userID string) ([]*dto.ExpenseResponse, error) {
	uid, err := valueobject.NewUserID(userID)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	// ユーザーの存在確認
	user, err := uc.userRepo.FindByID(ctx, uid)
	if err != nil {
		return nil, errors.NewApplicationError(errors.UserNotFound, "ユーザーが見つかりません")
	}

	expenses, err := uc.expenseRepo.FindByUserID(ctx, uid)
	if err != nil {
		return nil, errors.NewApplicationError("EXPENSE_FETCH_FAILED", "経費一覧の取得に失敗しました")
	}

	return uc.buildExpenseListResponse(ctx, expenses, user)
}

// GetExpensesByUserAndStatus ユーザーとステータスで経費一覧を取得
func (uc *ExpenseUseCase) GetExpensesByUserAndStatus(ctx context.Context, userID string, status string) ([]*dto.ExpenseResponse, error) {
	uid, err := valueobject.NewUserID(userID)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	// ユーザーの存在確認
	user, err := uc.userRepo.FindByID(ctx, uid)
	if err != nil {
		return nil, errors.NewApplicationError(errors.UserNotFound, "ユーザーが見つかりません")
	}

	// ステータスの検証
	expenseStatus := entity.ExpenseStatus(status)
	switch expenseStatus {
	case entity.ExpenseStatusDraft, entity.ExpenseStatusSubmitted, entity.ExpenseStatusApproved, entity.ExpenseStatusRejected:
		// 有効なステータス
	default:
		return nil, errors.NewApplicationError(errors.ValidationFailed, "無効なステータスです")
	}

	expenses, err := uc.expenseRepo.FindByUserIDAndStatus(ctx, uid, expenseStatus)
	if err != nil {
		return nil, errors.NewApplicationError("EXPENSE_FETCH_FAILED", "経費一覧の取得に失敗しました")
	}

	return uc.buildExpenseListResponse(ctx, expenses, user)
}

// SubmitExpense 経費を申請
func (uc *ExpenseUseCase) SubmitExpense(ctx context.Context, expenseID string) (*dto.ExpenseResponse, error) {
	return uc.changeExpenseStatus(ctx, expenseID, "submit")
}

// ApproveExpense 経費を承認
func (uc *ExpenseUseCase) ApproveExpense(ctx context.Context, expenseID string) (*dto.ExpenseResponse, error) {
	return uc.changeExpenseStatus(ctx, expenseID, "approve")
}

// RejectExpense 経費を却下
func (uc *ExpenseUseCase) RejectExpense(ctx context.Context, expenseID string) (*dto.ExpenseResponse, error) {
	return uc.changeExpenseStatus(ctx, expenseID, "reject")
}

// changeExpenseStatus 経費のステータスを変更
func (uc *ExpenseUseCase) changeExpenseStatus(ctx context.Context, expenseID string, action string) (*dto.ExpenseResponse, error) {
	id, err := valueobject.NewExpenseID(expenseID)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
	}

	expense, err := uc.expenseRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.NewApplicationError(errors.ExpenseNotFound, "経費が見つかりません")
	}

	// ステータス変更
	switch action {
	case "submit":
		if err := expense.Submit(); err != nil {
			return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
		}
	case "approve":
		if err := expense.Approve(); err != nil {
			return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
		}
	case "reject":
		if err := expense.Reject(); err != nil {
			return nil, errors.NewApplicationError(errors.ValidationFailed, err.Error())
		}
	default:
		return nil, errors.NewApplicationError(errors.ValidationFailed, "無効なアクションです")
	}

	// 経費を保存
	if err := uc.expenseRepo.Update(ctx, expense); err != nil {
		return nil, errors.NewApplicationError(errors.ExpenseUpdateFailed, "経費のステータス更新に失敗しました")
	}

	// ユーザーとカテゴリ情報を取得
	user, err := uc.userRepo.FindByID(ctx, expense.UserID())
	if err != nil {
		return nil, errors.NewApplicationError(errors.UserNotFound, "ユーザーが見つかりません")
	}

	category, err := uc.categoryRepo.FindByID(ctx, expense.CategoryID())
	if err != nil {
		return nil, errors.NewApplicationError(errors.CategoryNotFound, "カテゴリが見つかりません")
	}

	return uc.buildExpenseResponse(expense, user, category), nil
}

// buildExpenseResponse 経費レスポンスを構築
func (uc *ExpenseUseCase) buildExpenseResponse(expense *entity.Expense, user *entity.User, category *entity.Category) *dto.ExpenseResponse {
	return &dto.ExpenseResponse{
		ID:         expense.ID().String(),
		UserID:     expense.UserID().String(),
		CategoryID: expense.CategoryID().String(),
		Category: &dto.CategoryResponse{
			ID:          category.ID().String(),
			Name:        category.Name(),
			Description: category.Description(),
			Color:       category.Color(),
			CreatedAt:   category.CreatedAt(),
			UpdatedAt:   category.UpdatedAt(),
		},
		Amount:      expense.Amount().Amount(),
		Currency:    expense.Amount().Currency(),
		Title:       expense.Title(),
		Description: expense.Description(),
		Date:        expense.Date(),
		Status:      string(expense.Status()),
		CreatedAt:   expense.CreatedAt(),
		UpdatedAt:   expense.UpdatedAt(),
	}
}

// buildExpenseListResponse 経費リストレスポンスを構築
func (uc *ExpenseUseCase) buildExpenseListResponse(ctx context.Context, expenses []*entity.Expense, user *entity.User) ([]*dto.ExpenseResponse, error) {
	responses := make([]*dto.ExpenseResponse, len(expenses))

	for i, expense := range expenses {
		// カテゴリ情報を取得
		category, err := uc.categoryRepo.FindByID(ctx, expense.CategoryID())
		if err != nil {
			return nil, errors.NewApplicationError(errors.CategoryNotFound, "カテゴリが見つかりません")
		}

		responses[i] = uc.buildExpenseResponse(expense, user, category)
	}

	return responses, nil
}

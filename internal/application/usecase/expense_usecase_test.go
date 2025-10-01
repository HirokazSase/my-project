package usecase

import (
	"context"
	"expense-management-system/internal/application/dto"
	"expense-management-system/internal/domain/entity"
	"expense-management-system/internal/domain/valueobject"
	"expense-management-system/internal/infrastructure/persistence"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpenseUseCase_CreateExpense(t *testing.T) {
	ctx := context.Background()

	// リポジトリを初期化
	userRepo := persistence.NewMemoryUserRepository()
	categoryRepo := persistence.NewMemoryCategoryRepository()
	expenseRepo := persistence.NewMemoryExpenseRepository()

	// ユースケースを初期化
	useCase := NewExpenseUseCase(expenseRepo, userRepo, categoryRepo)

	// テスト用のユーザーとカテゴリを作成
	user, _ := entity.NewUser("テストユーザー", "test@example.com")
	err := userRepo.Save(ctx, user)
	require.NoError(t, err)

	category, _ := entity.NewCategory("交通費", "交通費カテゴリ", "#FF0000")
	err = categoryRepo.Save(ctx, category)
	require.NoError(t, err)

	tests := []struct {
		name    string
		userID  string
		req     *dto.CreateExpenseRequest
		wantErr bool
	}{
		{
			name:   "正常な経費作成",
			userID: user.ID().String(),
			req: &dto.CreateExpenseRequest{
				CategoryID:  category.ID().String(),
				Amount:      1000,
				Currency:    "JPY",
				Title:       "電車代",
				Description: "営業訪問のための電車代",
				Date:        time.Now().AddDate(0, 0, -1),
			},
			wantErr: false,
		},
		{
			name:   "存在しないユーザーID",
			userID: "invalid-user-id",
			req: &dto.CreateExpenseRequest{
				CategoryID:  category.ID().String(),
				Amount:      1000,
				Currency:    "JPY",
				Title:       "電車代",
				Description: "営業訪問のための電車代",
				Date:        time.Now().AddDate(0, 0, -1),
			},
			wantErr: true,
		},
		{
			name:   "存在しないカテゴリID",
			userID: user.ID().String(),
			req: &dto.CreateExpenseRequest{
				CategoryID:  "invalid-category-id",
				Amount:      1000,
				Currency:    "JPY",
				Title:       "電車代",
				Description: "営業訪問のための電車代",
				Date:        time.Now().AddDate(0, 0, -1),
			},
			wantErr: true,
		},
		{
			name:   "負の金額",
			userID: user.ID().String(),
			req: &dto.CreateExpenseRequest{
				CategoryID:  category.ID().String(),
				Amount:      -1000,
				Currency:    "JPY",
				Title:       "電車代",
				Description: "営業訪問のための電車代",
				Date:        time.Now().AddDate(0, 0, -1),
			},
			wantErr: true,
		},
		{
			name:   "空のタイトル",
			userID: user.ID().String(),
			req: &dto.CreateExpenseRequest{
				CategoryID:  category.ID().String(),
				Amount:      1000,
				Currency:    "JPY",
				Title:       "",
				Description: "営業訪問のための電車代",
				Date:        time.Now().AddDate(0, 0, -1),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := useCase.CreateExpense(ctx, tt.userID, tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.req.Title, result.Title)
				assert.Equal(t, tt.req.Amount, result.Amount)
				assert.Equal(t, "draft", result.Status)
				assert.NotNil(t, result.Category)
				assert.Equal(t, category.Name(), result.Category.Name)
			}
		})
	}
}

func TestExpenseUseCase_SubmitExpense(t *testing.T) {
	ctx := context.Background()

	// リポジトリを初期化
	userRepo := persistence.NewMemoryUserRepository()
	categoryRepo := persistence.NewMemoryCategoryRepository()
	expenseRepo := persistence.NewMemoryExpenseRepository()

	// ユースケースを初期化
	useCase := NewExpenseUseCase(expenseRepo, userRepo, categoryRepo)

	// テスト用のユーザー、カテゴリ、経費を作成
	user, _ := entity.NewUser("テストユーザー", "test@example.com")
	err := userRepo.Save(ctx, user)
	require.NoError(t, err)

	category, _ := entity.NewCategory("交通費", "交通費カテゴリ", "#FF0000")
	err = categoryRepo.Save(ctx, category)
	require.NoError(t, err)

	amount, _ := valueobject.NewMoney(1000, "JPY")
	expense, _ := entity.NewExpense(user.ID(), category.ID(), amount, "電車代", "営業訪問", time.Now().AddDate(0, 0, -1))
	err = expenseRepo.Save(ctx, expense)
	require.NoError(t, err)

	tests := []struct {
		name      string
		expenseID string
		wantErr   bool
	}{
		{
			name:      "正常な経費申請",
			expenseID: expense.ID().String(),
			wantErr:   false,
		},
		{
			name:      "存在しない経費ID",
			expenseID: "invalid-expense-id",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := useCase.SubmitExpense(ctx, tt.expenseID)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, "submitted", result.Status)
			}
		})
	}
}

func TestExpenseUseCase_ApproveExpense(t *testing.T) {
	ctx := context.Background()

	// リポジトリを初期化
	userRepo := persistence.NewMemoryUserRepository()
	categoryRepo := persistence.NewMemoryCategoryRepository()
	expenseRepo := persistence.NewMemoryExpenseRepository()

	// ユースケースを初期化
	useCase := NewExpenseUseCase(expenseRepo, userRepo, categoryRepo)

	// テスト用のユーザー、カテゴリ、経費を作成
	user, _ := entity.NewUser("テストユーザー", "test@example.com")
	err := userRepo.Save(ctx, user)
	require.NoError(t, err)

	category, _ := entity.NewCategory("交通費", "交通費カテゴリ", "#FF0000")
	err = categoryRepo.Save(ctx, category)
	require.NoError(t, err)

	amount, _ := valueobject.NewMoney(1000, "JPY")
	expense, _ := entity.NewExpense(user.ID(), category.ID(), amount, "電車代", "営業訪問", time.Now().AddDate(0, 0, -1))

	// 経費を申請状態にする
	err = expense.Submit()
	require.NoError(t, err)

	err = expenseRepo.Save(ctx, expense)
	require.NoError(t, err)

	t.Run("正常な経費承認", func(t *testing.T) {
		result, err := useCase.ApproveExpense(ctx, expense.ID().String())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "approved", result.Status)
	})
}

func TestExpenseUseCase_GetExpensesByUser(t *testing.T) {
	ctx := context.Background()

	// リポジトリを初期化
	userRepo := persistence.NewMemoryUserRepository()
	categoryRepo := persistence.NewMemoryCategoryRepository()
	expenseRepo := persistence.NewMemoryExpenseRepository()

	// ユースケースを初期化
	useCase := NewExpenseUseCase(expenseRepo, userRepo, categoryRepo)

	// テスト用のユーザーとカテゴリを作成
	user, _ := entity.NewUser("テストユーザー", "test@example.com")
	err := userRepo.Save(ctx, user)
	require.NoError(t, err)

	category, _ := entity.NewCategory("交通費", "交通費カテゴリ", "#FF0000")
	err = categoryRepo.Save(ctx, category)
	require.NoError(t, err)

	// テスト用の経費を複数作成
	amount1, _ := valueobject.NewMoney(1000, "JPY")
	expense1, _ := entity.NewExpense(user.ID(), category.ID(), amount1, "電車代1", "営業訪問1", time.Now().AddDate(0, 0, -1))
	err = expenseRepo.Save(ctx, expense1)
	require.NoError(t, err)

	amount2, _ := valueobject.NewMoney(2000, "JPY")
	expense2, _ := entity.NewExpense(user.ID(), category.ID(), amount2, "電車代2", "営業訪問2", time.Now().AddDate(0, 0, -2))
	err = expenseRepo.Save(ctx, expense2)
	require.NoError(t, err)

	t.Run("ユーザーの経費一覧取得", func(t *testing.T) {
		result, err := useCase.GetExpensesByUser(ctx, user.ID().String())
		require.NoError(t, err)
		assert.Len(t, result, 2)

		// IDで結果を特定
		var foundExpense1, foundExpense2 bool
		for _, exp := range result {
			if exp.ID == expense1.ID().String() {
				foundExpense1 = true
				assert.Equal(t, "電車代1", exp.Title)
				assert.Equal(t, 1000.0, exp.Amount)
			}
			if exp.ID == expense2.ID().String() {
				foundExpense2 = true
				assert.Equal(t, "電車代2", exp.Title)
				assert.Equal(t, 2000.0, exp.Amount)
			}
		}
		assert.True(t, foundExpense1, "expense1が見つかりませんでした")
		assert.True(t, foundExpense2, "expense2が見つかりませんでした")
	})

	t.Run("存在しないユーザーID", func(t *testing.T) {
		result, err := useCase.GetExpensesByUser(ctx, "invalid-user-id")
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

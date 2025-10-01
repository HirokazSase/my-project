package entity

import (
	"expense-management-system/internal/domain/valueobject"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewExpense(t *testing.T) {
	userID := valueobject.GenerateUserID()
	categoryID := valueobject.GenerateCategoryID()
	amount, _ := valueobject.NewMoney(1000, "JPY")
	validDate := time.Now().AddDate(0, 0, -1)

	tests := []struct {
		name        string
		userID      *valueobject.UserID
		categoryID  *valueobject.CategoryID
		amount      *valueobject.Money
		title       string
		description string
		date        time.Time
		wantErr     bool
	}{
		{
			name:        "正常な経費作成",
			userID:      userID,
			categoryID:  categoryID,
			amount:      amount,
			title:       "テスト経費",
			description: "テスト用の経費です",
			date:        validDate,
			wantErr:     false,
		},
		{
			name:        "ユーザーIDがnil",
			userID:      nil,
			categoryID:  categoryID,
			amount:      amount,
			title:       "テスト経費",
			description: "テスト用の経費です",
			date:        validDate,
			wantErr:     true,
		},
		{
			name:        "カテゴリIDがnil",
			userID:      userID,
			categoryID:  nil,
			amount:      amount,
			title:       "テスト経費",
			description: "テスト用の経費です",
			date:        validDate,
			wantErr:     true,
		},
		{
			name:        "金額がnil",
			userID:      userID,
			categoryID:  categoryID,
			amount:      nil,
			title:       "テスト経費",
			description: "テスト用の経費です",
			date:        validDate,
			wantErr:     true,
		},
		{
			name:        "タイトルが空文字",
			userID:      userID,
			categoryID:  categoryID,
			amount:      amount,
			title:       "",
			description: "テスト用の経費です",
			date:        validDate,
			wantErr:     true,
		},
		{
			name:        "未来の日付",
			userID:      userID,
			categoryID:  categoryID,
			amount:      amount,
			title:       "テスト経費",
			description: "テスト用の経費です",
			date:        time.Now().AddDate(0, 0, 2),
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expense, err := NewExpense(tt.userID, tt.categoryID, tt.amount, tt.title, tt.description, tt.date)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, expense)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, expense)
				assert.Equal(t, ExpenseStatusDraft, expense.Status())
				assert.Equal(t, tt.title, expense.Title())
				assert.Equal(t, tt.description, expense.Description())
				assert.True(t, expense.CanEdit())
				assert.True(t, expense.CanSubmit())
			}
		})
	}
}

func TestExpense_Submit(t *testing.T) {
	userID := valueobject.GenerateUserID()
	categoryID := valueobject.GenerateCategoryID()
	amount, _ := valueobject.NewMoney(1000, "JPY")
	validDate := time.Now().AddDate(0, 0, -1)

	t.Run("下書き状態から申請状態への変更", func(t *testing.T) {
		expense, err := NewExpense(userID, categoryID, amount, "テスト経費", "説明", validDate)
		require.NoError(t, err)

		err = expense.Submit()
		require.NoError(t, err)
		assert.Equal(t, ExpenseStatusSubmitted, expense.Status())
		assert.False(t, expense.CanEdit())
		assert.False(t, expense.CanSubmit())
	})

	t.Run("申請済み状態からの申請はエラー", func(t *testing.T) {
		expense, err := NewExpense(userID, categoryID, amount, "テスト経費", "説明", validDate)
		require.NoError(t, err)

		// 一度申請
		err = expense.Submit()
		require.NoError(t, err)

		// 再度申請を試行
		err = expense.Submit()
		assert.Error(t, err)
	})
}

func TestExpense_Approve(t *testing.T) {
	userID := valueobject.GenerateUserID()
	categoryID := valueobject.GenerateCategoryID()
	amount, _ := valueobject.NewMoney(1000, "JPY")
	validDate := time.Now().AddDate(0, 0, -1)

	t.Run("申請状態から承認状態への変更", func(t *testing.T) {
		expense, err := NewExpense(userID, categoryID, amount, "テスト経費", "説明", validDate)
		require.NoError(t, err)

		// 申請
		err = expense.Submit()
		require.NoError(t, err)

		// 承認
		err = expense.Approve()
		require.NoError(t, err)
		assert.Equal(t, ExpenseStatusApproved, expense.Status())
	})

	t.Run("下書き状態からの承認はエラー", func(t *testing.T) {
		expense, err := NewExpense(userID, categoryID, amount, "テスト経費", "説明", validDate)
		require.NoError(t, err)

		err = expense.Approve()
		assert.Error(t, err)
	})
}

func TestExpense_Reject(t *testing.T) {
	userID := valueobject.GenerateUserID()
	categoryID := valueobject.GenerateCategoryID()
	amount, _ := valueobject.NewMoney(1000, "JPY")
	validDate := time.Now().AddDate(0, 0, -1)

	t.Run("申請状態から却下状態への変更", func(t *testing.T) {
		expense, err := NewExpense(userID, categoryID, amount, "テスト経費", "説明", validDate)
		require.NoError(t, err)

		// 申請
		err = expense.Submit()
		require.NoError(t, err)

		// 却下
		err = expense.Reject()
		require.NoError(t, err)
		assert.Equal(t, ExpenseStatusRejected, expense.Status())
	})

	t.Run("下書き状態からの却下はエラー", func(t *testing.T) {
		expense, err := NewExpense(userID, categoryID, amount, "テスト経費", "説明", validDate)
		require.NoError(t, err)

		err = expense.Reject()
		assert.Error(t, err)
	})
}

func TestExpense_UpdateDetails(t *testing.T) {
	userID := valueobject.GenerateUserID()
	categoryID1 := valueobject.GenerateCategoryID()
	categoryID2 := valueobject.GenerateCategoryID()
	amount1, _ := valueobject.NewMoney(1000, "JPY")
	amount2, _ := valueobject.NewMoney(2000, "JPY")
	validDate := time.Now().AddDate(0, 0, -1)

	t.Run("下書き状態の経費詳細更新", func(t *testing.T) {
		expense, err := NewExpense(userID, categoryID1, amount1, "元のタイトル", "元の説明", validDate)
		require.NoError(t, err)

		newDate := time.Now().AddDate(0, 0, -2)
		err = expense.UpdateDetails(categoryID2, amount2, "新しいタイトル", "新しい説明", newDate)
		require.NoError(t, err)

		assert.Equal(t, categoryID2, expense.CategoryID())
		assert.Equal(t, amount2, expense.Amount())
		assert.Equal(t, "新しいタイトル", expense.Title())
		assert.Equal(t, "新しい説明", expense.Description())
		assert.Equal(t, newDate, expense.Date())
	})

	t.Run("申請済み状態の経費更新はエラー", func(t *testing.T) {
		expense, err := NewExpense(userID, categoryID1, amount1, "元のタイトル", "元の説明", validDate)
		require.NoError(t, err)

		// 申請
		err = expense.Submit()
		require.NoError(t, err)

		// 更新試行
		err = expense.UpdateDetails(categoryID2, amount2, "新しいタイトル", "新しい説明", validDate)
		assert.Error(t, err)
	})
}

func TestValidateExpenseDate(t *testing.T) {
	tests := []struct {
		name    string
		date    time.Time
		wantErr bool
	}{
		{
			name:    "昨日の日付",
			date:    time.Now().AddDate(0, 0, -1),
			wantErr: false,
		},
		{
			name:    "今日の日付",
			date:    time.Now(),
			wantErr: false,
		},
		{
			name:    "明日の日付（エラー）",
			date:    time.Now().AddDate(0, 0, 1),
			wantErr: true,
		},
		{
			name:    "1年以上前の日付（エラー）",
			date:    time.Now().AddDate(-1, -1, 0),
			wantErr: true,
		},
		{
			name:    "ゼロ値の日付（エラー）",
			date:    time.Time{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateExpenseDate(tt.date)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

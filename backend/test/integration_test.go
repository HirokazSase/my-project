package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"expense-management-system/internal/application/dto"
	"expense-management-system/internal/application/usecase"
	"expense-management-system/internal/infrastructure/persistence"
	"expense-management-system/internal/infrastructure/web"
	"expense-management-system/internal/infrastructure/web/handler"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupTestServer テスト用のサーバーをセットアップ
func setupTestServer() *httptest.Server {
	// リポジトリの初期化
	userRepo := persistence.NewMemoryUserRepository()
	categoryRepo := persistence.NewMemoryCategoryRepository()
	expenseRepo := persistence.NewMemoryExpenseRepository()

	// ユースケースの初期化
	userUseCase := usecase.NewUserUseCase(userRepo)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo, expenseRepo)
	expenseUseCase := usecase.NewExpenseUseCase(expenseRepo, userRepo, categoryRepo)

	// ハンドラーの初期化
	userHandler := handler.NewUserHandler(userUseCase)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)
	expenseHandler := handler.NewExpenseHandler(expenseUseCase)

	// ルーターの設定
	router := web.SetupRouter(userHandler, categoryHandler, expenseHandler)

	return httptest.NewServer(router)
}

// TestUserCRUD ユーザーのCRUD操作の統合テスト
func TestUserCRUD(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	client := &http.Client{}

	t.Run("ユーザー作成", func(t *testing.T) {
		reqBody := dto.CreateUserRequest{
			Name:  "田中太郎",
			Email: "tanaka@example.com",
		}

		body, _ := json.Marshal(reqBody)
		resp, err := client.Post(server.URL+"/api/v1/users", "application/json", bytes.NewBuffer(body))
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var user dto.UserResponse
		err = json.NewDecoder(resp.Body).Decode(&user)
		require.NoError(t, err)

		assert.NotEmpty(t, user.ID)
		assert.Equal(t, "田中太郎", user.Name)
		assert.Equal(t, "tanaka@example.com", user.Email)
	})

	t.Run("重複メールアドレスでのユーザー作成はエラー", func(t *testing.T) {
		// 同じメールアドレスで再度作成を試行
		reqBody := dto.CreateUserRequest{
			Name:  "佐藤花子",
			Email: "tanaka@example.com", // 重複
		}

		body, _ := json.Marshal(reqBody)
		resp, err := client.Post(server.URL+"/api/v1/users", "application/json", bytes.NewBuffer(body))
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusConflict, resp.StatusCode)
	})
}

// TestCategoryCRUD カテゴリのCRUD操作の統合テスト
func TestCategoryCRUD(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	client := &http.Client{}

	var categoryID string

	t.Run("カテゴリ作成", func(t *testing.T) {
		reqBody := dto.CreateCategoryRequest{
			Name:        "交通費",
			Description: "電車・バス・タクシーなどの交通費",
			Color:       "#FF6B6B",
		}

		body, _ := json.Marshal(reqBody)
		resp, err := client.Post(server.URL+"/api/v1/categories", "application/json", bytes.NewBuffer(body))
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var category dto.CategoryResponse
		err = json.NewDecoder(resp.Body).Decode(&category)
		require.NoError(t, err)

		categoryID = category.ID
		assert.NotEmpty(t, category.ID)
		assert.Equal(t, "交通費", category.Name)
		assert.Equal(t, "電車・バス・タクシーなどの交通費", category.Description)
		assert.Equal(t, "#FF6B6B", category.Color)
	})

	t.Run("カテゴリ取得", func(t *testing.T) {
		resp, err := client.Get(server.URL + "/api/v1/categories/" + categoryID)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var category dto.CategoryResponse
		err = json.NewDecoder(resp.Body).Decode(&category)
		require.NoError(t, err)

		assert.Equal(t, categoryID, category.ID)
		assert.Equal(t, "交通費", category.Name)
	})

	t.Run("カテゴリ更新", func(t *testing.T) {
		reqBody := dto.UpdateCategoryRequest{
			Name:        "交通費（更新）",
			Description: "更新された説明",
			Color:       "#00FF00",
		}

		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("PUT", server.URL+"/api/v1/categories/"+categoryID, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var category dto.CategoryResponse
		err = json.NewDecoder(resp.Body).Decode(&category)
		require.NoError(t, err)

		assert.Equal(t, "交通費（更新）", category.Name)
		assert.Equal(t, "更新された説明", category.Description)
		assert.Equal(t, "#00FF00", category.Color)
	})
}

// TestExpenseWorkflow 経費の全ワークフローの統合テスト
func TestExpenseWorkflow(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	client := &http.Client{}

	// 前提データの作成（ユーザーとカテゴリ）
	var userID, categoryID string

	// ユーザー作成
	userReq := dto.CreateUserRequest{
		Name:  "テストユーザー",
		Email: "test@example.com",
	}
	body, _ := json.Marshal(userReq)
	resp, err := client.Post(server.URL+"/api/v1/users", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	var user dto.UserResponse
	err = json.NewDecoder(resp.Body).Decode(&user)
	require.NoError(t, err)
	userID = user.ID

	// カテゴリ作成
	categoryReq := dto.CreateCategoryRequest{
		Name:        "交通費",
		Description: "交通費カテゴリ",
		Color:       "#FF0000",
	}
	body, _ = json.Marshal(categoryReq)
	resp, err = client.Post(server.URL+"/api/v1/categories", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	var category dto.CategoryResponse
	err = json.NewDecoder(resp.Body).Decode(&category)
	require.NoError(t, err)
	categoryID = category.ID

	var expenseID string

	t.Run("経費作成", func(t *testing.T) {
		expenseReq := dto.CreateExpenseRequest{
			CategoryID:  categoryID,
			Amount:      1500,
			Currency:    "JPY",
			Title:       "渋谷駅からオフィス",
			Description: "営業訪問のための交通費",
			Date:        time.Now().AddDate(0, 0, -1),
		}

		body, _ := json.Marshal(expenseReq)
		resp, err := client.Post(server.URL+"/api/v1/users/"+userID+"/expenses", "application/json", bytes.NewBuffer(body))
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var expense dto.ExpenseResponse
		err = json.NewDecoder(resp.Body).Decode(&expense)
		require.NoError(t, err)

		expenseID = expense.ID
		assert.NotEmpty(t, expense.ID)
		assert.Equal(t, "渋谷駅からオフィス", expense.Title)
		assert.Equal(t, 1500.0, expense.Amount)
		assert.Equal(t, "draft", expense.Status)
		assert.NotNil(t, expense.Category)
		assert.Equal(t, "交通費", expense.Category.Name)
	})

	t.Run("経費申請", func(t *testing.T) {
		req, _ := http.NewRequest("POST", server.URL+"/api/v1/expenses/"+expenseID+"/submit", nil)
		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var expense dto.ExpenseResponse
		err = json.NewDecoder(resp.Body).Decode(&expense)
		require.NoError(t, err)

		assert.Equal(t, "submitted", expense.Status)
	})

	t.Run("経費承認", func(t *testing.T) {
		req, _ := http.NewRequest("POST", server.URL+"/api/v1/expenses/"+expenseID+"/approve", nil)
		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var expense dto.ExpenseResponse
		err = json.NewDecoder(resp.Body).Decode(&expense)
		require.NoError(t, err)

		assert.Equal(t, "approved", expense.Status)
	})

	t.Run("ユーザーの経費一覧取得", func(t *testing.T) {
		resp, err := client.Get(server.URL + "/api/v1/users/" + userID + "/expenses")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var expenses []dto.ExpenseResponse
		err = json.NewDecoder(resp.Body).Decode(&expenses)
		require.NoError(t, err)

		assert.Len(t, expenses, 1)
		assert.Equal(t, expenseID, expenses[0].ID)
		assert.Equal(t, "approved", expenses[0].Status)
	})

	t.Run("ステータス別経費一覧取得", func(t *testing.T) {
		resp, err := client.Get(server.URL + "/api/v1/users/" + userID + "/expenses?status=approved")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var expenses []dto.ExpenseResponse
		err = json.NewDecoder(resp.Body).Decode(&expenses)
		require.NoError(t, err)

		assert.Len(t, expenses, 1)
		assert.Equal(t, "approved", expenses[0].Status)
	})
}

// TestHealthCheck ヘルスチェックエンドポイントのテスト
func TestHealthCheck(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/health")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, "ok", response["status"])
	assert.Equal(t, "Expense Management System is running", response["message"])
}

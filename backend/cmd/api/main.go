package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"expense-management-system/internal/application/dto"
	"expense-management-system/internal/application/usecase"
	"expense-management-system/internal/infrastructure/persistence"
	"expense-management-system/internal/infrastructure/web"
	"expense-management-system/internal/infrastructure/web/handler"
)

func main() {
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

	// サーバーの設定
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// サンプルデータの作成
	if err := createSampleData(userUseCase, categoryUseCase, expenseUseCase); err != nil {
		log.Printf("Failed to create sample data: %v", err)
	}

	// サーバー開始
	go func() {
		fmt.Printf("🚀 Server is running on port %s\n", port)
		fmt.Printf("📖 API Documentation: http://localhost:%s/health\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// グレースフルシャットダウン
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("✅ Server exited")
}

// createSampleData サンプルデータを作成
func createSampleData(userUseCase *usecase.UserUseCase, categoryUseCase *usecase.CategoryUseCase, expenseUseCase *usecase.ExpenseUseCase) error {
	ctx := context.Background()

	// サンプルユーザーを作成
	user1, err := userUseCase.CreateUser(ctx, &dto.CreateUserRequest{
		Name:  "田中太郎",
		Email: "tanaka@example.com",
	})
	if err != nil {
		return fmt.Errorf("failed to create user1: %w", err)
	}

	user2, err := userUseCase.CreateUser(ctx, &dto.CreateUserRequest{
		Name:  "佐藤花子",
		Email: "sato@example.com",
	})
	if err != nil {
		return fmt.Errorf("failed to create user2: %w", err)
	}

	// サンプルカテゴリを作成
	category1, err := categoryUseCase.CreateCategory(ctx, &dto.CreateCategoryRequest{
		Name:        "交通費",
		Description: "電車、バス、タクシーなどの交通費",
		Color:       "#FF6B6B",
	})
	if err != nil {
		return fmt.Errorf("failed to create category1: %w", err)
	}

	category2, err := categoryUseCase.CreateCategory(ctx, &dto.CreateCategoryRequest{
		Name:        "食費",
		Description: "会議や打ち合わせでの食事代",
		Color:       "#4ECDC4",
	})
	if err != nil {
		return fmt.Errorf("failed to create category2: %w", err)
	}

	category3, err := categoryUseCase.CreateCategory(ctx, &dto.CreateCategoryRequest{
		Name:        "事務用品",
		Description: "文具、オフィス用品等",
		Color:       "#45B7D1",
	})
	if err != nil {
		return fmt.Errorf("failed to create category3: %w", err)
	}

	// サンプル経費を作成
	_, err = expenseUseCase.CreateExpense(ctx, user1.ID, &dto.CreateExpenseRequest{
		CategoryID:  category1.ID,
		Amount:      500,
		Currency:    "JPY",
		Title:       "渋谷駅からオフィスまでの電車代",
		Description: "営業会議出席のための交通費",
		Date:        time.Now().AddDate(0, 0, -1),
	})
	if err != nil {
		return fmt.Errorf("failed to create expense1: %w", err)
	}

	_, err = expenseUseCase.CreateExpense(ctx, user1.ID, &dto.CreateExpenseRequest{
		CategoryID:  category2.ID,
		Amount:      1200,
		Currency:    "JPY",
		Title:       "クライアントとの会食",
		Description: "新規プロジェクトの打ち合わせランチ",
		Date:        time.Now().AddDate(0, 0, -2),
	})
	if err != nil {
		return fmt.Errorf("failed to create expense2: %w", err)
	}

	_, err = expenseUseCase.CreateExpense(ctx, user2.ID, &dto.CreateExpenseRequest{
		CategoryID:  category3.ID,
		Amount:      800,
		Currency:    "JPY",
		Title:       "プリンタ用紙購入",
		Description: "オフィス用のA4コピー用紙",
		Date:        time.Now().AddDate(0, 0, -3),
	})
	if err != nil {
		return fmt.Errorf("failed to create expense3: %w", err)
	}

	fmt.Println("✅ Sample data created successfully")
	fmt.Printf("👤 Users: %s (%s), %s (%s)\n", user1.Name, user1.Email, user2.Name, user2.Email)
	fmt.Printf("📁 Categories: %s, %s, %s\n", category1.Name, category2.Name, category3.Name)
	fmt.Println("💰 Sample expenses created for both users")

	return nil
}
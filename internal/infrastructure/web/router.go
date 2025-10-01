package web

import (
	"expense-management-system/internal/infrastructure/web/handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter ルーターを設定
func SetupRouter(
	userHandler *handler.UserHandler,
	categoryHandler *handler.CategoryHandler,
	expenseHandler *handler.ExpenseHandler,
) *gin.Engine {
	// Ginのモードを設定
	gin.SetMode(gin.ReleaseMode)
	
	router := gin.Default()

	// CORS対応
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// ヘルスチェック
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "Expense Management System is running",
		})
	})

	// APIバージョン1
	v1 := router.Group("/api/v1")
	{
		// ユーザー関連のルート
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)

			// ユーザーの経費関連のルート（同じパラメータ名を使用）
			users.GET("/:id/expenses", expenseHandler.GetExpensesByUser)
			users.POST("/:id/expenses", expenseHandler.CreateExpense)
		}

		// カテゴリ関連のルート
		categories := v1.Group("/categories")
		{
			categories.POST("", categoryHandler.CreateCategory)
			categories.GET("", categoryHandler.GetAllCategories)
			categories.GET("/:id", categoryHandler.GetCategory)
			categories.PUT("/:id", categoryHandler.UpdateCategory)
			categories.DELETE("/:id", categoryHandler.DeleteCategory)
		}

		// 経費関連のルート
		expenses := v1.Group("/expenses")
		{
			expenses.GET("/:id", expenseHandler.GetExpense)
			expenses.PUT("/:id", expenseHandler.UpdateExpense)
			expenses.DELETE("/:id", expenseHandler.DeleteExpense)

			// 経費ステータス変更のルート
			expenses.POST("/:id/submit", expenseHandler.SubmitExpense)
			expenses.POST("/:id/approve", expenseHandler.ApproveExpense)
			expenses.POST("/:id/reject", expenseHandler.RejectExpense)
		}
	}

	return router
}
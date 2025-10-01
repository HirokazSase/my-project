package handler

import (
	"expense-management-system/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse エラーレスポンス
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// handleError エラーハンドリング
func handleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *errors.DomainError:
		handleDomainError(c, e)
	case *errors.ApplicationError:
		handleApplicationError(c, e)
	default:
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "INTERNAL_SERVER_ERROR",
			Message: "内部サーバーエラーが発生しました",
			Details: err.Error(),
		})
	}
}

// handleDomainError ドメインエラーのハンドリング
func handleDomainError(c *gin.Context, err *errors.DomainError) {
	statusCode := http.StatusBadRequest

	switch err.Code {
	case errors.UserNotFound, errors.CategoryNotFound, errors.ExpenseNotFound:
		statusCode = http.StatusNotFound
	case errors.InvalidUserID, errors.InvalidCategoryID, errors.InvalidExpenseAmount:
		statusCode = http.StatusBadRequest
	}

	c.JSON(statusCode, ErrorResponse{
		Error:   err.Code,
		Message: err.Message,
	})
}

// handleApplicationError アプリケーションエラーのハンドリング
func handleApplicationError(c *gin.Context, err *errors.ApplicationError) {
	statusCode := http.StatusBadRequest

	switch err.Code {
	case errors.ValidationFailed:
		statusCode = http.StatusBadRequest
	case errors.ExpenseCreationFailed, errors.ExpenseUpdateFailed, errors.ExpenseDeletionFailed:
		statusCode = http.StatusInternalServerError
	case errors.UserCreationFailed, errors.UserUpdateFailed, errors.UserDeleteFailed:
		statusCode = http.StatusInternalServerError
	case errors.CategoryCreationFailed, errors.CategoryUpdateFailed, errors.CategoryDeleteFailed:
		statusCode = http.StatusInternalServerError
	case errors.EmailAlreadyExists, errors.CategoryNameExists:
		statusCode = http.StatusConflict
	case errors.CategoryInUse:
		statusCode = http.StatusConflict
	default:
		statusCode = http.StatusInternalServerError
	}

	c.JSON(statusCode, ErrorResponse{
		Error:   err.Code,
		Message: err.Message,
	})
}

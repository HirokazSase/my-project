package handler

import (
	"net/http"
	"expense-management-system/pkg/errors"

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
	case "USER_CREATION_FAILED", "USER_UPDATE_FAILED", "USER_DELETE_FAILED":
		statusCode = http.StatusInternalServerError
	case "CATEGORY_CREATION_FAILED", "CATEGORY_UPDATE_FAILED", "CATEGORY_DELETE_FAILED":
		statusCode = http.StatusInternalServerError
	case "EMAIL_ALREADY_EXISTS", "CATEGORY_NAME_ALREADY_EXISTS":
		statusCode = http.StatusConflict
	case "CATEGORY_IN_USE":
		statusCode = http.StatusConflict
	default:
		statusCode = http.StatusInternalServerError
	}

	c.JSON(statusCode, ErrorResponse{
		Error:   err.Code,
		Message: err.Message,
	})
}
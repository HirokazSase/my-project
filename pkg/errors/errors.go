// Package errors カスタムエラー型を定義
package errors

import "fmt"

// DomainError ドメインレイヤーのエラー
type DomainError struct {
	Code    string
	Message string
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("domain error [%s]: %s", e.Code, e.Message)
}

// NewDomainError ドメインエラーを作成
func NewDomainError(code, message string) error {
	return &DomainError{
		Code:    code,
		Message: message,
	}
}

// ApplicationError アプリケーションレイヤーのエラー
type ApplicationError struct {
	Code    string
	Message string
}

func (e *ApplicationError) Error() string {
	return fmt.Sprintf("application error [%s]: %s", e.Code, e.Message)
}

// NewApplicationError アプリケーションエラーを作成
func NewApplicationError(code, message string) error {
	return &ApplicationError{
		Code:    code,
		Message: message,
	}
}

// 定義済みエラーコード
const (
	// Domain errors
	InvalidExpenseAmount = "INVALID_EXPENSE_AMOUNT"
	InvalidUserID        = "INVALID_USER_ID"
	InvalidCategoryID    = "INVALID_CATEGORY_ID"
	ExpenseNotFound      = "EXPENSE_NOT_FOUND"
	UserNotFound         = "USER_NOT_FOUND"
	CategoryNotFound     = "CATEGORY_NOT_FOUND"

	// Application errors
	ValidationFailed     = "VALIDATION_FAILED"
	ExpenseCreationFailed = "EXPENSE_CREATION_FAILED"
	ExpenseUpdateFailed   = "EXPENSE_UPDATE_FAILED"
	ExpenseDeletionFailed = "EXPENSE_DELETION_FAILED"
)
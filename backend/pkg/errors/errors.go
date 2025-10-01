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
	InvalidUserName      = "INVALID_USER_NAME"
	InvalidUserEmail     = "INVALID_USER_EMAIL"
	InvalidCategoryID    = "INVALID_CATEGORY_ID"
	ExpenseNotFound      = "EXPENSE_NOT_FOUND"
	UserNotFound         = "USER_NOT_FOUND"
	CategoryNotFound     = "CATEGORY_NOT_FOUND"

	// Application errors
	ValidationFailed       = "VALIDATION_FAILED"
	EmailAlreadyExists     = "EMAIL_ALREADY_EXISTS"
	CategoryNameExists     = "CATEGORY_NAME_ALREADY_EXISTS"
	CategoryInUse          = "CATEGORY_IN_USE"
	ExpenseCreationFailed  = "EXPENSE_CREATION_FAILED"
	ExpenseUpdateFailed    = "EXPENSE_UPDATE_FAILED"
	ExpenseDeletionFailed  = "EXPENSE_DELETION_FAILED"
	UserCreationFailed     = "USER_CREATION_FAILED"
	UserUpdateFailed       = "USER_UPDATE_FAILED"
	UserDeleteFailed       = "USER_DELETE_FAILED"
	CategoryCreationFailed = "CATEGORY_CREATION_FAILED"
	CategoryUpdateFailed   = "CATEGORY_UPDATE_FAILED"
	CategoryDeleteFailed   = "CATEGORY_DELETE_FAILED"
)

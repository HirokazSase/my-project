package entity

import (
	"expense-management-system/internal/domain/valueobject"
	"expense-management-system/pkg/errors"
	"strings"
	"time"
)

// ExpenseStatus 経費の状態
type ExpenseStatus string

const (
	ExpenseStatusDraft     ExpenseStatus = "draft"     // 下書き
	ExpenseStatusSubmitted ExpenseStatus = "submitted" // 申請済み
	ExpenseStatusApproved  ExpenseStatus = "approved"  // 承認済み
	ExpenseStatusRejected  ExpenseStatus = "rejected"  // 却下
)

// Expense 経費エンティティ
type Expense struct {
	id          *valueobject.ExpenseID
	userID      *valueobject.UserID
	categoryID  *valueobject.CategoryID
	amount      *valueobject.Money
	title       string
	description string
	date        time.Time
	status      ExpenseStatus
	createdAt   time.Time
	updatedAt   time.Time
}

// NewExpense 新しいExpenseを作成
func NewExpense(userID *valueobject.UserID, categoryID *valueobject.CategoryID, amount *valueobject.Money, title, description string, date time.Time) (*Expense, error) {
	if userID == nil {
		return nil, errors.NewDomainError(errors.InvalidUserID, "ユーザーIDが必要です")
	}

	if categoryID == nil {
		return nil, errors.NewDomainError(errors.InvalidCategoryID, "カテゴリIDが必要です")
	}

	if amount == nil {
		return nil, errors.NewDomainError(errors.InvalidExpenseAmount, "金額が必要です")
	}

	if err := validateExpenseTitle(title); err != nil {
		return nil, err
	}

	if err := validateExpenseDescription(description); err != nil {
		return nil, err
	}

	if err := validateExpenseDate(date); err != nil {
		return nil, err
	}

	now := time.Now()
	return &Expense{
		id:          valueobject.GenerateExpenseID(),
		userID:      userID,
		categoryID:  categoryID,
		amount:      amount,
		title:       strings.TrimSpace(title),
		description: strings.TrimSpace(description),
		date:        date,
		status:      ExpenseStatusDraft,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

// ReconstructExpense 既存データからExpenseを再構築
func ReconstructExpense(
	id *valueobject.ExpenseID,
	userID *valueobject.UserID,
	categoryID *valueobject.CategoryID,
	amount *valueobject.Money,
	title, description string,
	date time.Time,
	status ExpenseStatus,
	createdAt, updatedAt time.Time,
) (*Expense, error) {
	if id == nil {
		return nil, errors.NewDomainError("INVALID_EXPENSE_ID", "経費IDが必要です")
	}

	if userID == nil {
		return nil, errors.NewDomainError(errors.InvalidUserID, "ユーザーIDが必要です")
	}

	if categoryID == nil {
		return nil, errors.NewDomainError(errors.InvalidCategoryID, "カテゴリIDが必要です")
	}

	if amount == nil {
		return nil, errors.NewDomainError(errors.InvalidExpenseAmount, "金額が必要です")
	}

	if err := validateExpenseTitle(title); err != nil {
		return nil, err
	}

	if err := validateExpenseDescription(description); err != nil {
		return nil, err
	}

	if err := validateExpenseDate(date); err != nil {
		return nil, err
	}

	if !isValidStatus(status) {
		return nil, errors.NewDomainError("INVALID_EXPENSE_STATUS", "無効な経費ステータスです")
	}

	return &Expense{
		id:          id,
		userID:      userID,
		categoryID:  categoryID,
		amount:      amount,
		title:       title,
		description: description,
		date:        date,
		status:      status,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}, nil
}

// ID IDを取得
func (e *Expense) ID() *valueobject.ExpenseID {
	return e.id
}

// UserID ユーザーIDを取得
func (e *Expense) UserID() *valueobject.UserID {
	return e.userID
}

// CategoryID カテゴリIDを取得
func (e *Expense) CategoryID() *valueobject.CategoryID {
	return e.categoryID
}

// Amount 金額を取得
func (e *Expense) Amount() *valueobject.Money {
	return e.amount
}

// Title タイトルを取得
func (e *Expense) Title() string {
	return e.title
}

// Description 説明を取得
func (e *Expense) Description() string {
	return e.description
}

// Date 日付を取得
func (e *Expense) Date() time.Time {
	return e.date
}

// Status ステータスを取得
func (e *Expense) Status() ExpenseStatus {
	return e.status
}

// CreatedAt 作成日時を取得
func (e *Expense) CreatedAt() time.Time {
	return e.createdAt
}

// UpdatedAt 更新日時を取得
func (e *Expense) UpdatedAt() time.Time {
	return e.updatedAt
}

// UpdateDetails 経費の詳細を更新
func (e *Expense) UpdateDetails(categoryID *valueobject.CategoryID, amount *valueobject.Money, title, description string, date time.Time) error {
	// 下書き状態でのみ更新可能
	if e.status != ExpenseStatusDraft {
		return errors.NewDomainError("EXPENSE_UPDATE_NOT_ALLOWED", "下書き状態の経費のみ更新できます")
	}

	if categoryID == nil {
		return errors.NewDomainError(errors.InvalidCategoryID, "カテゴリIDが必要です")
	}

	if amount == nil {
		return errors.NewDomainError(errors.InvalidExpenseAmount, "金額が必要です")
	}

	if err := validateExpenseTitle(title); err != nil {
		return err
	}

	if err := validateExpenseDescription(description); err != nil {
		return err
	}

	if err := validateExpenseDate(date); err != nil {
		return err
	}

	e.categoryID = categoryID
	e.amount = amount
	e.title = strings.TrimSpace(title)
	e.description = strings.TrimSpace(description)
	e.date = date
	e.updatedAt = time.Now()

	return nil
}

// Submit 経費を申請
func (e *Expense) Submit() error {
	if e.status != ExpenseStatusDraft {
		return errors.NewDomainError("EXPENSE_SUBMIT_NOT_ALLOWED", "下書き状態の経費のみ申請できます")
	}

	e.status = ExpenseStatusSubmitted
	e.updatedAt = time.Now()

	return nil
}

// Approve 経費を承認
func (e *Expense) Approve() error {
	if e.status != ExpenseStatusSubmitted {
		return errors.NewDomainError("EXPENSE_APPROVE_NOT_ALLOWED", "申請済み状態の経費のみ承認できます")
	}

	e.status = ExpenseStatusApproved
	e.updatedAt = time.Now()

	return nil
}

// Reject 経費を却下
func (e *Expense) Reject() error {
	if e.status != ExpenseStatusSubmitted {
		return errors.NewDomainError("EXPENSE_REJECT_NOT_ALLOWED", "申請済み状態の経費のみ却下できます")
	}

	e.status = ExpenseStatusRejected
	e.updatedAt = time.Now()

	return nil
}

// CanEdit 編集可能かどうか
func (e *Expense) CanEdit() bool {
	return e.status == ExpenseStatusDraft
}

// CanSubmit 申請可能かどうか
func (e *Expense) CanSubmit() bool {
	return e.status == ExpenseStatusDraft
}

// validateExpenseTitle 経費タイトルのバリデーション
func validateExpenseTitle(title string) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return errors.NewDomainError("INVALID_EXPENSE_TITLE", "経費タイトルは必須です")
	}

	if len(title) > 100 {
		return errors.NewDomainError("INVALID_EXPENSE_TITLE", "経費タイトルは100文字以内である必要があります")
	}

	return nil
}

// validateExpenseDescription 経費説明のバリデーション
func validateExpenseDescription(description string) error {
	if len(description) > 500 {
		return errors.NewDomainError("INVALID_EXPENSE_DESCRIPTION", "経費説明は500文字以内である必要があります")
	}

	return nil
}

// validateExpenseDate 経費日付のバリデーション
func validateExpenseDate(date time.Time) error {
	if date.IsZero() {
		return errors.NewDomainError("INVALID_EXPENSE_DATE", "経費日付が必要です")
	}

	// 未来の日付は許可しない（今日より後）
	now := time.Now()
	if date.After(time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())) {
		return errors.NewDomainError("INVALID_EXPENSE_DATE", "経費日付は未来の日付にできません")
	}

	// 1年以上前の日付は許可しない
	if date.Before(time.Now().AddDate(-1, 0, 0)) {
		return errors.NewDomainError("INVALID_EXPENSE_DATE", "経費日付は1年以内である必要があります")
	}

	return nil
}

// isValidStatus 有効なステータスかチェック
func isValidStatus(status ExpenseStatus) bool {
	switch status {
	case ExpenseStatusDraft, ExpenseStatusSubmitted, ExpenseStatusApproved, ExpenseStatusRejected:
		return true
	default:
		return false
	}
}

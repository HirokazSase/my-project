package valueobject

import (
	"expense-management-system/pkg/errors"
	"math"
)

// Money 金額を表すValue Object
type Money struct {
	amount   float64
	currency string
}

// NewMoney 新しいMoneyを作成
func NewMoney(amount float64, currency string) (*Money, error) {
	if amount < 0 {
		return nil, errors.NewDomainError(errors.InvalidExpenseAmount, "金額は負の値にできません")
	}

	if amount > math.MaxFloat64/100 {
		return nil, errors.NewDomainError(errors.InvalidExpenseAmount, "金額が大きすぎます")
	}

	if currency == "" {
		currency = "JPY" // デフォルトは日本円
	}

	// 小数点以下2桁に丸める
	roundedAmount := math.Round(amount*100) / 100

	return &Money{
		amount:   roundedAmount,
		currency: currency,
	}, nil
}

// Amount 金額を取得
func (m *Money) Amount() float64 {
	return m.amount
}

// Currency 通貨を取得
func (m *Money) Currency() string {
	return m.currency
}

// Add 金額を加算
func (m *Money) Add(other *Money) (*Money, error) {
	if m.currency != other.currency {
		return nil, errors.NewDomainError(errors.InvalidExpenseAmount, "異なる通貨の金額は加算できません")
	}

	return NewMoney(m.amount+other.amount, m.currency)
}

// Subtract 金額を減算
func (m *Money) Subtract(other *Money) (*Money, error) {
	if m.currency != other.currency {
		return nil, errors.NewDomainError(errors.InvalidExpenseAmount, "異なる通貨の金額は減算できません")
	}

	return NewMoney(m.amount-other.amount, m.currency)
}

// Multiply 金額を乗算
func (m *Money) Multiply(multiplier float64) (*Money, error) {
	return NewMoney(m.amount*multiplier, m.currency)
}

// Equals 等価性をチェック
func (m *Money) Equals(other *Money) bool {
	if other == nil {
		return false
	}
	return m.amount == other.amount && m.currency == other.currency
}

// IsGreaterThan 金額が他の金額より大きいかチェック
func (m *Money) IsGreaterThan(other *Money) bool {
	if m.currency != other.currency {
		return false
	}
	return m.amount > other.amount
}

// IsLessThan 金額が他の金額より小さいかチェック
func (m *Money) IsLessThan(other *Money) bool {
	if m.currency != other.currency {
		return false
	}
	return m.amount < other.amount
}

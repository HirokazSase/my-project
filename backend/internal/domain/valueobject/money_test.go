package valueobject

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMoney(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		currency string
		wantErr  bool
	}{
		{
			name:     "正常な金額とデフォルト通貨",
			amount:   100.50,
			currency: "",
			wantErr:  false,
		},
		{
			name:     "正常な金額と指定通貨",
			amount:   100.50,
			currency: "USD",
			wantErr:  false,
		},
		{
			name:     "ゼロ金額",
			amount:   0,
			currency: "JPY",
			wantErr:  false,
		},
		{
			name:     "負の金額",
			amount:   -100,
			currency: "JPY",
			wantErr:  true,
		},
		{
			name:     "非常に大きな金額",
			amount:   math.MaxFloat64,
			currency: "JPY",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			money, err := NewMoney(tt.amount, tt.currency)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, money)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, money)
				assert.Equal(t, tt.amount, money.Amount())

				expectedCurrency := tt.currency
				if expectedCurrency == "" {
					expectedCurrency = "JPY"
				}
				assert.Equal(t, expectedCurrency, money.Currency())
			}
		})
	}
}

func TestMoney_Add(t *testing.T) {
	money1, _ := NewMoney(100, "JPY")
	money2, _ := NewMoney(50, "JPY")
	money3, _ := NewMoney(75, "USD")

	t.Run("同じ通貨の加算", func(t *testing.T) {
		result, err := money1.Add(money2)
		require.NoError(t, err)
		assert.Equal(t, 150.0, result.Amount())
		assert.Equal(t, "JPY", result.Currency())
	})

	t.Run("異なる通貨の加算はエラー", func(t *testing.T) {
		result, err := money1.Add(money3)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestMoney_Subtract(t *testing.T) {
	money1, _ := NewMoney(100, "JPY")
	money2, _ := NewMoney(30, "JPY")
	money3, _ := NewMoney(75, "USD")

	t.Run("同じ通貨の減算", func(t *testing.T) {
		result, err := money1.Subtract(money2)
		require.NoError(t, err)
		assert.Equal(t, 70.0, result.Amount())
		assert.Equal(t, "JPY", result.Currency())
	})

	t.Run("異なる通貨の減算はエラー", func(t *testing.T) {
		result, err := money1.Subtract(money3)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("結果が負になる場合はエラー", func(t *testing.T) {
		result, err := money2.Subtract(money1)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestMoney_Multiply(t *testing.T) {
	money, _ := NewMoney(100, "JPY")

	t.Run("正の数での乗算", func(t *testing.T) {
		result, err := money.Multiply(1.5)
		require.NoError(t, err)
		assert.Equal(t, 150.0, result.Amount())
		assert.Equal(t, "JPY", result.Currency())
	})

	t.Run("ゼロでの乗算", func(t *testing.T) {
		result, err := money.Multiply(0)
		require.NoError(t, err)
		assert.Equal(t, 0.0, result.Amount())
	})

	t.Run("負の数での乗算はエラー", func(t *testing.T) {
		result, err := money.Multiply(-2)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestMoney_Equals(t *testing.T) {
	money1, _ := NewMoney(100, "JPY")
	money2, _ := NewMoney(100, "JPY")
	money3, _ := NewMoney(100, "USD")
	money4, _ := NewMoney(50, "JPY")

	t.Run("同じ金額と通貨", func(t *testing.T) {
		assert.True(t, money1.Equals(money2))
	})

	t.Run("同じ金額だが異なる通貨", func(t *testing.T) {
		assert.False(t, money1.Equals(money3))
	})

	t.Run("異なる金額", func(t *testing.T) {
		assert.False(t, money1.Equals(money4))
	})

	t.Run("nilとの比較", func(t *testing.T) {
		assert.False(t, money1.Equals(nil))
	})
}

func TestMoney_IsGreaterThan(t *testing.T) {
	money1, _ := NewMoney(100, "JPY")
	money2, _ := NewMoney(50, "JPY")
	money3, _ := NewMoney(100, "USD")

	t.Run("同じ通貨で大きい場合", func(t *testing.T) {
		assert.True(t, money1.IsGreaterThan(money2))
	})

	t.Run("同じ通貨で小さい場合", func(t *testing.T) {
		assert.False(t, money2.IsGreaterThan(money1))
	})

	t.Run("異なる通貨の比較", func(t *testing.T) {
		assert.False(t, money1.IsGreaterThan(money3))
	})
}

func TestMoney_IsLessThan(t *testing.T) {
	money1, _ := NewMoney(100, "JPY")
	money2, _ := NewMoney(50, "JPY")
	money3, _ := NewMoney(100, "USD")

	t.Run("同じ通貨で小さい場合", func(t *testing.T) {
		assert.True(t, money2.IsLessThan(money1))
	})

	t.Run("同じ通貨で大きい場合", func(t *testing.T) {
		assert.False(t, money1.IsLessThan(money2))
	})

	t.Run("異なる通貨の比較", func(t *testing.T) {
		assert.False(t, money1.IsLessThan(money3))
	})
}

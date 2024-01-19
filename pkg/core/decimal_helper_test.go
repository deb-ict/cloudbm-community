package core

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestTryGetDecimalFromString(t *testing.T) {
	tests := []struct {
		value    string
		expected decimal.Decimal
	}{
		{"10.5", decimal.NewFromFloat(10.5)},
		{"-5.25", decimal.NewFromFloat(-5.25)},
		{"0", decimal.NewFromFloat(0)},
		{"invalid", decimal.Zero},
	}

	for _, test := range tests {
		result := TryGetDecimalFromString(test.value)
		assert.Equal(t, 0, result.Cmp(test.expected), "TryGetDecimalFromString(%s) = %s, expected %s", test.value, result.String(), test.expected.String())
	}
}

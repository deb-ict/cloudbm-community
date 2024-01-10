package core

import "github.com/shopspring/decimal"

func TryGetDecimalFromString(value string) decimal.Decimal {
	d, err := decimal.NewFromString(value)
	if err != nil {
		return decimal.Zero
	}
	return d
}

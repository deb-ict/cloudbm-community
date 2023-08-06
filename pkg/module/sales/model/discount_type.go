package model

type DiscountType uint8

const (
	DiscountType_Undefined DiscountType = iota
	DiscountType_Rate
	DiscountType_Amount
)

func (t DiscountType) String() string {
	switch t {
	case DiscountType_Rate:
		return "Rate"
	case DiscountType_Amount:
		return "Amount"
	default:
		return "Undefined"
	}
}

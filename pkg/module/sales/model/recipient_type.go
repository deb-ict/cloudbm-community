package model

type RecipientType uint8

const (
	RecipientType_Undefined RecipientType = iota
	RecipientType_Company
	RecipientType_Private
)

func (t RecipientType) String() string {
	switch t {
	case RecipientType_Company:
		return "Company"
	case RecipientType_Private:
		return "Private"
	default:
		return "Undefined"
	}
}

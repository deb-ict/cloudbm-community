package model

type ProductVariantValue struct {
	AttributeId string
	ValueId     string
}

func (m *ProductVariantValue) Clone() *ProductVariantValue {
	if m == nil {
		return nil
	}
	return &ProductVariantValue{
		AttributeId: m.AttributeId,
		ValueId:     m.ValueId,
	}
}

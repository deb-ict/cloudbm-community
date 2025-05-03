package model

type ProductAttribute struct {
	AttributeId string
	ValueId     string
}

func (m *ProductAttribute) Clone() *ProductAttribute {
	if m == nil {
		return nil
	}
	return &ProductAttribute{
		AttributeId: m.AttributeId,
		ValueId:     m.ValueId,
	}
}

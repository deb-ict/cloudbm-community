package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
)

type ProductVariant struct {
	Id         string
	ProductId  string
	Attributes []*ProductVariantValue
	Details    *ProductDetail
}

func (m *ProductVariant) Normalize(normalizer core.StringNormalizer) {
	m.Details.Normalize(normalizer)
}

func (m *ProductVariant) UpdateModel(other *ProductVariant) {
	m.Attributes = make([]*ProductVariantValue, 0)
	m.Details = other.Details.Clone()
	for _, attribute := range other.Attributes {
		m.Attributes = append(m.Attributes, attribute.Clone())
	}
}

func (m *ProductVariant) GetTranslation(language string, defaultLanguage string) *ProductTranslation {
	return m.Details.GetTranslation(language, defaultLanguage)
}

func (m *ProductVariant) TryGetTranslation(language string) (*ProductTranslation, error) {
	return m.Details.TryGetTranslation(language)
}

func (m *ProductVariant) IsTransient() bool {
	return m.Id == ""
}

func (m *ProductVariant) Clone() *ProductVariant {
	if m == nil {
		return nil
	}
	model := &ProductVariant{
		Id:         m.Id,
		ProductId:  m.ProductId,
		Details:    m.Details.Clone(),
		Attributes: make([]*ProductVariantValue, 0),
	}
	for _, attribute := range m.Attributes {
		model.Attributes = append(model.Attributes, attribute.Clone())
	}
	return model
}

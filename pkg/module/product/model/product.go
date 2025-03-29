package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
)

type Product struct {
	Id           string
	CategoryIds  []string
	AttributeIds []string
	Details      *ProductDetail
	Variants     []*ProductVariant
}

func (m *Product) Normalize(normalizer core.StringNormalizer) {
	m.Details.Normalize(normalizer)
	for _, variant := range m.Variants {
		variant.Normalize(normalizer)
	}
}

func (m *Product) UpdateModel(other *Product) {
	m.CategoryIds = make([]string, 0)
	m.CategoryIds = append(m.CategoryIds, other.CategoryIds...)
	m.AttributeIds = make([]string, 0)
	m.AttributeIds = append(m.AttributeIds, other.AttributeIds...)
	m.Variants = make([]*ProductVariant, 0)
	m.Details = other.Details.Clone()
	for _, variant := range other.Variants {
		m.Variants = append(m.Variants, variant.Clone())
	}
}

func (m *Product) GetTranslation(language string, defaultLanguage string) *ProductTranslation {
	return m.Details.GetTranslation(language, defaultLanguage)
}

func (m *Product) TryGetTranslation(language string) (*ProductTranslation, error) {
	return m.Details.TryGetTranslation(language)
}

func (m *Product) IsTransient() bool {
	return m.Id == ""
}

func (m *Product) Clone() *Product {
	if m == nil {
		return nil
	}
	model := &Product{
		Id:           m.Id,
		CategoryIds:  make([]string, 0),
		AttributeIds: make([]string, 0),
		Variants:     make([]*ProductVariant, 0),
		Details:      m.Details.Clone(),
	}
	model.CategoryIds = append(model.CategoryIds, m.CategoryIds...)
	model.AttributeIds = append(model.AttributeIds, m.AttributeIds...)
	for _, variant := range m.Variants {
		model.Variants = append(model.Variants, variant.Clone())
	}
	return model
}

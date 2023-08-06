package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	metadata "github.com/deb-ict/cloudbm-community/pkg/module/metadata/model"
	"github.com/shopspring/decimal"
)

type Product struct {
	Id           string
	CategoryIds  []string
	Translations []*ProductTranslation
	ThumbnailId  string
	ThumbnailUri string
	Price        decimal.Decimal
	TaxProfile   *metadata.TaxProfile
	IsEnabled    bool
}

type ProductTranslation struct {
	Language    string
	Name        string
	Slug        string
	Summary     string
	Description string
}

type ProductFilter struct {
	CategoryId string
}

func (m *Product) GetTranslation(language string, defaultLanguage string) *ProductTranslation {
	if len(m.Translations) == 0 {
		return &ProductTranslation{}
	}

	translation, err := m.TryGetTranslation(language)
	if err == core.ErrTranslationNotFound && language != defaultLanguage {
		translation, err = m.TryGetTranslation(defaultLanguage)
	}
	if err == core.ErrTranslationNotFound {
		translation = m.Translations[0]
	}

	return translation
}

func (m *Product) TryGetTranslation(language string) (*ProductTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}

	return nil, core.ErrTranslationNotFound
}

func (m *Product) IsTransient() bool {
	return m.Id == ""
}

func (f *ProductFilter) HasFilter() bool {
	if f.CategoryId != "" {
		return true
	}
	return false
}

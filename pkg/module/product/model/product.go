package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type Product struct {
	Id              string
	CategoryIds     []string
	Translations    []*ProductTranslation
	ThumbnailId     string
	ThumbnailUri    string
	Price           int64
	PriceMultiplier int
	IsEnabled       bool
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

func (m *Product) GetTranslation(language string) *ProductTranslation {
	if len(m.Translations) == 0 {
		return &ProductTranslation{}
	}

	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t
		}
	}

	return m.Translations[0]
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

package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/gosimple/slug"
	"github.com/shopspring/decimal"
)

type Product struct {
	Id           string
	CategoryIds  []string
	Translations []*ProductTranslation
	ThumbnailId  string
	ThumbnailUri string
	Gtin         string
	RegularPrice decimal.Decimal
	IsEnabled    bool
}

type ProductTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Slug           string
	Summary        string
	Description    string
}

type ProductFilter struct {
	Language   string
	Name       string
	CategoryId string
}

func (m *Product) Normalize(normalizer core.StringNormalizer) {
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
		if translation.Slug == "" {
			translation.Slug = slug.MakeLang(translation.NormalizedName, translation.Language)
		}
		translation.Slug = normalizer.NormalizeString(translation.Slug)
	}
}

func (m *Product) UpdateModel(other *Product) {
	m.CategoryIds = make([]string, 0)
	m.CategoryIds = append(m.CategoryIds, other.CategoryIds...)
	m.Translations = make([]*ProductTranslation, 0)
	m.Translations = append(m.Translations, other.Translations...)
	m.ThumbnailId = other.ThumbnailId
	m.ThumbnailUri = other.ThumbnailUri
	m.Price = other.Price
	m.IsEnabled = other.IsEnabled
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

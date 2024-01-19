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
	Ean          string
	Sku          string
	Mpn          string
	RegularPrice decimal.Decimal
	SalesPrice   decimal.Decimal
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
	m.ThumbnailId = other.ThumbnailId
	m.Ean = other.Ean
	m.Sku = other.Sku
	m.Mpn = other.Mpn
	m.RegularPrice = other.RegularPrice
	m.SalesPrice = other.SalesPrice
	m.IsEnabled = other.IsEnabled
	for _, translation := range other.Translations {
		m.Translations = append(m.Translations, translation.Clone())
	}
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

func (m *Product) Clone() *Product {
	if m == nil {
		return nil
	}
	model := &Product{
		Id:           m.Id,
		CategoryIds:  make([]string, 0),
		Translations: make([]*ProductTranslation, 0),
		ThumbnailId:  m.ThumbnailId,
		Ean:          m.Ean,
		Sku:          m.Sku,
		Mpn:          m.Mpn,
		RegularPrice: m.RegularPrice,
		SalesPrice:   m.SalesPrice,
		IsEnabled:    m.IsEnabled,
	}
	model.CategoryIds = append(model.CategoryIds, m.CategoryIds...)
	for _, translation := range m.Translations {
		model.Translations = append(model.Translations, translation.Clone())
	}
	return model
}

func (m *ProductTranslation) Clone() *ProductTranslation {
	if m == nil {
		return nil
	}
	return &ProductTranslation{
		Language:       m.Language,
		Name:           m.Name,
		NormalizedName: m.NormalizedName,
		Slug:           m.Slug,
		Summary:        m.Summary,
		Description:    m.Description,
	}
}

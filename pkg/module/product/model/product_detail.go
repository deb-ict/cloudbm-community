package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/gosimple/slug"
	"github.com/shopspring/decimal"
)

type ProductDetail struct {
	Translations []*ProductTranslation
	ThumbnailId  string
	Gtin         string
	Sku          string
	Mpn          string
	RegularPrice decimal.Decimal
	SalesPrice   decimal.Decimal
	IsEnabled    bool
}

func (m *ProductDetail) Normalize(normalizer core.StringNormalizer) {
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
		if translation.Slug == "" {
			translation.Slug = slug.MakeLang(translation.NormalizedName, translation.Language)
		}
		translation.Slug = normalizer.NormalizeString(translation.Slug)
	}
}

func (m *ProductDetail) UpdateModel(other *ProductDetail) {
	m.Translations = make([]*ProductTranslation, 0)
	m.ThumbnailId = other.ThumbnailId
	m.Gtin = other.Gtin
	m.Sku = other.Sku
	m.Mpn = other.Mpn
	m.RegularPrice = other.RegularPrice
	m.SalesPrice = other.SalesPrice
	m.IsEnabled = other.IsEnabled
	for _, translation := range other.Translations {
		m.Translations = append(m.Translations, translation.Clone())
	}
}

func (m *ProductDetail) GetTranslation(language string, defaultLanguage string) *ProductTranslation {
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

func (m *ProductDetail) TryGetTranslation(language string) (*ProductTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *ProductDetail) Clone() *ProductDetail {
	return &ProductDetail{
		ThumbnailId:  m.ThumbnailId,
		Gtin:         m.Gtin,
		Sku:          m.Sku,
		Mpn:          m.Mpn,
		RegularPrice: m.RegularPrice,
		SalesPrice:   m.SalesPrice,
		IsEnabled:    m.IsEnabled,
	}
}

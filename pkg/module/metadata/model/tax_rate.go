package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/shopspring/decimal"
)

type TaxRate struct {
	Id           string
	Key          string
	Translations []*TaxRateTranslation
	Rate         decimal.Decimal
	IsEnabled    bool
}

type TaxRateTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Description    string
}

type TaxRateFilter struct {
	Language string
	Name     string
}

func (m *TaxRate) Normalize(normalizer core.StringNormalizer) {
	m.Key = normalizer.NormalizeString(m.Key)
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
	}
}

func (m *TaxRate) UpdateModel(other *TaxRate) {
	m.Translations = make([]*TaxRateTranslation, 0)
	for _, translation := range other.Translations {
		m.Translations = append(m.Translations, translation.Clone())
	}
	m.Rate = other.Rate
}

func (m *TaxRate) GetTranslation(language string, defaultLanguage string) *TaxRateTranslation {
	if len(m.Translations) == 0 {
		return &TaxRateTranslation{}
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

func (m *TaxRate) TryGetTranslation(language string) (*TaxRateTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *TaxRate) IsTransient() bool {
	return m.Id == ""
}

func (m *TaxRate) Clone() *TaxRate {
	if m == nil {
		return nil
	}
	model := &TaxRate{
		Id:           m.Id,
		Key:          m.Key,
		Translations: make([]*TaxRateTranslation, 0),
		Rate:         m.Rate,
		IsEnabled:    m.IsEnabled,
	}
	for _, translation := range m.Translations {
		model.Translations = append(model.Translations, translation.Clone())
	}
	return model
}

func (m *TaxRateTranslation) Clone() *TaxRateTranslation {
	if m == nil {
		return nil
	}
	return &TaxRateTranslation{
		Language:       m.Language,
		Name:           m.Name,
		NormalizedName: m.NormalizedName,
		Description:    m.Description,
	}
}

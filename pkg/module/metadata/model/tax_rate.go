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
	Language    string
	Name        string
	Description string
}

type TaxRateFilter struct {
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

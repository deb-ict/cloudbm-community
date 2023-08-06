package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/shopspring/decimal"
)

type TaxProfile struct {
	Id           string
	Key          string
	Translations []*TaxProfileTranslation
	Rate         decimal.Decimal
}

type TaxProfileTranslation struct {
	Language    string
	Name        string
	Description string
}

type TaxProfileFilter struct {
}

func (m *TaxProfile) GetTranslation(language string, defaultLanguage string) *TaxProfileTranslation {
	if len(m.Translations) == 0 {
		return &TaxProfileTranslation{}
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

func (m *TaxProfile) TryGetTranslation(language string) (*TaxProfileTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *TaxProfile) IsTransient() bool {
	return m.Id == ""
}

func (f *TaxProfileFilter) HasFilter() bool {
	return false
}

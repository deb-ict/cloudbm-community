package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type Unit struct {
	Id           string
	Key          string
	Translations []*UnitTranslation
	IsSystem     bool
	IsEnabled    bool
}

type UnitTranslation struct {
	Language    string
	Name        string
	Description string
}

type UnitFilter struct {
}

func (m *Unit) GetTranslation(language string, defaultLanguage string) *UnitTranslation {
	if len(m.Translations) == 0 {
		return &UnitTranslation{}
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

func (m *Unit) TryGetTranslation(language string) (*UnitTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *Unit) IsTransient() bool {
	return m.Id == ""
}

package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type Industry struct {
	Id           string
	Key          string
	Translations []*IndustryTranslation
	IsSystem     bool
}

type IndustryTranslation struct {
	Language    string
	Name        string
	Description string
}

type IndustryFilter struct {
	Name string
}

func (m *Industry) GetTranslation(language string, defaultLanguage string) *IndustryTranslation {
	if len(m.Translations) == 0 {
		return &IndustryTranslation{}
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

func (m *Industry) TryGetTranslation(language string) (*IndustryTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *Industry) IsTransient() bool {
	return m.Id == ""
}

func (f *IndustryFilter) HasFilter() bool {
	return false
}

package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type ContactTitle struct {
	Id           string
	Key          string
	Translations []*ContactTitleTranslation
	IsSystem     bool
}

type ContactTitleTranslation struct {
	Language    string
	Name        string
	Description string
}

type ContactTitleFilter struct {
	Language string
	Name     string
}

func (m *ContactTitle) GetTranslation(language string, defaultLanguage string) *ContactTitleTranslation {
	if len(m.Translations) == 0 {
		return &ContactTitleTranslation{}
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

func (m *ContactTitle) TryGetTranslation(language string) (*ContactTitleTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *ContactTitle) IsTransient() bool {
	return m.Id == ""
}

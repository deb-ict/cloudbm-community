package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type EmailType struct {
	Id           string
	Key          string
	Translations []*EmailTypeTranslation
	IsDefault    bool
	IsSystem     bool
}

type EmailTypeTranslation struct {
	Language    string
	Name        string
	Description string
}

type EmailTypeFilter struct {
	Language string
	Name     string
}

func (m *EmailType) GetTranslation(language string, defaultLanguage string) *EmailTypeTranslation {
	if len(m.Translations) == 0 {
		return &EmailTypeTranslation{}
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

func (m *EmailType) TryGetTranslation(language string) (*EmailTypeTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *EmailType) IsTransient() bool {
	return m.Id == ""
}

package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type PhoneType struct {
	Id           string
	Key          string
	Translations []*PhoneTypeTranslation
	IsDefault    bool
	IsSystem     bool
}

type PhoneTypeTranslation struct {
	Language    string
	Name        string
	Description string
}

type PhoneTypeFilter struct {
	Name string
}

func (m *PhoneType) GetTranslation(language string, defaultLanguage string) *PhoneTypeTranslation {
	if len(m.Translations) == 0 {
		return &PhoneTypeTranslation{}
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

func (m *PhoneType) TryGetTranslation(language string) (*PhoneTypeTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *PhoneType) IsTransient() bool {
	return m.Id == ""
}

func (m *PhoneType) CanDelete() bool {
	return !m.IsDefault && !m.IsSystem
}

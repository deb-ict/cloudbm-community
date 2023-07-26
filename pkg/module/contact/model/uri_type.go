package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type UriType struct {
	Id           string
	Translations []*UriTypeTranslation
	IsDefault    bool
	IsSystem     bool
}

type UriTypeTranslation struct {
	Language    string
	Name        string
	Description string
}

type UriTypeFilter struct {
	Name string
}

func (m *UriType) GetTranslation(language string, defaultLanguage string) *UriTypeTranslation {
	if len(m.Translations) == 0 {
		return &UriTypeTranslation{}
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

func (m *UriType) TryGetTranslation(language string) (*UriTypeTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *UriType) IsTransient() bool {
	return m.Id == ""
}

func (m *UriType) CanDelete() bool {
	return !m.IsDefault && !m.IsSystem
}

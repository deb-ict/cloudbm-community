package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type AddressType struct {
	Id           string
	Key          string
	Translations []*AddressTypeTranslation
	IsDefault    bool
	IsSystem     bool
}

type AddressTypeTranslation struct {
	Language    string
	Name        string
	Description string
}

type AddressTypeFilter struct {
	Name string
}

func (m *AddressType) GetTranslation(language string, defaultLanguage string) *AddressTypeTranslation {
	if len(m.Translations) == 0 {
		return &AddressTypeTranslation{}
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

func (m *AddressType) TryGetTranslation(language string) (*AddressTypeTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *AddressType) IsTransient() bool {
	return m.Id == ""
}

func (m *AddressType) CanDelete() bool {
	return !m.IsDefault && !m.IsSystem
}

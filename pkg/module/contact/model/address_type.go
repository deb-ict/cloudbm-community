package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type AddressType struct {
	Id           string
	Translations []AddressTypeTranslation
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

func (m *AddressType) GetTranslation(language string) AddressTypeTranslation {
	if len(m.Translations) == 0 {
		return AddressTypeTranslation{}
	}

	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t
		}
	}

	return m.Translations[0]
}

func (m *AddressType) IsTransient() bool {
	return m.Id == ""
}

func (m *AddressType) CanDelete() bool {
	return !m.IsDefault && !m.IsSystem
}

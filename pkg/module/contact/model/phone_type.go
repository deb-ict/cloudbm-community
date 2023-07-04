package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type PhoneType struct {
	Id           string
	Translations []PhoneTypeTranslation
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

func (m *PhoneType) GetTranslation(language string) PhoneTypeTranslation {
	if len(m.Translations) == 0 {
		return PhoneTypeTranslation{}
	}

	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t
		}
	}

	return m.Translations[0]
}

func (m *PhoneType) IsTransient() bool {
	return m.Id == ""
}

func (m *PhoneType) CanDelete() bool {
	return !m.IsDefault && !m.IsSystem
}

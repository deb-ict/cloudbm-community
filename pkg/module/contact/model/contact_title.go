package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type ContactTitle struct {
	Id           string
	Translations []ContactTitleTranslation
	IsDefault    bool
	IsSystem     bool
}

type ContactTitleTranslation struct {
	Language    string
	Name        string
	Description string
}

type ContactTitleFilter struct {
	Name string
}

func (m *ContactTitle) GetTranslation(language string) ContactTitleTranslation {
	if len(m.Translations) == 0 {
		return ContactTitleTranslation{}
	}

	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t
		}
	}

	return m.Translations[0]
}

func (m *ContactTitle) IsTransient() bool {
	return m.Id == ""
}

func (m *ContactTitle) CanDelete() bool {
	return !m.IsDefault && !m.IsSystem
}

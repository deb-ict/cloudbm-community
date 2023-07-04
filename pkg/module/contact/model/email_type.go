package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type EmailType struct {
	Id           string
	Translations []EmailTypeTranslation
	IsDefault    bool
	IsSystem     bool
}

type EmailTypeTranslation struct {
	Language    string
	Name        string
	Description string
}

type EmailTypeFilter struct {
	Name string
}

func (m *EmailType) GetTranslation(language string) EmailTypeTranslation {
	if len(m.Translations) == 0 {
		return EmailTypeTranslation{}
	}

	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t
		}
	}

	return m.Translations[0]
}

func (m *EmailType) IsTransient() bool {
	return m.Id == ""
}

func (m *EmailType) CanDelete() bool {
	return !m.IsDefault && !m.IsSystem
}

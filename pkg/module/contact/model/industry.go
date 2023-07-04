package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type Industry struct {
	Id           string
	Translations []IndustryTranslation
	IsDefault    bool
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

func (m *Industry) GetTranslation(language string) IndustryTranslation {
	if len(m.Translations) == 0 {
		return IndustryTranslation{}
	}

	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t
		}
	}

	return m.Translations[0]
}

func (m *Industry) IsTransient() bool {
	return m.Id == ""
}

func (m *Industry) CanDelete() bool {
	return !m.IsDefault && !m.IsSystem
}

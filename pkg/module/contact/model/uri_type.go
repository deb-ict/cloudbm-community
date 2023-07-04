package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type UriType struct {
	Id           string
	Translations []UriTypeTranslation
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

func (m *UriType) GetTranslation(language string) UriTypeTranslation {
	if len(m.Translations) == 0 {
		return UriTypeTranslation{}
	}

	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t
		}
	}

	return m.Translations[0]
}

func (m *UriType) IsTransient() bool {
	return m.Id == ""
}

func (m *UriType) CanDelete() bool {
	return !m.IsDefault && !m.IsSystem
}

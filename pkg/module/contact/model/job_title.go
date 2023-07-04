package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type JobTitle struct {
	Id           string
	Translations []JobTitleTranslation
	IsDefault    bool
	IsSystem     bool
}

type JobTitleTranslation struct {
	Language    string
	Name        string
	Description string
}

type JobTitleFilter struct {
	Name string
}

func (m *JobTitle) GetTranslation(language string) JobTitleTranslation {
	if len(m.Translations) == 0 {
		return JobTitleTranslation{}
	}

	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t
		}
	}

	return m.Translations[0]
}

func (m *JobTitle) IsTransient() bool {
	return m.Id == ""
}

func (m *JobTitle) CanDelete() bool {
	return !m.IsDefault && !m.IsSystem
}

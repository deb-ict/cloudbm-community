package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type JobTitle struct {
	Id           string
	Key          string
	Translations []*JobTitleTranslation
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

func (m *JobTitle) GetTranslation(language string, defaultLanguage string) *JobTitleTranslation {
	if len(m.Translations) == 0 {
		return &JobTitleTranslation{}
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

func (m *JobTitle) TryGetTranslation(language string) (*JobTitleTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *JobTitle) IsTransient() bool {
	return m.Id == ""
}

func (m *JobTitle) CanDelete() bool {
	return !m.IsDefault && !m.IsSystem
}

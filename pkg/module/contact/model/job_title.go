package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type JobTitle struct {
	Id           string
	Key          string
	Translations []*JobTitleTranslation
	IsSystem     bool
}

type JobTitleTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Description    string
}

type JobTitleFilter struct {
	Language string
	Name     string
}

func (m *JobTitle) Normalize(normalizer core.StringNormalizer) {
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
	}
}

func (m *JobTitle) UpdateModel(other *JobTitle) {
	m.Translations = make([]*JobTitleTranslation, 0)
	for _, translation := range other.Translations {
		m.Translations = append(m.Translations, translation.Clone())
	}
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

func (m *JobTitle) Clone() *JobTitle {
	if m == nil {
		return nil
	}
	model := &JobTitle{
		Id:           m.Id,
		Key:          m.Key,
		Translations: make([]*JobTitleTranslation, 0),
		IsSystem:     m.IsSystem,
	}
	for _, translation := range m.Translations {
		model.Translations = append(model.Translations, translation.Clone())
	}
	return model
}

func (m *JobTitleTranslation) Clone() *JobTitleTranslation {
	if m == nil {
		return nil
	}
	return &JobTitleTranslation{
		Language:       m.Language,
		Name:           m.Name,
		NormalizedName: m.NormalizedName,
		Description:    m.Description,
	}
}

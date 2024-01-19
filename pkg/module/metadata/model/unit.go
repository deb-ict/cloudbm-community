package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type Unit struct {
	Id           string
	Key          string
	Translations []*UnitTranslation
	IsEnabled    bool
}

type UnitTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Description    string
}

type UnitFilter struct {
	Language string
	Name     string
}

func (m *Unit) Normalize(normalizer core.StringNormalizer) {
	m.Key = normalizer.NormalizeString(m.Key)
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
	}
}

func (m *Unit) UpdateModel(other *Unit) {
	m.Translations = make([]*UnitTranslation, 0)
	for _, translation := range other.Translations {
		m.Translations = append(m.Translations, translation.Clone())
	}
}

func (m *Unit) GetTranslation(language string, defaultLanguage string) *UnitTranslation {
	if len(m.Translations) == 0 {
		return &UnitTranslation{}
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

func (m *Unit) TryGetTranslation(language string) (*UnitTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *Unit) IsTransient() bool {
	return m.Id == ""
}

func (m *Unit) Clone() *Unit {
	if m == nil {
		return nil
	}
	model := &Unit{
		Id:           m.Id,
		Key:          m.Key,
		Translations: make([]*UnitTranslation, 0),
		IsEnabled:    m.IsEnabled,
	}
	for _, translation := range m.Translations {
		model.Translations = append(model.Translations, translation.Clone())
	}
	return model
}

func (m *UnitTranslation) Clone() *UnitTranslation {
	if m == nil {
		return nil
	}
	return &UnitTranslation{
		Language:       m.Language,
		Name:           m.Name,
		NormalizedName: m.NormalizedName,
		Description:    m.Description,
	}
}

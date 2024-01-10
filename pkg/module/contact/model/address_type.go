package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type AddressType struct {
	Id           string
	Key          string
	Translations []*AddressTypeTranslation
	IsDefault    bool
	IsSystem     bool
}

type AddressTypeTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Description    string
}

type AddressTypeFilter struct {
	Language string
	Name     string
}

func (m *AddressType) Normalize(normalizer core.StringNormalizer) {
	m.Key = normalizer.NormalizeString(m.Key)
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
	}
}

func (m *AddressType) UpdateModel(other *AddressType) {
	m.Translations = make([]*AddressTypeTranslation, 0)
	m.IsDefault = other.IsDefault
	for _, translation := range other.Translations {
		m.Translations = append(m.Translations, translation.Clone())
	}
}

func (m *AddressType) GetTranslation(language string, defaultLanguage string) *AddressTypeTranslation {
	if len(m.Translations) == 0 {
		return &AddressTypeTranslation{}
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

func (m *AddressType) TryGetTranslation(language string) (*AddressTypeTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *AddressType) IsTransient() bool {
	return m.Id == ""
}

func (m *AddressType) Clone() *AddressType {
	if m == nil {
		return nil
	}
	model := &AddressType{
		Id:           m.Id,
		Key:          m.Key,
		Translations: make([]*AddressTypeTranslation, 0),
		IsDefault:    m.IsDefault,
		IsSystem:     m.IsSystem,
	}
	for _, translation := range m.Translations {
		model.Translations = append(model.Translations, translation.Clone())
	}
	return model
}

func (m *AddressTypeTranslation) Clone() *AddressTypeTranslation {
	if m == nil {
		return nil
	}
	return &AddressTypeTranslation{
		Language:       m.Language,
		Name:           m.Name,
		NormalizedName: m.NormalizedName,
		Description:    m.Description,
	}
}

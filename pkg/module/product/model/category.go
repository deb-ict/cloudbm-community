package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/gosimple/slug"
)

type Category struct {
	Id           string
	ParentId     string
	Translations []*CategoryTranslation
	ThumbnailId  string
	ThumbnailUri string
	IsEnabled    bool
}

type CategoryTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Slug           string
	Summary        string
	Description    string
}

type CategoryFilter struct {
	Language  string
	Name      string
	ParentId  string
	AllLevels bool
}

func (m *Category) Normalize(normalizer core.StringNormalizer) {
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
		if translation.Slug == "" {
			translation.Slug = slug.MakeLang(translation.NormalizedName, translation.Language)
		}
		translation.Slug = normalizer.NormalizeString(translation.Slug)
	}
}

func (m *Category) UpdateModel(other *Category) {
	m.ParentId = other.ParentId
	m.Translations = make([]*CategoryTranslation, 0)
	m.Translations = append(m.Translations, other.Translations...)
	m.ThumbnailId = other.ThumbnailId
	m.ThumbnailUri = other.ThumbnailUri
	m.IsEnabled = other.IsEnabled
}

func (m *Category) GetTranslation(language string, defaultLanguage string) *CategoryTranslation {
	if len(m.Translations) == 0 {
		return &CategoryTranslation{}
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

func (m *Category) TryGetTranslation(language string) (*CategoryTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *Category) IsTransient() bool {
	return m.Id == ""
}

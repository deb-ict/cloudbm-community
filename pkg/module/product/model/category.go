package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type Category struct {
	Id           string
	ParentId     string
	Translations []*CategoryTranslation
	ThumbnailId  string
	ThumbnailUri string
	SortOrder    int64
	IsEnabled    bool
}

type CategoryTranslation struct {
	Language    string
	Name        string
	Slug        string
	Summary     string
	Description string
}

type CategoryFilter struct {
	ParentId string
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

func (f *CategoryFilter) HasFilter() bool {
	if f.ParentId != "" {
		return true
	}
	return false
}

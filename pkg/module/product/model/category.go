package model

import (
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

func (m *Category) GetTranslation(language string) *CategoryTranslation {
	if len(m.Translations) == 0 {
		return &CategoryTranslation{}
	}

	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t
		}
	}

	return m.Translations[0]
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

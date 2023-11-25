package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type Media struct {
	Id            string
	CategoryId    string
	Translations  []*MediaTranslation
	StorageFolder string
	FileName      string
	FileType      string
	FileSize      int64
	Width         int
	Height        int
}

type MediaTranslation struct {
	Language    string
	Name        string
	Slug        string
	Summary     string
	Description string
}

type MediaFilter struct {
	CategoryId string
}

func (m *Media) GetTranslation(language string, defaultLanguage string) *MediaTranslation {
	if len(m.Translations) == 0 {
		return &MediaTranslation{}
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

func (m *Media) TryGetTranslation(language string) (*MediaTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *Media) IsTransient() bool {
	return m.Id == ""
}

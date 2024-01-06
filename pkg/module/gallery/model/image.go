package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/gosimple/slug"
)

type Image struct {
	Id               string
	Translations     []*ImageTranslation
	StorageFolder    string
	FileName         string
	OriginalFileName string
	FileSize         int64
	MimeType         string
	Width            int
	Height           int
}

type ImageTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Slug           string
	Summary        string
	Description    string
}

type ImageFilter struct {
	Language string
	Name     string
	MimeType string
}

func (m *Image) Normalize(normalizer core.StringNormalizer) {
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
		if translation.Slug == "" {
			translation.Slug = slug.MakeLang(translation.NormalizedName, translation.Language)
		}
		translation.Slug = normalizer.NormalizeString(translation.Slug)
	}
}

func (m *Image) UpdateModel(other *Image) {
	m.Translations = make([]*ImageTranslation, 0)
	m.Translations = append(m.Translations, other.Translations...)
	m.OriginalFileName = other.OriginalFileName
}

func (m *Image) GetTranslation(language string, defaultLanguage string) *ImageTranslation {
	if len(m.Translations) == 0 {
		return &ImageTranslation{}
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

func (m *Image) TryGetTranslation(language string) (*ImageTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *Image) IsTransient() bool {
	return m.Id == ""
}

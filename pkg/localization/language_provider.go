package localization

import (
	"context"
)

const (
	DEFAULT_LANGUAGE_ID string = "en"
)

type LanguageProvider interface {
	UserLanguage(ctx context.Context) string
	DefaultLanguage(ctx context.Context) string
	SupportedLanguages(ctx context.Context) []string
	IsDefaultLanguage(ctx context.Context, language string) bool
	IsSupportedLanguage(ctx context.Context, language string) bool
}

type defaultLanguageProvider struct {
}

func NewDefaultLanguageProvider() LanguageProvider {
	return &defaultLanguageProvider{}
}

func (p *defaultLanguageProvider) UserLanguage(ctx context.Context) string {
	return NormalizeLanguage(DEFAULT_LANGUAGE_ID)
}

func (p *defaultLanguageProvider) DefaultLanguage(ctx context.Context) string {
	return NormalizeLanguage(DEFAULT_LANGUAGE_ID)
}

func (p *defaultLanguageProvider) SupportedLanguages(ctx context.Context) []string {
	return []string{
		NormalizeLanguage(DEFAULT_LANGUAGE_ID),
	}
}

func (p *defaultLanguageProvider) IsDefaultLanguage(ctx context.Context, language string) bool {
	return NormalizeLanguage(language) == NormalizeLanguage(DEFAULT_LANGUAGE_ID)
}

func (p *defaultLanguageProvider) IsSupportedLanguage(ctx context.Context, language string) bool {
	normalizedLanguage := NormalizeLanguage(language)
	for _, supportedLanguage := range p.SupportedLanguages(ctx) {
		if supportedLanguage == normalizedLanguage {
			return true
		}
	}
	return false
}

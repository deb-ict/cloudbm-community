package localization

import "context"

const (
	DEFAULT_LANGUAGE_ID string = "EN"
)

type LanguageProvider interface {
	UserLanguage(ctx context.Context) string
	DefaultLanguage(ctx context.Context) string
}

type defaultLanguageProvider struct {
}

func NewDefaultLanguageProvider() LanguageProvider {
	return &defaultLanguageProvider{}
}

func (p *defaultLanguageProvider) UserLanguage(ctx context.Context) string {
	return DEFAULT_LANGUAGE_ID
}

func (p *defaultLanguageProvider) DefaultLanguage(ctx context.Context) string {
	return DEFAULT_LANGUAGE_ID
}

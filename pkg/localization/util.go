package localization

import (
	"net/http"
	"strings"
)

func NormalizeLanguage(language string) string {
	return strings.ToLower(language)
}

func GetHttpRequestLanguage(r *http.Request, l LanguageProvider) string {
	language := r.URL.Query().Get("language")
	if language == "" {
		language = l.UserLanguage(r.Context())
	}
	return language
}

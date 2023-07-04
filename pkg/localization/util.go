package localization

import (
	"strings"
)

func NormalizeLanguage(language string) string {
	return strings.ToLower(language)
}

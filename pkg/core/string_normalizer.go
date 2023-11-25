package core

import "strings"

type StringNormalizer interface {
	NormalizeString(value string) string
}

type defaultStringNormalizer struct {
}

func DefaultStringNormalizer() StringNormalizer {
	return &defaultStringNormalizer{}
}

func (n *defaultStringNormalizer) NormalizeString(value string) string {
	return strings.ToLower(value)
}

package module

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type ServiceOptions struct {
	StringNormalizer core.StringNormalizer
	FeatureProvider  core.FeatureProvider
	LanguageProvider localization.LanguageProvider
}

type Service interface {
	GetStringNormalizer() core.StringNormalizer
	GetFeatureProvider() core.FeatureProvider
	GetLanguageProvider() localization.LanguageProvider
}

func (opt *ServiceOptions) EnsureDefaults() {
	if opt.StringNormalizer == nil {
		opt.StringNormalizer = core.NewDefaultStringNormalizer()
	}
	if opt.FeatureProvider == nil {
		opt.FeatureProvider = core.NewDefaultFeatureProvider()
	}
	if opt.LanguageProvider == nil {
		opt.LanguageProvider = localization.NewDefaultLanguageProvider()
	}
}

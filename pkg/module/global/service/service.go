package service

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/global"
)

type ServiceOptions struct {
	StringNormalizer core.StringNormalizer
	FeatureProvider  core.FeatureProvider
	LanguageProvider localization.LanguageProvider
}

type service struct {
	stringNormalizer core.StringNormalizer
	featureProvider  core.FeatureProvider
	languageProvider localization.LanguageProvider
	database         global.Database
}

func NewService(database global.Database, opts *ServiceOptions) global.Service {
	if opts == nil {
		opts = &ServiceOptions{}
	}
	opts.EnsureDefaults()

	svc := &service{
		stringNormalizer: opts.StringNormalizer,
		featureProvider:  opts.FeatureProvider,
		languageProvider: opts.LanguageProvider,
		database:         database,
	}

	return svc
}

func (svc *service) StringNormalizer() core.StringNormalizer {
	return svc.stringNormalizer
}

func (svc *service) FeatureProvider() core.FeatureProvider {
	return svc.featureProvider
}

func (svc *service) LanguageProvider() localization.LanguageProvider {
	return svc.languageProvider
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

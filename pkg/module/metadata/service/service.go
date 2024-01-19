package service

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/metadata"
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
	database         metadata.Database
}

func NewService(database metadata.Database, opts *ServiceOptions) metadata.Service {
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

func (opts *ServiceOptions) EnsureDefaults() {
	if opts.StringNormalizer == nil {
		opts.StringNormalizer = core.DefaultStringNormalizer()
	}
	if opts.FeatureProvider == nil {
		opts.FeatureProvider = core.DefaultFeatureProvider()
	}
	if opts.LanguageProvider == nil {
		opts.LanguageProvider = localization.NewDefaultLanguageProvider()
	}
}

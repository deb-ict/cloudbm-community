package service

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
)

type ServiceOptions struct {
	module.ServiceOptions
}

type service struct {
	stringNormalizer core.StringNormalizer
	featureProvider  core.FeatureProvider
	languageProvider localization.LanguageProvider
	database         product.Database
}

func NewService(database product.Database, opts *ServiceOptions) product.Service {
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

func (svc *service) GetStringNormalizer() core.StringNormalizer {
	return svc.stringNormalizer
}

func (svc *service) GetFeatureProvider() core.FeatureProvider {
	return svc.featureProvider
}

func (svc *service) GetLanguageProvider() localization.LanguageProvider {
	return svc.languageProvider
}

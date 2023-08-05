package service

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
)

type ServiceOptions struct {
	FeatureProvider  core.FeatureProvider
	LanguageProvider localization.LanguageProvider
}

type service struct {
	featureProvider  core.FeatureProvider
	languageProvider localization.LanguageProvider
	database         product.Database
}

func NewService(database product.Database, opts *ServiceOptions) product.Service {
	svc := &service{
		database: database,
	}

	if opts != nil {
		svc.featureProvider = opts.FeatureProvider
		svc.languageProvider = opts.LanguageProvider
	}

	svc.ensureDefaults()
	return svc
}

func (svc *service) GetFeatureProvider() core.FeatureProvider {
	return svc.featureProvider
}

func (svc *service) GetLanguageProvider() localization.LanguageProvider {
	return svc.languageProvider
}

func (svc *service) ensureDefaults() {
	if svc.featureProvider == nil {
		svc.featureProvider = core.NewDefaultFeatureProvider()
	}
	if svc.languageProvider == nil {
		svc.languageProvider = localization.NewDefaultLanguageProvider()
	}
}

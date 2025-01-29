package service

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/gallery"
)

type ServiceOptions struct {
	StringNormalizer core.StringNormalizer
	FeatureProvider  core.FeatureProvider
	StorageProvider  core.StorageProvider
	LanguageProvider localization.LanguageProvider
}

type service struct {
	stringNormalizer core.StringNormalizer
	featureProvider  core.FeatureProvider
	storageProvider  core.StorageProvider
	languageProvider localization.LanguageProvider
	database         gallery.Database
}

func NewService(database gallery.Database, opts *ServiceOptions) gallery.Service {
	if opts == nil {
		opts = &ServiceOptions{}
	}
	opts.EnsureDefaults()

	svc := &service{
		stringNormalizer: opts.StringNormalizer,
		featureProvider:  opts.FeatureProvider,
		storageProvider:  opts.StorageProvider,
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

func (svc *service) StorageProvider() core.StorageProvider {
	return svc.storageProvider
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
	if opts.StorageProvider == nil {
		opts.StorageProvider = core.DefaultStorageProvider("/var/lib/cloudbm-community")
	}
	if opts.LanguageProvider == nil {
		opts.LanguageProvider = localization.NewDefaultLanguageProvider()
	}
}

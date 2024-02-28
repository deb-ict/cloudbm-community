package service

import (
	"os/user"
	"path/filepath"
	"strings"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/gallery"
)

type ServiceOptions struct {
	StringNormalizer core.StringNormalizer
	FeatureProvider  core.FeatureProvider
	LanguageProvider localization.LanguageProvider
	StorageFolder    string
}

type service struct {
	stringNormalizer core.StringNormalizer
	featureProvider  core.FeatureProvider
	languageProvider localization.LanguageProvider
	storageFolder    string
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
		languageProvider: opts.LanguageProvider,
		storageFolder:    opts.StorageFolder,
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

func (svc *service) StorageFolder() string {
	return svc.storageFolder
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
	if opts.StorageFolder == "" {
		opts.StorageFolder = "/data/gallery"
	}
	if strings.HasPrefix(opts.StorageFolder, "~") {
		usr, _ := user.Current()
		opts.StorageFolder = filepath.Join(usr.HomeDir, opts.StorageFolder[2:])
	}
}

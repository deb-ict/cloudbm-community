package service

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth"
)

type ServiceOptions struct {
	FeatureProvider core.FeatureProvider
	UserNormalizer  auth.UserNormalizer
	PasswordHasher  auth.PasswordHasher
}

type service struct {
	featureProvider core.FeatureProvider
	userNormalizer  auth.UserNormalizer
	passwordHasher  auth.PasswordHasher
	database        auth.Database
}

func NewService(database auth.Database, opts *ServiceOptions) auth.Service {
	if opts == nil {
		opts = &ServiceOptions{}
	}
	opts.EnsureDefaults()

	svc := &service{
		featureProvider: opts.FeatureProvider,
		userNormalizer:  opts.UserNormalizer,
		passwordHasher:  opts.PasswordHasher,
		database:        database,
	}

	return svc
}

func (svc *service) FeatureProvider() core.FeatureProvider {
	return svc.featureProvider
}

func (opts *ServiceOptions) EnsureDefaults() {
	if opts.FeatureProvider == nil {
		opts.FeatureProvider = core.DefaultFeatureProvider()
	}
	if opts.UserNormalizer == nil {
		opts.UserNormalizer = auth.DefaultUserNormalizer()
	}
	if opts.PasswordHasher == nil {
		opts.PasswordHasher = auth.DefaultPasswordHasher()
	}
}

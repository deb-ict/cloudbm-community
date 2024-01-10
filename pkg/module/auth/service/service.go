package service

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/security"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/util"
)

type ServiceOptions struct {
	FeatureProvider core.FeatureProvider
	UserNormalizer  util.UserNormalizer
	PasswordHasher  security.PasswordHasher
}

type service struct {
	featureProvider core.FeatureProvider
	userNormalizer  util.UserNormalizer
	passwordHasher  security.PasswordHasher
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

func (svc *service) UserNormalizer() util.UserNormalizer {
	return svc.userNormalizer
}

func (svc *service) PasswordHasher() security.PasswordHasher {
	return svc.passwordHasher
}

func (svc *service) FeatureProvider() core.FeatureProvider {
	return svc.featureProvider
}

func (opts *ServiceOptions) EnsureDefaults() {
	if opts.FeatureProvider == nil {
		opts.FeatureProvider = core.DefaultFeatureProvider()
	}
	if opts.UserNormalizer == nil {
		opts.UserNormalizer = util.DefaultUserNormalizer()
	}
	if opts.PasswordHasher == nil {
		opts.PasswordHasher = security.DefaultPasswordHasher()
	}
}

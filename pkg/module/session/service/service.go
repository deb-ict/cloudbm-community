package service

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/session"
)

type ServiceOptions struct {
	FeatureProvider core.FeatureProvider
}

type service struct {
	featureProvider core.FeatureProvider
	database        session.Database
}

func NewService(database session.Database, opts *ServiceOptions) session.Service {
	if opts == nil {
		opts = &ServiceOptions{}
	}
	opts.EnsureDefaults()

	svc := &service{
		featureProvider: opts.FeatureProvider,
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
}

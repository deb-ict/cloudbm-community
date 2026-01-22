package service

import (
	"time"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/session"
)

type ServiceOptions struct {
	FeatureProvider       core.FeatureProvider
	SessionTimeoutMinutes int64
}

type service struct {
	featureProvider core.FeatureProvider
	sessionTimeout  time.Duration
	database        session.Database
}

func NewService(database session.Database, opts *ServiceOptions) session.Service {
	if opts == nil {
		opts = &ServiceOptions{}
	}
	opts.EnsureDefaults()

	svc := &service{
		featureProvider: opts.FeatureProvider,
		sessionTimeout:  time.Duration(opts.SessionTimeoutMinutes) * time.Minute,
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
	if opts.SessionTimeoutMinutes == 0 {
		opts.SessionTimeoutMinutes = 30
	}
}

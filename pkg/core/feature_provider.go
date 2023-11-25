package core

import (
	"context"
)

type FeatureProvider interface {
	FeatureEnabled(ctx context.Context, name string) bool
}

type defaultFeatureProvider struct {
}

func DefaultFeatureProvider() FeatureProvider {
	return &defaultFeatureProvider{}
}

func (p *defaultFeatureProvider) FeatureEnabled(ctx context.Context, name string) bool {
	return false
}

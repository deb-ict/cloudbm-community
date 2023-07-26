package core

import (
	"context"
)

type FeatureProvider interface {
	FeatureEnabled(ctx context.Context, name string) bool
}

type defaultFeatureProvider struct {
}

func NewDefaultFeatureProvider() FeatureProvider {
	return &defaultFeatureProvider{}
}

func (p *defaultFeatureProvider) FeatureEnabled(ctx context.Context, name string) bool {
	return false
}

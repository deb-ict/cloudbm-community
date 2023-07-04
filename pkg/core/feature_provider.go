package core

import (
	"context"
)

type FeatureProvider interface {
	FeatureEnabled(ctx context.Context, name string) bool
}

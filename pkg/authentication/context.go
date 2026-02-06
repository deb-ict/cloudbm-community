package authentication

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/logging"
)

type ContextKey string

const (
	IdentityContextKey ContextKey = "cbm.identity"
)

func WithIdentityInContext(ctx context.Context, identity Identity) context.Context {
	return context.WithValue(ctx, IdentityContextKey, identity)
}

func GetIdentityFromContext(ctx context.Context) Identity {
	value := ctx.Value(IdentityContextKey)
	if value == nil {
		logging.GetLoggerFromContext(ctx).WarnContext(ctx, "Failed to get identity from context: identity is missing")
		return nil
	}
	identity, ok := value.(Identity)
	if !ok {
		logging.GetLoggerFromContext(ctx).WarnContext(ctx, "Failed to get identity from context: invalid type")
		return nil
	}
	return identity
}

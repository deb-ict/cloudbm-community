package authorization

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/authentication"
	"github.com/deb-ict/cloudbm-community/pkg/logging"
	"github.com/gorilla/mux"
)

type AuthorizationMiddleware struct {
	requiredPolicies map[string]string
}

func NewMiddleware() *AuthorizationMiddleware {
	return &AuthorizationMiddleware{}
}

func (m *AuthorizationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logging.GetLoggerFromContext(ctx)
		identity := authentication.GetIdentityFromContext(ctx)

		if identity == nil {
			logger.WarnContext(ctx, "Failed to authorize: Identity is missing in context")
			ForbiddenResponse(w)
			return
		}

		// TODO: We must setup authorization policies here
		// Create a map with the api name and required policy
		// In case the route has no name, fail
		routeName := mux.CurrentRoute(r).GetName()
		if routeName == "" {
			logger.WarnContext(ctx, "Failed to authorize: Route has no name")
			ForbiddenResponse(w)
			return
		}

		logger.InfoContext(ctx, "User is authorized to access the route")

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func ForbiddenResponse(w http.ResponseWriter) {
	http.Error(w, "Access denied", http.StatusForbidden)
}

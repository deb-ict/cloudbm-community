package authorization

import (
	"net/http"
)

type AuthorizationMiddleware struct {
}

func NewMiddleware() *AuthorizationMiddleware {
	return &AuthorizationMiddleware{}
}

func (m *AuthorizationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement authentication logic here

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

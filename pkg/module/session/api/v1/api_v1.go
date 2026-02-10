package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/session"
	"github.com/deb-ict/go-router"
	"github.com/deb-ict/go-router/authorization"
)

const (
	PolicyReadSessionsV1    = "session_api:ReadSessions:v1"
	PolicyCreateSessionsV1  = "session_api:CreateSessions:v1"
	PolicyUpdateSessionsV1  = "session_api:UpdateSessions:v1"
	PolicyDeleteSessionsV1  = "session_api:DeleteSessions:v1"
	PolicyCleanupSessionsV1 = "session_api:CleanupSessions:v1"
)

type ApiV1 interface {
	RegisterAuthorizationPolicies(middleware *authorization.Middleware)
	RegisterRoutes(r *router.Router)
}

type apiV1 struct {
	service session.Service
}

func NewApiV1(service session.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterAuthorizationPolicies(middleware *authorization.Middleware) {
	middleware.SetPolicy(authorization.NewPolicy(PolicyReadSessionsV1,
		authorization.NewScopeRequirement("session.read"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyCreateSessionsV1,
		authorization.NewScopeRequirement("session.create"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyUpdateSessionsV1,
		authorization.NewScopeRequirement("session.update"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyDeleteSessionsV1,
		authorization.NewScopeRequirement("session.delete"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyCleanupSessionsV1,
		authorization.NewScopeRequirement("session.cleanup"),
	))
}

func (api *apiV1) RegisterRoutes(r *router.Router) {
	r.HandleFunc("/v1/session", api.GetSessionsHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadSessionsV1),
	)
	r.HandleFunc("/v1/session/{id}", api.GetSessionByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadSessionsV1),
	)
	r.HandleFunc("/v1/session", api.CreateSessionHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyCreateSessionsV1),
	)
	r.HandleFunc("/v1/session/{id}", api.UpdateSessionHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyUpdateSessionsV1),
	)
	r.HandleFunc("/v1/session/{id}", api.DeleteSessionHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyDeleteSessionsV1),
	)
	r.HandleFunc("/v1/session/cleanup", api.CleanupExpiredSessionsHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyCleanupSessionsV1),
	)
}

func (api *apiV1) handleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	switch err {
	case session.ErrSessionNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case session.ErrSessionExpired:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case core.ErrInvalidId:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	default:
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
	}
	return true
}

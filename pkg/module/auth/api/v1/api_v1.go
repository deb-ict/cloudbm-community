package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth"
	"github.com/deb-ict/go-router"
	"github.com/deb-ict/go-router/authorization"
)

const (
	PolicyReadUsersV1   = "auth_api:ReadUsers:v1"
	PolicyCreateUsersV1 = "auth_api:CreateUsers:v1"
	PolicyUpdateUsersV1 = "auth_api:UpdateUsers:v1"
	PolicyDeleteUsersV1 = "auth_api:DeleteUsers:v1"
)

type ApiV1 interface {
	RegisterAuthorizationPolicies(middleware *authorization.Middleware)
	RegisterRoutes(r *router.Router)
}

type apiV1 struct {
	service auth.Service
}

func NewApiV1(service auth.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterAuthorizationPolicies(middleware *authorization.Middleware) {
	middleware.SetPolicy(authorization.NewPolicy(PolicyReadUsersV1,
		authorization.NewScopeRequirement("user.read"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyCreateUsersV1,
		authorization.NewScopeRequirement("user.create"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyUpdateUsersV1,
		authorization.NewScopeRequirement("user.update"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyDeleteUsersV1,
		authorization.NewScopeRequirement("user.delete"),
	))
}

func (api *apiV1) RegisterRoutes(r *router.Router) {
	// Users
	r.HandleFunc("/v1/user", api.GetUsersHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadUsersV1),
	)
	r.HandleFunc("/v1/user/{id}", api.GetUserByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadUsersV1),
	)
	r.HandleFunc("/v1/user", api.CreateUserHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyCreateUsersV1),
	)
	r.HandleFunc("/v1/user/{id}", api.UpdateUserHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyUpdateUsersV1),
	)
	r.HandleFunc("/v1/user/{id}", api.DeleteUserHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyDeleteUsersV1),
	)
}

func (api *apiV1) handleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	switch err {
	case auth.ErrUserNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case auth.ErrDuplicateUsername:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case auth.ErrDuplicateEmail:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case core.ErrInvalidId:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	default:
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
	}
	return true
}

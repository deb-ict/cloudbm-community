package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth"
	"github.com/gorilla/mux"
)

type ApiV1 interface {
	RegisterRoutes(r *mux.Router)
}

type apiV1 struct {
	service auth.Service
}

func NewApiV1(service auth.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterRoutes(r *mux.Router) {
	// Users
	r.HandleFunc("/v1/user", api.GetUsersHandlerV1).Methods(http.MethodGet).Name("auth_api:GetUsersHandlerV1")
	r.HandleFunc("/v1/user/{id}", api.GetUserByIdHandlerV1).Methods(http.MethodGet).Name("auth_api:GetUserByIdHandlerV1")
	r.HandleFunc("/v1/user", api.CreateUserHandlerV1).Methods(http.MethodPost).Name("auth_api:CreateUserHandlerV1")
	r.HandleFunc("/v1/user/{id}", api.UpdateUserHandlerV1).Methods(http.MethodPut).Name("auth_api:UpdateUserHandlerV1")
	r.HandleFunc("/v1/user/{id}", api.DeleteUserHandlerV1).Methods(http.MethodDelete).Name("auth_api:DeleteUserHandlerV1")
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

package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/session"
	"github.com/gorilla/mux"
)

type ApiV1 interface {
	RegisterRoutes(r *mux.Router)
}

type apiV1 struct {
	service session.Service
}

func NewApiV1(service session.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/v1/session", api.GetSessionsHandlerV1).Methods(http.MethodGet).Name("session_api:GetSessionsHandlerV1")
	r.HandleFunc("/v1/session/{id}", api.GetSessionByIdHandlerV1).Methods(http.MethodGet).Name("session_api:GetSessionByIdHandlerV1")
	r.HandleFunc("/v1/session", api.CreateSessionHandlerV1).Methods(http.MethodPost).Name("session_api:CreateSessionHandlerV1")
	r.HandleFunc("/v1/session/{id}", api.UpdateSessionHandlerV1).Methods(http.MethodPut).Name("session_api:UpdateSessionHandlerV1")
	r.HandleFunc("/v1/session/{id}", api.DeleteSessionHandlerV1).Methods(http.MethodDelete).Name("session_api:DeleteSessionHandlerV1")
	r.HandleFunc("/v1/session/cleanup", api.CleanupExpiredSessionsHandlerV1).Methods(http.MethodPost).Name("session_api:CleanupExpiredSessionsHandlerV1")
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

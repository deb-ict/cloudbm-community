package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/session"
	"github.com/gorilla/mux"
)

const (
	RouteGetSessionsV1            = "session_api:GetSessions:v1"
	RouteGetSessionByIdV1         = "session_api:GetSessionById:v1"
	RouteCreateSessionV1          = "session_api:CreateSession:v1"
	RouteUpdateSessionV1          = "session_api:UpdateSession:v1"
	RouteDeleteSessionV1          = "session_api:DeleteSession:v1"
	RouteCleanupExpiredSessionsV1 = "session_api:CleanupExpiredSessions:v1"
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
	r.HandleFunc("/v1/session", api.GetSessionsHandlerV1).Methods(http.MethodGet).Name(RouteGetSessionsV1)
	r.HandleFunc("/v1/session/{id}", api.GetSessionByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetSessionByIdV1)
	r.HandleFunc("/v1/session", api.CreateSessionHandlerV1).Methods(http.MethodPost).Name(RouteCreateSessionV1)
	r.HandleFunc("/v1/session/{id}", api.UpdateSessionHandlerV1).Methods(http.MethodPut).Name(RouteUpdateSessionV1)
	r.HandleFunc("/v1/session/{id}", api.DeleteSessionHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteSessionV1)
	r.HandleFunc("/v1/session/cleanup", api.CleanupExpiredSessionsHandlerV1).Methods(http.MethodPost).Name(RouteCleanupExpiredSessionsV1)
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

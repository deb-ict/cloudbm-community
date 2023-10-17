package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/metadata"
	"github.com/gorilla/mux"
)

type ApiV1 interface {
	RegisterRoutes(r *mux.Router)
}

type apiV1 struct {
	service metadata.Service
}

func NewApi(service metadata.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterRoutes(r *mux.Router) {
	// Tax profiles
	r.HandleFunc("/v1/taxProfile", api.GetTaxProfilesHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/taxProfile/{id}", api.GetTaxProfileByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/taxProfile", api.CreateTaxProfileHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/taxProfile/{id}", api.UpdateTaxProfileHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/taxProfile/{id}", api.DeleteTaxProfileHandlerV1).Methods(http.MethodDelete)
}

func (api *apiV1) handleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	switch err {
	case metadata.ErrTaxProfileNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case metadata.ErrTaxProfileDuplicateKey:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case metadata.ErrTaxProfileDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	default:
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
	}
	return true
}

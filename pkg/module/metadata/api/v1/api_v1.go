package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/metadata"
	"github.com/deb-ict/go-router"
)

type ApiV1 interface {
	RegisterRoutes(r *router.Router)
}

type apiV1 struct {
	service metadata.Service
}

func NewApi(service metadata.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterRoutes(r *router.Router) {
	// Tax profiles
	r.HandleFunc(
		"/v1/taxProfile",
		api.GetTaxProfilesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized("taxProfile.read"),
	)
	r.HandleFunc(
		"/v1/taxProfile/{id}",
		api.GetTaxProfileByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized("taxProfile.read"),
	)
	r.HandleFunc(
		"/v1/taxProfile",
		api.CreateTaxProfileHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized("taxProfile.create"),
	)
	r.HandleFunc(
		"/v1/taxProfile/{id}",
		api.UpdateTaxProfileHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized("taxProfile.update"),
	)
	r.HandleFunc(
		"/v1/taxProfile/{id}",
		api.DeleteTaxProfileHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized("taxProfile.delete"),
	)
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

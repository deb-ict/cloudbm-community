package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/metadata"
	"github.com/gorilla/mux"
)

const (
	RouteGetTaxRatesV1    = "metadata_api:GetTaxRates:v1"
	RouteGetTaxRateByIdV1 = "metadata_api:GetTaxRateById:v1"
	RouteCreateTaxRateV1  = "metadata_api:CreateTaxRate:v1"
	RouteUpdateTaxRateV1  = "metadata_api:UpdateTaxRate:v1"
	RouteDeleteTaxRateV1  = "metadata_api:DeleteTaxRate:v1"
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
	r.HandleFunc("/v1/taxRate", api.GetTaxRatesHandlerV1).Methods(http.MethodGet).Name(RouteGetTaxRatesV1)
	r.HandleFunc("/v1/taxRate/{id}", api.GetTaxRateByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetTaxRateByIdV1)
	r.HandleFunc("/v1/taxRate", api.CreateTaxRateHandlerV1).Methods(http.MethodPost).Name(RouteCreateTaxRateV1)
	r.HandleFunc("/v1/taxRate/{id}", api.UpdateTaxRateHandlerV1).Methods(http.MethodPut).Name(RouteUpdateTaxRateV1)
	r.HandleFunc("/v1/taxRate/{id}", api.DeleteTaxRateHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteTaxRateV1)
}

func (api *apiV1) handleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	switch err {
	case metadata.ErrTaxRateNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case metadata.ErrTaxRateDuplicateKey:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case metadata.ErrTaxRateDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	default:
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
	}
	return true
}

package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/metadata"
	"github.com/deb-ict/go-router"
	"github.com/deb-ict/go-router/authorization"
)

const (
	PolicyReadMetadataV1   = "metadata_api:ReadMetadata:v1"
	PolicyCreateMetadataV1 = "metadata_api:CreateMetadata:v1"
	PolicyUpdateMetadataV1 = "metadata_api:UpdateMetadata:v1"
	PolicyDeleteMetadataV1 = "metadata_api:DeleteMetadata:v1"
)

type ApiV1 interface {
	RegisterAuthorizationPolicies(middleware *authorization.Middleware)
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

func (api *apiV1) RegisterAuthorizationPolicies(middleware *authorization.Middleware) {
	middleware.SetPolicy(authorization.NewPolicy(PolicyReadMetadataV1,
		authorization.NewScopeRequirement("metadata.read"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyCreateMetadataV1,
		authorization.NewScopeRequirement("metadata.create"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyUpdateMetadataV1,
		authorization.NewScopeRequirement("metadata.update"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyDeleteMetadataV1,
		authorization.NewScopeRequirement("metadata.delete"),
	))
}

func (api *apiV1) RegisterRoutes(r *router.Router) {
	// Tax profiles
	r.HandleFunc("/v1/taxRate", api.GetTaxRatesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadMetadataV1),
	)
	r.HandleFunc("/v1/taxRate/{id}", api.GetTaxRateByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadMetadataV1),
	)
	r.HandleFunc("/v1/taxRate", api.CreateTaxRateHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyCreateMetadataV1),
	)
	r.HandleFunc("/v1/taxRate/{id}", api.UpdateTaxRateHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyUpdateMetadataV1),
	)
	r.HandleFunc("/v1/taxRate/{id}", api.DeleteTaxRateHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyDeleteMetadataV1),
	)
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

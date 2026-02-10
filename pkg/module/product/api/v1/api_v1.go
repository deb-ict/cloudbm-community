package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
	"github.com/deb-ict/go-router"
	"github.com/deb-ict/go-router/authorization"
)

const (
	PolicyReadProductsV1   = "product_api:ReadProducts:v1"
	PolicyCreateProductsV1 = "product_api:CreateProducts:v1"
	PolicyUpdateProductsV1 = "product_api:UpdateProducts:v1"
	PolicyDeleteProductsV1 = "product_api:DeleteProducts:v1"
)

type ApiV1 interface {
	RegisterAuthorizationPolicies(middleware *authorization.Middleware)
	RegisterRoutes(r *router.Router)
}

type apiV1 struct {
	service product.Service
}

func NewApiV1(service product.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterAuthorizationPolicies(middleware *authorization.Middleware) {
	middleware.SetPolicy(authorization.NewPolicy(PolicyReadProductsV1,
		authorization.NewScopeRequirement("product.read"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyCreateProductsV1,
		authorization.NewScopeRequirement("product.create"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyUpdateProductsV1,
		authorization.NewScopeRequirement("product.update"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyDeleteProductsV1,
		authorization.NewScopeRequirement("product.delete"),
	))
}

func (api *apiV1) RegisterRoutes(r *router.Router) {
	// Attributes
	r.HandleFunc("/v1/attribute", api.GetAttributesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadProductsV1),
	)
	r.HandleFunc("/v1/attribute/{id}", api.GetAttributeByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadProductsV1),
	)
	r.HandleFunc("/v1/attribute", api.CreateAttributeHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyCreateProductsV1),
	)
	r.HandleFunc("/v1/attribute/{id}", api.UpdateAttributeHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyUpdateProductsV1),
	)
	r.HandleFunc("/v1/attribute/{id}", api.DeleteAttributeHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyDeleteProductsV1),
	)

	// Attribute Values
	r.HandleFunc("/v1/attribute/{attributeId}/value", api.GetAttributeValuesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadProductsV1),
	)
	r.HandleFunc("/v1/attribute/{attributeId}/value/{id}", api.GetAttributeValueByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadProductsV1),
	)
	r.HandleFunc("/v1/attribute/{attributeId}/value", api.CreateAttributeValueHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyCreateProductsV1),
	)
	r.HandleFunc("/v1/attribute/{attributeId}/value/{id}", api.UpdateAttributeValueHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyUpdateProductsV1),
	)
	r.HandleFunc("/v1/attribute/{attributeId}/value/{id}", api.DeleteAttributeValueHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyDeleteProductsV1),
	)

	// Categories
	r.HandleFunc("/v1/category", api.GetCategoriesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadProductsV1),
	)
	r.HandleFunc("/v1/category/{id}", api.GetCategoryByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadProductsV1),
	)
	r.HandleFunc("/v1/category", api.CreateCategoryHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyCreateProductsV1),
	)
	r.HandleFunc("/v1/category/{id}", api.UpdateCategoryHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyUpdateProductsV1),
	)
	r.HandleFunc("/v1/category/{id}", api.DeleteCategoryHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyDeleteProductsV1),
	)

	// Products
	r.HandleFunc("/v1/product", api.GetProductsHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadProductsV1),
	)
	r.HandleFunc("/v1/product/{id}", api.GetProductByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadProductsV1),
	)
	r.HandleFunc("/v1/product", api.CreateProductHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyCreateProductsV1),
	)
	r.HandleFunc("/v1/product/{id}", api.UpdateProductHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyUpdateProductsV1),
	)
	r.HandleFunc("/v1/product/{id}", api.DeleteProductHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyDeleteProductsV1),
	)
}

func (api *apiV1) handleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	switch err {
	case product.ErrAttributeNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case product.ErrAttributeDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case product.ErrAttributeDuplicateSlug:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case product.ErrAttributeValueNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case product.ErrAttributeValueDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case product.ErrAttributeValueDuplicateSlug:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case product.ErrCategoryNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case product.ErrCategoryDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case product.ErrCategoryDuplicateSlug:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case product.ErrProductNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case product.ErrProductDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case product.ErrProductDuplicateSlug:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case product.ErrProductVariantNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case product.ErrProductVariantDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case product.ErrProductVariantDuplicateSlug:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case core.ErrTranslationNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case core.ErrInvalidId:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	default:
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
	}
	return true
}

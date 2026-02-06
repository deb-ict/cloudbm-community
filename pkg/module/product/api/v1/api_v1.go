package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
	"github.com/gorilla/mux"
)

const (
	RouteGetAttributesV1         = "product_api:GetAttributes:v1"
	RouteGetAttributeByIdV1      = "product_api:GetAttributeById:v1"
	RouteCreateAttributeV1       = "product_api:CreateAttribute:v1"
	RouteUpdateAttributeV1       = "product_api:UpdateAttribute:v1"
	RouteDeleteAttributeV1       = "product_api:DeleteAttribute:v1"
	RouteGetAttributeValuesV1    = "product_api:GetAttributeValues:v1"
	RouteGetAttributeValueByIdV1 = "product_api:GetAttributeValueById:v1"
	RouteCreateAttributeValueV1  = "product_api:CreateAttributeValue:v1"
	RouteUpdateAttributeValueV1  = "product_api:UpdateAttributeValue:v1"
	RouteDeleteAttributeValueV1  = "product_api:DeleteAttributeValue:v1"
	RouteGetCategoriesV1         = "product_api:GetCategories:v1"
	RouteGetCategoryByIdV1       = "product_api:GetCategoryById:v1"
	RouteCreateCategoryV1        = "product_api:CreateCategory:v1"
	RouteUpdateCategoryV1        = "product_api:UpdateCategory:v1"
	RouteDeleteCategoryV1        = "product_api:DeleteCategory:v1"
	RouteGetProductsV1           = "product_api:GetProducts:v1"
	RouteGetProductByIdV1        = "product_api:GetProductById:v1"
	RouteCreateProductV1         = "product_api:CreateProduct:v1"
	RouteUpdateProductV1         = "product_api:UpdateProduct:v1"
	RouteDeleteProductV1         = "product_api:DeleteProduct:v1"
)

type ApiV1 interface {
	RegisterRoutes(r *mux.Router)
}

type apiV1 struct {
	service product.Service
}

func NewApiV1(service product.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterRoutes(r *mux.Router) {
	// Attributes
	r.HandleFunc("/v1/attribute", api.GetAttributesHandlerV1).Methods(http.MethodGet).Name(RouteGetAttributesV1)
	r.HandleFunc("/v1/attribute/{id}", api.GetAttributeByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetAttributeByIdV1)
	r.HandleFunc("/v1/attribute", api.CreateAttributeHandlerV1).Methods(http.MethodPost).Name(RouteCreateAttributeV1)
	r.HandleFunc("/v1/attribute/{id}", api.UpdateAttributeHandlerV1).Methods(http.MethodPut).Name(RouteUpdateAttributeV1)
	r.HandleFunc("/v1/attribute/{id}", api.DeleteAttributeHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteAttributeV1)

	// Attribute Values
	r.HandleFunc("/v1/attribute/{attributeId}/value", api.GetAttributeValuesHandlerV1).Methods(http.MethodGet).Name(RouteGetAttributeValuesV1)
	r.HandleFunc("/v1/attribute/{attributeId}/value/{id}", api.GetAttributeValueByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetAttributeValueByIdV1)
	r.HandleFunc("/v1/attribute/{attributeId}/value", api.CreateAttributeValueHandlerV1).Methods(http.MethodPost).Name(RouteCreateAttributeValueV1)
	r.HandleFunc("/v1/attribute/{attributeId}/value/{id}", api.UpdateAttributeValueHandlerV1).Methods(http.MethodPut).Name(RouteUpdateAttributeValueV1)
	r.HandleFunc("/v1/attribute/{attributeId}/value/{id}", api.DeleteAttributeValueHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteAttributeValueV1)
	// Categories
	r.HandleFunc("/v1/category", api.GetCategoriesHandlerV1).Methods(http.MethodGet).Name(RouteGetCategoriesV1)
	r.HandleFunc("/v1/category/{id}", api.GetCategoryByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetCategoryByIdV1)
	r.HandleFunc("/v1/category", api.CreateCategoryHandlerV1).Methods(http.MethodPost).Name(RouteCreateCategoryV1)
	r.HandleFunc("/v1/category/{id}", api.UpdateCategoryHandlerV1).Methods(http.MethodPut).Name(RouteUpdateCategoryV1)
	r.HandleFunc("/v1/category/{id}", api.DeleteCategoryHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteCategoryV1)

	// Products
	r.HandleFunc("/v1/product", api.GetProductsHandlerV1).Methods(http.MethodGet).Name(RouteGetProductsV1)
	r.HandleFunc("/v1/product/{id}", api.GetProductByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetProductByIdV1)
	r.HandleFunc("/v1/product", api.CreateProductHandlerV1).Methods(http.MethodPost).Name(RouteCreateProductV1)
	r.HandleFunc("/v1/product/{id}", api.UpdateProductHandlerV1).Methods(http.MethodPut).Name(RouteUpdateProductV1)
	r.HandleFunc("/v1/product/{id}", api.DeleteProductHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteProductV1)
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

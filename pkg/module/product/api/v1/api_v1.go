package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
	"github.com/gorilla/mux"
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
	r.HandleFunc("/v1/attribute", api.GetAttributesHandlerV1).Methods(http.MethodGet).Name("product_api:GetAttributesHandlerV1")
	r.HandleFunc("/v1/attribute/{id}", api.GetAttributeByIdHandlerV1).Methods(http.MethodGet).Name("product_api:GetAttributeByIdHandlerV1")
	r.HandleFunc("/v1/attribute", api.CreateAttributeHandlerV1).Methods(http.MethodPost).Name("product_api:CreateAttributeHandlerV1")
	r.HandleFunc("/v1/attribute/{id}", api.UpdateAttributeHandlerV1).Methods(http.MethodPut).Name("product_api:UpdateAttributeHandlerV1")
	r.HandleFunc("/v1/attribute/{id}", api.DeleteAttributeHandlerV1).Methods(http.MethodDelete).Name("product_api:DeleteAttributeHandlerV1")

	// Attribute Values
	r.HandleFunc("/v1/attribute/{attributeId}/value", api.GetAttributeValuesHandlerV1).Methods(http.MethodGet).Name("product_api:GetAttributeValuesHandlerV1")
	r.HandleFunc("/v1/attribute/{attributeId}/value/{id}", api.GetAttributeValueByIdHandlerV1).Methods(http.MethodGet).Name("product_api:GetAttributeValueByIdHandlerV1")
	r.HandleFunc("/v1/attribute/{attributeId}/value", api.CreateAttributeValueHandlerV1).Methods(http.MethodPost).Name("product_api:CreateAttributeValueHandlerV1")
	r.HandleFunc("/v1/attribute/{attributeId}/value/{id}", api.UpdateAttributeValueHandlerV1).Methods(http.MethodPut).Name("product_api:UpdateAttributeValueHandlerV1")
	r.HandleFunc("/v1/attribute/{attributeId}/value/{id}", api.DeleteAttributeValueHandlerV1).Methods(http.MethodDelete).Name("product_api:DeleteAttributeValueHandlerV1")

	// Categories
	r.HandleFunc("/v1/category", api.GetCategoriesHandlerV1).Methods(http.MethodGet).Name("product_api:GetCategoriesHandlerV1")
	r.HandleFunc("/v1/category/{id}", api.GetCategoryByIdHandlerV1).Methods(http.MethodGet).Name("product_api:GetCategoryByIdHandlerV1")
	r.HandleFunc("/v1/category", api.CreateCategoryHandlerV1).Methods(http.MethodPost).Name("product_api:CreateCategoryHandlerV1")
	r.HandleFunc("/v1/category/{id}", api.UpdateCategoryHandlerV1).Methods(http.MethodPut).Name("product_api:UpdateCategoryHandlerV1")
	r.HandleFunc("/v1/category/{id}", api.DeleteCategoryHandlerV1).Methods(http.MethodDelete).Name("product_api:DeleteCategoryHandlerV1")

	// Products
	r.HandleFunc("/v1/product", api.GetProductsHandlerV1).Methods(http.MethodGet).Name("product_api:GetProductsHandlerV1")
	r.HandleFunc("/v1/product/{id}", api.GetProductByIdHandlerV1).Methods(http.MethodGet).Name("product_api:GetProductByIdHandlerV1")
	r.HandleFunc("/v1/product", api.CreateProductHandlerV1).Methods(http.MethodPost).Name("product_api:CreateProductHandlerV1")
	r.HandleFunc("/v1/product/{id}", api.UpdateProductHandlerV1).Methods(http.MethodPut).Name("product_api:UpdateProductHandlerV1")
	r.HandleFunc("/v1/product/{id}", api.DeleteProductHandlerV1).Methods(http.MethodDelete).Name("product_api:DeleteProductHandlerV1")

	// Product variants
	r.HandleFunc("/v1/product/{productId}/variant", api.GetProductVariantsHandlerV1).Methods(http.MethodGet).Name("product_api:GetProductVariantsHandlerV1")
	r.HandleFunc("/v1/product/{productId}/variant/{id}", api.GetProductVariantByIdHandlerV1).Methods(http.MethodGet).Name("product_api:GetProductVariantByIdHandlerV1")
	r.HandleFunc("/v1/product/{productId}/variant", api.CreateProductVariantHandlerV1).Methods(http.MethodPost).Name("product_api:CreateProductVariantHandlerV1")
	r.HandleFunc("/v1/product/{productId}/variant/{id}", api.UpdateProductVariantHandlerV1).Methods(http.MethodPut).Name("product_api:UpdateProductVariantHandlerV1")
	r.HandleFunc("/v1/product/{productId}/variant/{id}", api.DeleteProductVariantHandlerV1).Methods(http.MethodDelete).Name("product_api:DeleteProductVariantHandlerV1")
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

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
	// Products
	r.HandleFunc("/v1/product", api.GetProductsHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/product/{id}", api.GetProductByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/product", api.CreateProductHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/product/{id}", api.UpdateProductHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/product/{id}", api.DeleteProductHandlerV1).Methods(http.MethodDelete)

	// Categories
	r.HandleFunc("/v1/category", api.GetCateogiesHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/category/{id}", api.GetCategoryByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/category", api.CreateCategoryHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/category/{id}", api.UpdateCategoryHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/category/{id}", api.DeleteCategoryHandlerV1).Methods(http.MethodDelete)
}

func (api *apiV1) handleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	switch err {
	case product.ErrProductNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case product.ErrCategoryNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case core.ErrTranslationNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case core.ErrInvalidId:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	default:
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
	}
	return true
}

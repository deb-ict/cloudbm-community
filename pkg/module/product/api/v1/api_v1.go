package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
	"github.com/deb-ict/go-router"
)

type ApiV1 interface {
	RegisterRoutes(r *router.Router)
}

type apiV1 struct {
	service product.Service
}

func NewApi(service product.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterRoutes(r *router.Router) {
	// Products
	r.HandleFunc(
		"/v1/product",
		api.GetProductsHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized("product.read"),
	)
	r.HandleFunc(
		"/v1/product/{id}",
		api.GetProductByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized("product.read"),
	)
	r.HandleFunc(
		"/v1/product",
		api.CreateProductHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized("product.create"),
	)
	r.HandleFunc(
		"/v1/product/{id}",
		api.UpdateProductHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized("product.update"),
	)
	r.HandleFunc(
		"/v1/product/{id}",
		api.DeleteProductHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized("product.delete"),
	)

	// Categories
	r.HandleFunc(
		"/v1/category",
		api.GetCateogiesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized("product.read"),
	)
	r.HandleFunc(
		"/v1/category/{id}",
		api.GetCategoryByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized("product.read"),
	)
	r.HandleFunc(
		"/v1/category",
		api.CreateCategoryHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized("product.create"),
	)
	r.HandleFunc(
		"/v1/category/{id}",
		api.UpdateCategoryHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized("product.update"),
	)
	r.HandleFunc(
		"/v1/category/{id}",
		api.DeleteCategoryHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized("product.delete"),
	)
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

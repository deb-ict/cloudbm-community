package v1

import (
	"net/http"

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
	api.registerProductRoutes(r)
	api.registerCategoryRoutes(r)
}

func (api *apiV1) handleError(w http.ResponseWriter, err error) bool {
	if err != nil {
		switch err {
		case product.ErrProductNotFound:
			rest.WriteError(w, http.StatusNotFound, err.Error())
		case product.ErrCategoryNotFound:
			rest.WriteError(w, http.StatusNotFound, err.Error())
		default:
			rest.WriteError(w, http.StatusInternalServerError, err.Error())
		}
		return true
	}
	return false
}

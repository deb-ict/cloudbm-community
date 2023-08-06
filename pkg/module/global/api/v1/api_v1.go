package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/global"
	"github.com/deb-ict/go-router"
)

type ApiV1 interface {
	RegisterRoutes(r *router.Router)
	RegisterTaxProfileRoutes(r *router.Router)
}

type apiV1 struct {
	service global.Service
}

func NewApi(service global.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterRoutes(r *router.Router) {
	api.RegisterTaxProfileRoutes(r)
}

func (api *apiV1) handleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	switch err {
	case global.ErrTaxProfileNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case global.ErrTaxProfileDuplicateKey:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case global.ErrTaxProfileDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	default:
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
	}
	return true
}

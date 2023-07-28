package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/go-router"
)

type ApiV1 interface {
	RegisterRoutes(r *router.Router)
}

type apiV1 struct {
	service contact.Service
}

func NewApi(service contact.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterRoutes(r *router.Router) {

}

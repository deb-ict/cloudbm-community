package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/module/session"
	"github.com/gorilla/mux"
)

type ApiV1 interface {
	RegisterRoutes(r *mux.Router)
}

type apiV1 struct {
	service session.Service
}

func NewApiV1(service session.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterRoutes(r *mux.Router) {
}

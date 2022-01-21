package customer

import (
	"github.com/gorilla/mux"
)

type Api interface {
	RegisterRoutes(r *mux.Router)
}

type api struct {
	svc Service
}

func NewApi(s Service) Api {
	return &api{
		svc: s,
	}
}

func (api *api) RegisterRoutes(r *mux.Router) {

}

func (api *api) GetService() Service {
	return api.svc
}

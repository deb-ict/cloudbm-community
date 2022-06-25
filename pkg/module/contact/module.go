package contact

import (
	"github.com/deb-ict/cloudbm-community/pkg/hosting"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func NewModule(db Database) hosting.Module {
	svc := newService(db)
	api := newApiHandler(svc)
	return &module{
		svc: svc,
		api: api,
	}
}

type module struct {
	svc Service
	api ApiHandler
}

func (mod *module) LoadConfig(configPath string) error {
	err := mod.svc.GetDatabase().LoadConfig(configPath)
	if err != nil {
		return err
	}
	return nil
}

func (mod *module) RegisterApiRoutes(router *mux.Router) error {
	mod.api.RegisterRoutes(router)
	return nil
}

func (mod *module) RegisterGrpcServices(server *grpc.Server) error {
	return nil
}

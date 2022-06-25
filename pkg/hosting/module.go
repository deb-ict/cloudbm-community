package hosting

import (
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type Module interface {
	LoadConfig(configPath string) error
	RegisterApiRoutes(router *mux.Router) error
	RegisterGrpcServices(server *grpc.Server) error
}

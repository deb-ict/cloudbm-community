package hosting

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	ErrAlreadyRunning  error = errors.New("host already running")
	ErrModuleDuplicate error = errors.New("module already exists")
	ErrModuleNotFound  error = errors.New("module not found")
)

type MiddlewareFunc func(http.Handler) http.Handler

type Host interface {
	Run() error
	GetConfig() HostConfig
	GetConfigPath() string
	SetConfigPath(configPath string)
	GetHttpRouter() *mux.Router
	SetHttpRouter(router *mux.Router)
	GetGrpcServer() *grpc.Server
	SetGrpcServer(server *grpc.Server)
	UseMiddleware(middleware MiddlewareFunc)
	AddModule(name string, module Module) error
	GetModule(name string) (Module, error)
	IsRunning() bool
}

func NewHost() Host {
	return &host{
		config:      newHostConfig(),
		middlewares: make([]MiddlewareFunc, 0),
		modules:     make(map[string]Module),
	}
}

type host struct {
	config       HostConfig
	configPath   string
	httpRouter   *mux.Router
	httpServer   *http.Server
	grpcServer   *grpc.Server
	grpcListener net.Listener
	middlewares  []MiddlewareFunc
	modules      map[string]Module
	isRunning    bool
}

func (host *host) Run() error {
	var err error

	// Check if not already running
	if host.isRunning {
		logrus.Error("Host already running")
		return ErrAlreadyRunning
	}

	// Set as running
	logrus.Info("Starting host")
	host.isRunning = true

	// Load the configuration
	configPath := host.getConfigPath()
	err = host.GetConfig().Load(configPath)
	if err != nil {
		return err
	}

	// Load the module configs
	err = host.loadModuleConfigs(configPath)
	if err != nil {
		return err
	}

	// Start the http server
	if host.httpRouter != nil {
		defer host.stopHttpServer()
		err = host.startHttpServer()
		if err != nil {
			logrus.Errorf("Failed to start the http server: %v", err)
			return err
		}
	}

	// Start the grpc server
	if host.grpcServer != nil {
		defer host.stopGrpcServer()
		err = host.startGrpcServer()
		if err != nil {
			logrus.Errorf("Failed to start the grpc server: %v", err)
			return err
		}
	}

	// Wait for cancel signal (CTRL+C)
	cancelSignal := make(chan os.Signal, 1)
	signal.Notify(cancelSignal, os.Interrupt)
	<-cancelSignal

	// Set as not running
	logrus.Info("Stopping host")
	host.isRunning = false

	return nil
}

func (host *host) GetConfig() HostConfig {
	return host.config
}

func (host *host) GetConfigPath() string {
	return host.configPath
}

func (host *host) SetConfigPath(configPath string) {
	host.configPath = configPath
}

func (host *host) GetHttpRouter() *mux.Router {
	if !host.isRunning && host.httpRouter == nil {
		host.httpRouter = mux.NewRouter().StrictSlash(true)
	}
	return host.httpRouter
}

func (host *host) SetHttpRouter(router *mux.Router) {
	if !host.isRunning {
		host.httpRouter = router
	}
}

func (host *host) GetGrpcServer() *grpc.Server {
	if !host.isRunning && host.grpcServer == nil {
		host.grpcServer = grpc.NewServer()
	}
	return host.grpcServer
}

func (host *host) SetGrpcServer(server *grpc.Server) {
	if !host.isRunning {
		host.grpcServer = server
	}
}

func (host *host) UseMiddleware(middleware MiddlewareFunc) {
	host.middlewares = append(host.middlewares, middleware)
}

func (host *host) AddModule(name string, module Module) error {
	_, found := host.modules[name]
	if found {
		return ErrModuleDuplicate
	}
	host.modules[name] = module
	return nil
}

func (host *host) GetModule(name string) (Module, error) {
	module, found := host.modules[name]
	if !found {
		return nil, ErrModuleNotFound
	}
	return module, nil
}

func (host *host) IsRunning() bool {
	return host.isRunning
}

func (host *host) getConfigPath() string {
	configPath := host.configPath
	if len(configPath) == 0 {
		configPath = os.Getenv("CONFIG_PATH")
	}
	return configPath
}

func (host *host) startHttpServer() error {
	httpServerAddress := host.config.GetHttpConfig().GetBindAddress()
	logrus.Infof("Starting http server: %s", httpServerAddress)

	// Register the api routes
	httpRouter := host.GetHttpRouter()
	err := host.registerModuleApiRoutes(httpRouter)
	if err != nil {
		return err
	}

	// Setup the http handler chain
	var httpHandler http.Handler
	httpHandler = host.GetHttpRouter()
	for _, middleware := range host.middlewares {
		httpHandler = middleware(httpHandler)
	}

	// Create the http server
	host.httpServer = &http.Server{
		Handler:      httpHandler,
		Addr:         httpServerAddress,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := host.httpServer.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				// Ignore
			} else {
				logrus.Fatalf("Server failed to run: %v", err)
			}
		}
	}()
	return nil
}

func (host *host) stopHttpServer() {
	logrus.Info("Stopping http server")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	host.httpServer.Shutdown(ctx)
}

func (host *host) startGrpcServer() error {
	grpcServerAddress := host.config.GetGrpcConfig().GetBindAddress()
	logrus.Infof("Starting grpc server: %s", grpcServerAddress)

	// Register the module grpc services
	server := host.GetGrpcServer()
	err := host.registerModuleGrpcServices(server)
	if err != nil {
		return err
	}

	// Run the grpc server
	listener, err := net.Listen("tcp", grpcServerAddress)
	if err != nil {
		return err
	}
	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (host *host) stopGrpcServer() {
	logrus.Info("Stopping grpc server")

	// Stop the GRPC server
	host.grpcServer.Stop()
	host.grpcListener.Close()
}

func (host *host) loadModuleConfigs(configPath string) error {
	var err error
	for _, module := range host.modules {
		err = module.LoadConfig(configPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (host *host) registerModuleApiRoutes(router *mux.Router) error {
	var err error
	for _, module := range host.modules {
		err = module.RegisterApiRoutes(router)
		if err != nil {
			return err
		}
	}
	return nil
}

func (host *host) registerModuleGrpcServices(server *grpc.Server) error {
	var err error
	for _, module := range host.modules {
		err = module.RegisterGrpcServices(server)
		if err != nil {
			return err
		}
	}
	return nil
}

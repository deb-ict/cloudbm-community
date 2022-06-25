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
	ErrAlreadyRunning error = errors.New("host already running")
)

type Host interface {
	GetConfig() HostConfig
	GetConfigPath() string
	SetConfigPath(configPath string)
	GetHttpHandler() http.Handler
	SetHttpHandler(handler http.Handler)
	GetGrpcServer() *grpc.Server
	SetGrpcServer(server *grpc.Server)
	Run() error
	IsRunning() bool
}

func NewHost() Host {
	return &host{
		config: newHostConfig(),
	}
}

type host struct {
	config       HostConfig
	configPath   string
	httpHandler  http.Handler
	httpServer   *http.Server
	grpcServer   *grpc.Server
	grpcListener net.Listener
	isRunning    bool
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

func (host *host) GetHttpHandler() http.Handler {
	if !host.isRunning && host.httpHandler == nil {
		host.httpHandler = mux.NewRouter().StrictSlash(true)
	}
	return host.httpHandler
}

func (host *host) SetHttpHandler(handler http.Handler) {
	if !host.isRunning {
		host.httpHandler = handler
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

	// Start the http server
	if host.httpHandler != nil {
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
	logrus.Info("Starting http server: %s", httpServerAddress)

	// Create the http server
	host.httpServer = &http.Server{
		Handler:      host.GetHttpHandler(),
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
	logrus.Info("Starting grpc server: %s", grpcServerAddress)

	listener, err := net.Listen("tcp", grpcServerAddress)
	if err != nil {
		return err
	}
	if err := host.GetGrpcServer().Serve(listener); err != nil {
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

package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	auth_api_v1 "github.com/deb-ict/cloudbm-community/pkg/module/auth/api/v1"
	auth_svc "github.com/deb-ict/cloudbm-community/pkg/module/auth/service"
	contact_api_v1 "github.com/deb-ict/cloudbm-community/pkg/module/contact/api/v1"
	contact_svc "github.com/deb-ict/cloudbm-community/pkg/module/contact/service"
	gallery_api_v1 "github.com/deb-ict/cloudbm-community/pkg/module/gallery/api/v1"
	gallery_svc "github.com/deb-ict/cloudbm-community/pkg/module/gallery/service"
	product_api_v1 "github.com/deb-ict/cloudbm-community/pkg/module/product/api/v1"
	product_svc "github.com/deb-ict/cloudbm-community/pkg/module/product/service"
	session_api_v1 "github.com/deb-ict/cloudbm-community/pkg/module/session/api/v1"
	session_svc "github.com/deb-ict/cloudbm-community/pkg/module/session/service"
	"github.com/gorilla/mux"
)

func main() {
	// Parse arguments
	var configPath string
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&configPath, "config", "/etc/cloudbm/portal.conf", "the path of the configuration file")
	flag.Parse()

	// Setup the default logger
	slogJsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	slog.SetDefault(slog.New(slogJsonHandler))
	slog.SetLogLoggerLevel(slog.LevelInfo)

	// Load configuration
	config, err := LoadConfig(configPath)
	if err != nil {
		os.Exit(1)
	}

	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Setup the HTTP server and routes
	router := mux.NewRouter().StrictSlash(true)
	initServices(router)

	// Start the HTTP server
	httpServerAddr := config.Http.GetBindAddress()
	slog.InfoContext(context.Background(), "Starting http server",
		slog.String("address", httpServerAddr),
	)
	httpServer := &http.Server{
		Handler:      router,
		Addr:         httpServerAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	httpServerErr := make(chan error, 1)
	go func() {
		httpServerErr <- httpServer.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err := <-httpServerErr:
		// Error when starting HTTP server.
		slog.ErrorContext(context.Background(), "Error while running http server",
			slog.Any("error", err),
		)
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		slog.InfoContext(context.Background(), "Application termination requested")
		stop()
	}

	// Stop http server
	slog.InfoContext(context.Background(), "Stopping http server")
	httpServer.Shutdown(ctx)

	// Shutdown
	os.Exit(0)
}

func initServices(router *mux.Router) {
	//TODO: Database connection
	authSvc := auth_svc.NewService(nil, &auth_svc.ServiceOptions{})
	authApiV1 := auth_api_v1.NewApiV1(authSvc)
	authApiV1.RegisterRoutes(router.PathPrefix("auth").Subrouter())

	//TODO: Database connection
	gallerySvc := gallery_svc.NewService(nil, &gallery_svc.ServiceOptions{})
	galleryApiV1 := gallery_api_v1.NewApiV1(gallerySvc)
	galleryApiV1.RegisterRoutes(router.PathPrefix("gallery").Subrouter())

	//TODO: Database connection
	contactSvc := contact_svc.NewService(nil, &contact_svc.ServiceOptions{})
	contactApiV1 := contact_api_v1.NewApiV1(contactSvc)
	contactApiV1.RegisterRoutes(router.PathPrefix("contact").Subrouter())

	//TODO: Database connection
	productSvc := product_svc.NewService(nil, &product_svc.ServiceOptions{})
	productApiV1 := product_api_v1.NewApiV1(productSvc)
	productApiV1.RegisterRoutes(router.PathPrefix("product").Subrouter())

	//TODO: Database connection
	sessionSvc := session_svc.NewService(nil, &session_svc.ServiceOptions{})
	sessionApiV1 := session_api_v1.NewApiV1(sessionSvc)
	sessionApiV1.RegisterRoutes(router.PathPrefix("session").Subrouter())
}

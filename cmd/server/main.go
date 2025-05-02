package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
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
	"github.com/sirupsen/logrus"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	// Parse arguments
	var configPath string
	var wait time.Duration
	var err error
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&configPath, "config", "/etc/cloudbm/cloudbm.conf", "the path of the configuration file")
	flag.Parse()

	// Load configuration
	config, err = LoadConfig(configPath)
	if err != nil {
		logrus.Fatalf("Failed to load config (%s): %v", configPath, err)
	}

	// Initialize the router
	httpRouter := mux.NewRouter().StrictSlash(true)

	// Setup the
	spaHandler := spaHandler{staticPath: "web/dist/browser", indexPath: "index.html"}
	httpRouter.PathPrefix("/").Handler(spaHandler)

	authSvc := auth_svc.NewService(nil, &auth_svc.ServiceOptions{})
	authApiV1 := auth_api_v1.NewApiV1(authSvc)
	authApiV1.RegisterRoutes(httpRouter.PathPrefix("auth").Subrouter())

	gallerySvc := gallery_svc.NewService(nil, &gallery_svc.ServiceOptions{})
	galleryApiV1 := gallery_api_v1.NewApiV1(gallerySvc)
	galleryApiV1.RegisterRoutes(httpRouter.PathPrefix("gallery").Subrouter())

	contactSvc := contact_svc.NewService(nil, &contact_svc.ServiceOptions{})
	contactApiV1 := contact_api_v1.NewApiV1(contactSvc)
	contactApiV1.RegisterRoutes(httpRouter.PathPrefix("contact").Subrouter())

	productSvc := product_svc.NewService(nil, &product_svc.ServiceOptions{})
	productApiV1 := product_api_v1.NewApiV1(productSvc)
	productApiV1.RegisterRoutes(httpRouter.PathPrefix("product").Subrouter())

	sessionSvc := session_svc.NewService(nil, &session_svc.ServiceOptions{})
	sessionApiV1 := session_api_v1.NewApiV1(sessionSvc)
	sessionApiV1.RegisterRoutes(httpRouter.PathPrefix("session").Subrouter())

	// Run the web server
	httpServerAddress := config.GetHttpConfig().GetBindAddress()
	logrus.Infof("Starting http server: %s", httpServerAddress)
	httpServer := &http.Server{
		Handler:      httpRouter,
		Addr:         httpServerAddress,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				// Ignore
			} else {
				logrus.Fatalf("Failed to run http server: %v", err)
			}
		}
	}()

	// Create interruption signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Stop http server
	logrus.Info("Stopping http server")
	httpServer.Shutdown(ctx)

	// Shutdown
	os.Exit(0)
}

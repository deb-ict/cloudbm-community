package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/deb-ict/cloudbm-community/pkg/authentication"
	"github.com/deb-ict/cloudbm-community/pkg/authorization"
	"github.com/deb-ict/cloudbm-community/pkg/logging"
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
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func main() {
	// Parse arguments
	var configPath string
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&configPath, "config", "/etc/cloudbm/cloudbm.conf", "the path of the configuration file")
	flag.Parse()

	// Setup the default logger
	slogJsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	slog.SetDefault(slog.New(slogJsonHandler))
	slog.SetLogLoggerLevel(slog.LevelInfo)

	claims := jwt.MapClaims{}
	claims["sub"] = "1234567890"
	claims["iss"] = "https://localhost:8000"
	claims["aud"] = "cloudbm-users"
	claims["name"] = "John Doe"
	claims["iat"] = jwt.NewNumericDate(time.Now().UTC())
	claims["nbf"] = jwt.NewNumericDate(time.Now().UTC())
	claims["exp"] = jwt.NewNumericDate(time.Now().UTC().Add(1 * time.Hour))
	claims["jti"] = "unique-token-id"
	claims["role"] = "user admin"
	claims["scope"] = "product:read product:write"
	claims["email"] = "john.doe@deb-ict.com"
	claims["email_verified"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your-256-bit-secret"))
	if err != nil {
		slog.ErrorContext(context.Background(), "Failed to generate JWT token",
			slog.Any("error", err),
		)
	}
	slog.InfoContext(context.Background(), "Generated JWT token",
		slog.String("token", tokenString),
	)

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
	registerAuthService(router, &config.AuthService)
	registerGalleryService(router, &config.GalleryService)
	registerContactService(router, &config.ContactService)
	registerProductService(router, &config.ProductService)
	registerSessionService(router, &config.SessionService)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to CloudBM!"))
	})

	// Setup the middlewares
	//router.Use(logging.NewMiddleware().Middleware)
	router.Use(authentication.NewMiddleware(nil).Middleware)
	router.Use(authorization.NewMiddleware().Middleware)
	//router.Use(metrics) // prometheus metrics middleware
	//router.Use(cors)
	//router.Use(tracing) // otel tracing middleware

	customerHeaderMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					slog.ErrorContext(context.Background(), "Panic occurred while handling request",
						slog.Any("error", r),
					)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()

			w.Header().Add("Server", "CBM-WebHost/1.0")
			w.Header().Add("X-Powered-By", "https://www.deb-ict.com")
			next.ServeHTTP(w, r)
		})
	}
	httpHandler := logging.NewMiddleware().Middleware(router)
	httpHandler = customerHeaderMiddleware(httpHandler)

	// Start the HTTP server
	httpServerAddr := config.Http.GetBindAddress()
	slog.InfoContext(context.Background(), "Starting http server",
		slog.String("address", httpServerAddr),
	)
	httpServer := &http.Server{
		Handler:      httpHandler,
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

func registerAuthService(router *mux.Router, opts *auth_svc.ServiceOptions) {
	authSvc := auth_svc.NewService(nil, opts)
	authApiV1 := auth_api_v1.NewApiV1(authSvc)
	authApiV1.RegisterRoutes(router.PathPrefix("auth").Subrouter())
}

func registerGalleryService(router *mux.Router, opts *gallery_svc.ServiceOptions) {
	gallerySvc := gallery_svc.NewService(nil, opts)
	galleryApiV1 := gallery_api_v1.NewApiV1(gallerySvc)
	galleryApiV1.RegisterRoutes(router.PathPrefix("/gallery").Subrouter())
}

func registerContactService(router *mux.Router, opts *contact_svc.ServiceOptions) {
	contactSvc := contact_svc.NewService(nil, opts)
	contactApiV1 := contact_api_v1.NewApiV1(contactSvc)
	contactApiV1.RegisterRoutes(router.PathPrefix("/contact").Subrouter())
}

func registerProductService(router *mux.Router, opts *product_svc.ServiceOptions) {
	productSvc := product_svc.NewService(nil, opts)
	productApiV1 := product_api_v1.NewApiV1(productSvc)
	productApiV1.RegisterRoutes(router.PathPrefix("/product").Subrouter())
}

func registerSessionService(router *mux.Router, opts *session_svc.ServiceOptions) {
	sessionSvc := session_svc.NewService(nil, opts)
	sessionApiV1 := session_api_v1.NewApiV1(sessionSvc)
	sessionApiV1.RegisterRoutes(router.PathPrefix("/session").Subrouter())
}

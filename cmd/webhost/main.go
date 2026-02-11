package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/deb-ict/cloudbm-community/pkg/logging"
	auth_api_v1 "github.com/deb-ict/cloudbm-community/pkg/module/auth/api/v1"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/database"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/model"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/oauth"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/security"
	auth_svc "github.com/deb-ict/cloudbm-community/pkg/module/auth/service"
	contact_api_v1 "github.com/deb-ict/cloudbm-community/pkg/module/contact/api/v1"
	contact_svc "github.com/deb-ict/cloudbm-community/pkg/module/contact/service"
	gallery_api_v1 "github.com/deb-ict/cloudbm-community/pkg/module/gallery/api/v1"
	gallery_svc "github.com/deb-ict/cloudbm-community/pkg/module/gallery/service"
	product_api_v1 "github.com/deb-ict/cloudbm-community/pkg/module/product/api/v1"
	product_svc "github.com/deb-ict/cloudbm-community/pkg/module/product/service"
	session_api_v1 "github.com/deb-ict/cloudbm-community/pkg/module/session/api/v1"
	session_svc "github.com/deb-ict/cloudbm-community/pkg/module/session/service"
	"github.com/deb-ict/go-router"
	"github.com/deb-ict/go-router/authentication"
	"github.com/deb-ict/go-router/authorization"
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

	// Load configuration
	config, err := LoadConfig(configPath)
	if err != nil {
		os.Exit(1)
	}

	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Initialize the middlewares
	authenticationValidator := oauth.NewTokenValidator()
	authenticationHandler := authentication.NewBearerAuthenticationHandler(authenticationValidator)
	authenticationMiddleware := authentication.NewMiddleware(authenticationHandler)
	authorizationMiddleware := authorization.NewMiddleware()

	// Setup the HTTP server and routes
	router := router.NewRouter()
	registerAuthService(router, authorizationMiddleware, &config.AuthService)
	registerGalleryService(router, authorizationMiddleware, &config.GalleryService)
	registerContactService(router, authorizationMiddleware, &config.ContactService)
	registerProductService(router, authorizationMiddleware, &config.ProductService)
	registerSessionService(router, authorizationMiddleware, &config.SessionService)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to CloudBM!"))
	})

	// Setup the authentication middleware
	router.Use(authenticationMiddleware.Middleware)

	// Setup the authorization middleware
	router.Use(authorizationMiddleware.Middleware)

	// Setup the middlewares
	//router.Use(logging.NewMiddleware().Middleware)
	//router.Use(authentication.NewMiddleware(nil).Middleware)
	//router.Use(authorization.NewMiddleware().Middleware)
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

func registerAuthService(router *router.Router, authorization *authorization.Middleware, opts *auth_svc.ServiceOptions) {
	authDb, err := database.NewDatabase()
	if err != nil {
		slog.ErrorContext(context.Background(), "Failed to initialize auth database",
			slog.Any("error", err),
		)
		os.Exit(1)
	}

	authSvc := auth_svc.NewService(authDb, opts)
	authApiV1 := auth_api_v1.NewApiV1(authSvc)
	authApiV1.RegisterAuthorizationPolicies(authorization)
	authApiV1.RegisterRoutes(router.PathPrefix("/api/auth").SubRouter())

	tokenHandler := oauth.NewTokenHandler(authSvc, "http://localhost:8000")
	tokenHandler.RegisterRoutes(router.PathPrefix("/oauth").SubRouter())

	_, count, err := authSvc.GetUsers(context.Background(), 0, 1, nil, nil)
	if err != nil {
		slog.ErrorContext(context.Background(), "Failed to get users from auth service",
			slog.Any("error", err),
		)
		os.Exit(1)
	}
	if count == 0 {
		password := security.GeneratePassword(16)
		hash, err := authSvc.PasswordHasher().HashPassword(password)
		if err != nil {
			slog.ErrorContext(context.Background(), "Failed to hash password",
				slog.Any("error", err),
			)
			os.Exit(1)
		}
		user := &model.User{
			Username:      "admin",
			Email:         "admin@cloudbm.eu",
			EmailVerified: true,
			PasswordHash:  hash,
			IsEnabled:     true,
			LockEnd:       time.Now().UTC(),
		}
		_, err = authSvc.CreateUser(context.Background(), user)
		if err != nil {
			slog.ErrorContext(context.Background(), "Failed to create admin user",
				slog.Any("error", err),
			)
			os.Exit(1)
		}

		slog.WarnContext(context.Background(), "Admin user created. Store the password in a safe place and change it after first login.",
			slog.String("username", user.Username),
			slog.String("email", user.Email),
			slog.String("password", password),
		)
	}
}

func registerGalleryService(router *router.Router, authorization *authorization.Middleware, opts *gallery_svc.ServiceOptions) {
	gallerySvc := gallery_svc.NewService(nil, opts)
	galleryApiV1 := gallery_api_v1.NewApiV1(gallerySvc)
	galleryApiV1.RegisterAuthorizationPolicies(authorization)
	galleryApiV1.RegisterRoutes(router.PathPrefix("/api/gallery").SubRouter())
}

func registerContactService(router *router.Router, authorization *authorization.Middleware, opts *contact_svc.ServiceOptions) {
	contactSvc := contact_svc.NewService(nil, opts)
	contactApiV1 := contact_api_v1.NewApiV1(contactSvc)
	contactApiV1.RegisterAuthorizationPolicies(authorization)
	contactApiV1.RegisterRoutes(router.PathPrefix("/api/contact").SubRouter())
}

func registerProductService(router *router.Router, authorization *authorization.Middleware, opts *product_svc.ServiceOptions) {
	productSvc := product_svc.NewService(nil, opts)
	productApiV1 := product_api_v1.NewApiV1(productSvc)
	productApiV1.RegisterAuthorizationPolicies(authorization)
	productApiV1.RegisterRoutes(router.PathPrefix("/api/product").SubRouter())
}

func registerSessionService(router *router.Router, authorization *authorization.Middleware, opts *session_svc.ServiceOptions) {
	sessionSvc := session_svc.NewService(nil, opts)
	sessionApiV1 := session_api_v1.NewApiV1(sessionSvc)
	sessionApiV1.RegisterAuthorizationPolicies(authorization)
	sessionApiV1.RegisterRoutes(router.PathPrefix("/api/session").SubRouter())
}

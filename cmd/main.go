package main

import (
	"log"
	"net/http"
	"time"

	auth_api_v1 "github.com/deb-ict/cloudbm-community/pkg/module/auth/api/v1"
	auth_svc "github.com/deb-ict/cloudbm-community/pkg/module/auth/service"
	contact_api_v1 "github.com/deb-ict/cloudbm-community/pkg/module/contact/api/v1"
	contact_svc "github.com/deb-ict/cloudbm-community/pkg/module/contact/service"
	gallery_api_v1 "github.com/deb-ict/cloudbm-community/pkg/module/gallery/api/v1"
	gallery_svc "github.com/deb-ict/cloudbm-community/pkg/module/gallery/service"
	product_api_v1 "github.com/deb-ict/cloudbm-community/pkg/module/product/api/v1"
	product_svc "github.com/deb-ict/cloudbm-community/pkg/module/product/service"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	authSvc := auth_svc.NewService(nil, &auth_svc.ServiceOptions{})
	authApiV1 := auth_api_v1.NewApiV1(authSvc)
	authApiV1.RegisterRoutes(r.PathPrefix("auth").Subrouter())

	gallerySvc := gallery_svc.NewService(nil, &gallery_svc.ServiceOptions{})
	galleryApiV1 := gallery_api_v1.NewApiV1(gallerySvc)
	galleryApiV1.RegisterRoutes(r.PathPrefix("gallery").Subrouter())

	contactSvc := contact_svc.NewService(nil, &contact_svc.ServiceOptions{})
	contactApiV1 := contact_api_v1.NewApiV1(contactSvc)
	contactApiV1.RegisterRoutes(r.PathPrefix("contact").Subrouter())

	productSvc := product_svc.NewService(nil, &product_svc.ServiceOptions{})
	productApiV1 := product_api_v1.NewApiV1(productSvc)
	productApiV1.RegisterRoutes(r.PathPrefix("product").Subrouter())

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

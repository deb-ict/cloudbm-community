package main

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/storage/mongodb"
	"github.com/deb-ict/cloudbm-community/pkg/user"
	"github.com/gorilla/mux"
)

type ApiRouter struct {
	*mux.Router
}

func (r *ApiRouter) RegisterApi() *ApiRouter {
	return r
}

func main() {
	db := mongodb.NewDatabase()

	router := mux.NewRouter().StrictSlash(true)
	r := &ApiRouter{router}
	r.RegisterApi()

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	user.NewMuxRouterApi(apiRouter, user.NewService(db.GetUserRepository()))

	http.ListenAndServe("127.0.0.1:5000", router)
}

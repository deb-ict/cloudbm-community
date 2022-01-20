package main

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/storage/mongodb"
	"github.com/deb-ict/cloudbm-community/pkg/user"
	"github.com/gorilla/mux"
)

func main() {
	db := mongodb.NewDatabase()

	userApi := user.NewApi(user.NewService(db.GetUserRepository()))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/user", userApi.GetUsers).Methods("GET")

	http.ListenAndServe("127.0.0.1:5000", router)
}

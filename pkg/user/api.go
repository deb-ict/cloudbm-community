package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Api interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type api struct {
	svc Service
}

func NewApi(s Service) Api {
	return &api{
		svc: s,
	}
}

func NewMuxRouterApi(r *mux.Router, s Service) Api {
	api := NewApi(s).(*api)
	api.RegisterRoutes(r)
	return api
}

func (api *api) RegisterRoutes(r *mux.Router) {
	//r.HandleFunc("/user", api.CreateUser).Methods("POST")
	r.HandleFunc("/user", api.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", api.GetUserById).Methods("GET")
	//r.HandleFunc("/user/{id}", api.UpdateUser).Methods("PUT", "PATCH")
	//r.HandleFunc("/user/{id}", api.DeleteUser).Methods("DELETE")
}

func (api *api) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := api.svc.GetUsers(r.Context(), 0, 25)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (api *api) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result, err := api.svc.GetUserById(r.Context(), id)
	if err == ErrorNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (api *api) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (api *api) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (api *api) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

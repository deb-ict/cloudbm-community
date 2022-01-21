package user

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/gorilla/mux"
)

type Api interface {
	RegisterRoutes(r *mux.Router)
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
	r.HandleFunc("/api/v1/user", api.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/user", api.GetUsers).Methods("GET")
	r.HandleFunc("/api/v1/user/{id}", api.GetUserById).Methods("GET")
	r.HandleFunc("/api/v1/user/{id}", api.UpdateUser).Methods("PUT", "PATCH")
	r.HandleFunc("/api/v1/user/{id}", api.DeleteUser).Methods("DELETE")
}

func (api *api) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := api.svc.GetUsers(r.Context(), 0, 25)
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	rest.WriteResult(w, result)
}

func (api *api) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result, err := api.svc.GetUserById(r.Context(), id)
	if err == ErrNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *api) CreateUser(w http.ResponseWriter, r *http.Request) {
	var model User
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.svc.CreateUser(r.Context(), model)
	if err == ErrDuplicateUserName {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err == ErrDuplicateEmail {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *api) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var model User
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	model.UserName = ""

	result, err := api.svc.UpdateUser(r.Context(), id, model)
	if err == ErrNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *api) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := api.svc.DeleteUser(r.Context(), id)
	if err == ErrNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

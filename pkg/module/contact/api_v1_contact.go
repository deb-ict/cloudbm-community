package contact

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/gorilla/mux"
)

func (api *apiHandler) GetContacts(w http.ResponseWriter, r *http.Request) {
	paging := rest.GetPaging(r)

	result, err := api.service.GetContacts(r.Context(), paging.PageIndex, paging.PageSize)
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	rest.WriteResult(w, result)
}

func (api *apiHandler) GetContactById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result, err := api.service.GetContactById(r.Context(), id)
	if err == ErrContactNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	var model Contact
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.CreateContact(r.Context(), model)
	if err == ErrContactDuplicateName {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var model Contact
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.UpdateContact(r.Context(), id, model)
	if err == ErrContactNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := api.service.DeleteContact(r.Context(), id)
	if err == ErrContactNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

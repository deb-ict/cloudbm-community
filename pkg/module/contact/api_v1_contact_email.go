package contact

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/gorilla/mux"
)

func (api *apiHandler) GetContactEmails(w http.ResponseWriter, r *http.Request) {
	contactId := mux.Vars(r)["contactId"]
	paging := rest.GetPaging(r)

	result, err := api.service.GetContactEmails(r.Context(), contactId, paging.PageIndex, paging.PageSize)
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

func (api *apiHandler) GetContactEmailById(w http.ResponseWriter, r *http.Request) {
	contactId := mux.Vars(r)["contactId"]
	id := mux.Vars(r)["id"]

	result, err := api.service.GetContactEmailById(r.Context(), contactId, id)
	if err == ErrContactNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrContactEmailNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) CreateContactEmail(w http.ResponseWriter, r *http.Request) {
	contactId := mux.Vars(r)["contactId"]

	var model Email
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.CreateContactEmail(r.Context(), contactId, model)
	if err == ErrContactNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrContactEmailDuplicate {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) UpdateContactEmail(w http.ResponseWriter, r *http.Request) {
	contactId := mux.Vars(r)["contactId"]
	id := mux.Vars(r)["id"]

	var model Email
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.UpdateContactEmail(r.Context(), contactId, id, model)
	if err == ErrContactNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrContactEmailNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) DeleteContactEmail(w http.ResponseWriter, r *http.Request) {
	contactId := mux.Vars(r)["contactId"]
	id := mux.Vars(r)["id"]

	err := api.service.DeleteContactEmail(r.Context(), contactId, id)
	if err == ErrContactNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrContactEmailNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

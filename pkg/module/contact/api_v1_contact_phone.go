package contact

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/gorilla/mux"
)

func (api *apiHandler) GetContactPhones(w http.ResponseWriter, r *http.Request) {
	contactId := mux.Vars(r)["contactId"]
	paging := rest.GetPaging(r)

	result, err := api.service.GetContactPhones(r.Context(), contactId, paging.PageIndex, paging.PageSize)
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

func (api *apiHandler) GetContactPhoneById(w http.ResponseWriter, r *http.Request) {
	contactId := mux.Vars(r)["contactId"]
	id := mux.Vars(r)["id"]

	result, err := api.service.GetContactPhoneById(r.Context(), contactId, id)
	if err == ErrContactNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrContactPhoneNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) CreateContactPhone(w http.ResponseWriter, r *http.Request) {
	contactId := mux.Vars(r)["contactId"]
	var model Phone
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.CreateContactPhone(r.Context(), contactId, model)
	if err == ErrContactNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrContactPhoneDuplicate {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) UpdateContactPhone(w http.ResponseWriter, r *http.Request) {
	contactId := mux.Vars(r)["contactId"]
	id := mux.Vars(r)["id"]

	var model Phone
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.UpdateContactPhone(r.Context(), contactId, id, model)
	if err == ErrContactNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrContactPhoneNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) DeleteContactPhone(w http.ResponseWriter, r *http.Request) {
	contactId := mux.Vars(r)["contactId"]
	id := mux.Vars(r)["id"]

	err := api.service.DeleteContactPhone(r.Context(), contactId, id)
	if err == ErrContactNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrContactPhoneNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

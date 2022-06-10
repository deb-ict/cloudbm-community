package contact

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/gorilla/mux"
)

func (api *apiHandler) GetContactAddresses(w http.ResponseWriter, r *http.Request) {
	contactId := mux.Vars(r)["contactId"]
	paging := rest.GetPaging(r)

	result, err := api.service.GetContactAddresses(r.Context(), contactId, paging.PageIndex, paging.PageSize)
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

func (api *apiHandler) GetContactAddressById(w http.ResponseWriter, r *http.Request) {
	contactId := mux.Vars(r)["contactId"]
	id := mux.Vars(r)["id"]

	result, err := api.service.GetContactAddressById(r.Context(), contactId, id)
	if err == ErrContactNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrContactAddressNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) CreateContactAddress(w http.ResponseWriter, r *http.Request) {
	contactId := mux.Vars(r)["contactId"]

	var model Address
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.CreateContactAddress(r.Context(), contactId, model)
	if err == ErrContactNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrContactAddressDuplicate {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) UpdateContactAddress(w http.ResponseWriter, r *http.Request) {
	contactId := mux.Vars(r)["contactId"]
	id := mux.Vars(r)["id"]

	var model Address
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.UpdateContactAddress(r.Context(), contactId, id, model)
	if err == ErrContactNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrContactAddressNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) DeleteContactAddress(w http.ResponseWriter, r *http.Request) {
	contactId := mux.Vars(r)["contactId"]
	id := mux.Vars(r)["id"]

	err := api.service.DeleteContactAddress(r.Context(), contactId, id)
	if err == ErrContactNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrContactAddressNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

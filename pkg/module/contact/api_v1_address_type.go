package contact

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/gorilla/mux"
)

func (api *apiHandler) GetAddressTypes(w http.ResponseWriter, r *http.Request) {
	paging := rest.GetPaging(r)

	result, err := api.service.GetAddressTypes(r.Context(), paging.PageIndex, paging.PageSize)
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	rest.WriteResult(w, result)
}

func (api *apiHandler) GetAddressTypeById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result, err := api.service.GetAddressTypeById(r.Context(), id)
	if err == ErrAddressTypeNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) CreateAddressType(w http.ResponseWriter, r *http.Request) {
	var model AddressType
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	model.IsSystem = false

	result, err := api.service.CreateAddressType(r.Context(), model)
	if err == ErrAddressTypeDuplicate {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) UpdateAddressType(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var model AddressType
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.UpdateAddressType(r.Context(), id, model)
	if err == ErrAddressTypeNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrAddressTypeReadOnly {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) DeleteAddressType(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := api.service.DeleteAddressType(r.Context(), id)
	if err == ErrAddressTypeNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrAddressTypeReadOnly {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiHandler) SetAddressTypeAsDefault(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := api.service.SetDefaultAddressType(r.Context(), id)
	if err == ErrAddressTypeNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

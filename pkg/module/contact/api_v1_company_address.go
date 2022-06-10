package contact

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/gorilla/mux"
)

func (api *apiHandler) GetCompanyAddresses(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	paging := rest.GetPaging(r)

	result, err := api.service.GetCompanyAddresses(r.Context(), companyId, paging.PageIndex, paging.PageSize)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	rest.WriteResult(w, result)
}

func (api *apiHandler) GetCompanyAddressById(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	id := mux.Vars(r)["id"]

	result, err := api.service.GetCompanyAddressById(r.Context(), companyId, id)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrCompanyAddressNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) CreateCompanyAddress(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]

	var model Address
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.CreateCompanyAddress(r.Context(), companyId, model)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrCompanyAddressDuplicate {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) UpdateCompanyAddress(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	id := mux.Vars(r)["id"]

	var model Address
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.UpdateCompanyAddress(r.Context(), companyId, id, model)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrCompanyAddressNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) DeleteCompanyAddress(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	id := mux.Vars(r)["id"]

	err := api.service.DeleteCompanyAddress(r.Context(), companyId, id)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrCompanyAddressNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

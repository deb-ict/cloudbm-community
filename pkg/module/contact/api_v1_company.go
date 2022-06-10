package contact

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/gorilla/mux"
)

func (api *apiHandler) GetCompanies(w http.ResponseWriter, r *http.Request) {
	paging := rest.GetPaging(r)

	result, err := api.service.GetCompanies(r.Context(), paging.PageIndex, paging.PageSize)
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	rest.WriteResult(w, result)
}

func (api *apiHandler) GetCompanyById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result, err := api.service.GetCompanyById(r.Context(), id)
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

func (api *apiHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var model Company
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.CreateCompany(r.Context(), model)
	if err == ErrCompanyDuplicateName {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var model Company
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.UpdateCompany(r.Context(), id, model)
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

func (api *apiHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := api.service.DeleteCompany(r.Context(), id)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

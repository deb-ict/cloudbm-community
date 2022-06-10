package contact

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/gorilla/mux"
)

func (api *apiHandler) GetCompanyUrls(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	paging := rest.GetPaging(r)

	result, err := api.service.GetCompanyUrls(r.Context(), companyId, paging.PageIndex, paging.PageSize)
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

func (api *apiHandler) GetCompanyUrlById(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	id := mux.Vars(r)["id"]

	result, err := api.service.GetCompanyUrlById(r.Context(), companyId, id)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrCompanyUrlNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) CreateCompanyUrl(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]

	var model Url
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.CreateCompanyUrl(r.Context(), companyId, model)
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

func (api *apiHandler) UpdateCompanyUrl(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	id := mux.Vars(r)["id"]

	var model Url
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.UpdateCompanyUrl(r.Context(), companyId, id, model)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrCompanyUrlNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) DeleteCompanyUrl(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	id := mux.Vars(r)["id"]

	err := api.service.DeleteCompanyUrl(r.Context(), companyId, id)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrCompanyUrlNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

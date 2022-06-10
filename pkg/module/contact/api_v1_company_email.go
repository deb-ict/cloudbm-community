package contact

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/gorilla/mux"
)

func (api *apiHandler) GetCompanyEmails(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	paging := rest.GetPaging(r)

	result, err := api.service.GetCompanyEmails(r.Context(), companyId, paging.PageIndex, paging.PageSize)
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

func (api *apiHandler) GetCompanyEmailById(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	id := mux.Vars(r)["id"]

	result, err := api.service.GetCompanyEmailById(r.Context(), companyId, id)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrCompanyEmailNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) CreateCompanyEmail(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]

	var model Email
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.CreateCompanyEmail(r.Context(), companyId, model)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrCompanyEmailDuplicate {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) UpdateCompanyEmail(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	id := mux.Vars(r)["id"]

	var model Email
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.UpdateCompanyEmail(r.Context(), companyId, id, model)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrCompanyEmailNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) DeleteCompanyEmail(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	id := mux.Vars(r)["id"]

	err := api.service.DeleteCompanyEmail(r.Context(), companyId, id)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrCompanyEmailNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

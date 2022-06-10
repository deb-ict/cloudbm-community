package contact

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/gorilla/mux"
)

func (api *apiHandler) GetCompanyPhones(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	paging := rest.GetPaging(r)

	result, err := api.service.GetCompanyPhones(r.Context(), companyId, paging.PageIndex, paging.PageSize)
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

func (api *apiHandler) GetCompanyPhoneById(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	id := mux.Vars(r)["id"]

	result, err := api.service.GetCompanyPhoneById(r.Context(), companyId, id)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrCompanyPhoneNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) CreateCompanyPhone(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]

	var model Phone
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.CreateCompanyPhone(r.Context(), companyId, model)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrCompanyPhoneDuplicate {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) UpdateCompanyPhone(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	id := mux.Vars(r)["id"]

	var model Phone
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.UpdateCompanyPhone(r.Context(), companyId, id, model)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrCompanyPhoneNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) DeleteCompanyPhone(w http.ResponseWriter, r *http.Request) {
	companyId := mux.Vars(r)["companyId"]
	id := mux.Vars(r)["id"]

	err := api.service.DeleteCompanyPhone(r.Context(), companyId, id)
	if err == ErrCompanyNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrCompanyPhoneNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

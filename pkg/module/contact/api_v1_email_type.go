package contact

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/gorilla/mux"
)

func (api *apiHandler) GetEmailTypes(w http.ResponseWriter, r *http.Request) {
	paging := rest.GetPaging(r)

	result, err := api.service.GetEmailTypes(r.Context(), paging.PageIndex, paging.PageSize)
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	rest.WriteResult(w, result)
}

func (api *apiHandler) GetEmailTypeById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result, err := api.service.GetEmailTypeById(r.Context(), id)
	if err == ErrEmailTypeNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) CreateEmailType(w http.ResponseWriter, r *http.Request) {
	var model EmailType
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.CreateEmailType(r.Context(), model)
	if err == ErrEmailTypeDuplicate {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) UpdateEmailType(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var model EmailType
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.UpdateEmailType(r.Context(), id, model)
	if err == ErrEmailTypeNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrEmailTypeReadOnly {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) DeleteEmailType(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := api.service.DeleteEmailType(r.Context(), id)
	if err == ErrEmailTypeNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrEmailTypeReadOnly {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

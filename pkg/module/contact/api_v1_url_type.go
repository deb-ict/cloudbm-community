package contact

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/gorilla/mux"
)

func (api *apiHandler) GetUrlTypes(w http.ResponseWriter, r *http.Request) {
	paging := rest.GetPaging(r)

	result, err := api.service.GetUrlTypes(r.Context(), paging.PageIndex, paging.PageSize)
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	rest.WriteResult(w, result)
}

func (api *apiHandler) GetUrlTypeById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result, err := api.service.GetUrlTypeById(r.Context(), id)
	if err == ErrUrlTypeNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) CreateUrlType(w http.ResponseWriter, r *http.Request) {
	var model UrlType
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.CreateUrlType(r.Context(), model)
	if err == ErrUrlTypeDuplicate {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) UpdateUrlType(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var model UrlType
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := api.service.UpdateUrlType(r.Context(), id, model)
	if err == ErrUrlTypeNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrUrlTypeReadOnly {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteResult(w, result)
}

func (api *apiHandler) DeleteUrlType(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := api.service.DeleteUrlType(r.Context(), id)
	if err == ErrUrlTypeNotFound {
		rest.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	if err == ErrUrlTypeReadOnly {
		rest.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

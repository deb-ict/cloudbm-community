package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/gallery"
	"github.com/gorilla/mux"
)

type ApiV1 interface {
	RegisterRoutes(r *mux.Router)
}

type apiV1 struct {
	service gallery.Service
}

func NewApiV1(service gallery.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterRoutes(r *mux.Router) {
	// Images
	/*
		r.HandleFunc("/v1/image", api.GetImagesHandlerV1).Methods(http.MethodGet)
		r.HandleFunc("/v1/image/{id}", api.GetImageByIdHandlerV1).Methods(http.MethodGet)
		r.HandleFunc("/v1/image", api.CreateImageHandlerV1).Methods(http.MethodPost)
		r.HandleFunc("/v1/image/{id}", api.UpdateImageHandlerV1).Methods(http.MethodPut)
		r.HandleFunc("/v1/image/{id}", api.DeleteImageHandlerV1).Methods(http.MethodDelete)
	*/
}

func (api *apiV1) handleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	switch err {
	case gallery.ErrImageNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case core.ErrTranslationNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case core.ErrInvalidId:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	default:
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
	}
	return true
}

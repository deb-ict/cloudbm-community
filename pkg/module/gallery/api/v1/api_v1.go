package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/gallery"
	"github.com/deb-ict/go-router"
	"github.com/deb-ict/go-router/authorization"
)

const (
	PolicyReadImages     = "gallery_api:ReadImages"
	PolicyCreateImages   = "gallery_api:CreateImages"
	PolicyUpdateImages   = "gallery_api:UpdateImages"
	PolicyDeleteImages   = "gallery_api:DeleteImages"
	PolicyUploadImages   = "gallery_api:UploadImages"
	PolicyDownloadImages = "gallery_api:DownloadImages"
)

type ApiV1 interface {
	RegisterAuthorizationPolicies(middleware *authorization.Middleware)
	RegisterRoutes(r *router.Router)
}

type apiV1 struct {
	service gallery.Service
}

func NewApiV1(service gallery.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterAuthorizationPolicies(middleware *authorization.Middleware) {
	middleware.SetPolicy(authorization.NewPolicy(PolicyReadImages,
		authorization.NewScopeRequirement("gallery.image.read"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyCreateImages,
		authorization.NewScopeRequirement("gallery.image.create"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyUpdateImages,
		authorization.NewScopeRequirement("gallery.image.update"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyDeleteImages,
		authorization.NewScopeRequirement("gallery.image.delete"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyUploadImages,
		authorization.NewScopeRequirement("gallery.image.upload"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyDownloadImages,
		authorization.NewScopeRequirement("gallery.image.download"),
	))
}

func (api *apiV1) RegisterRoutes(r *router.Router) {
	// Images
	r.HandleFunc("/v1/image", api.GetImagesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadImages),
	)
	r.HandleFunc("/v1/image/{id}", api.GetImageByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyReadImages),
	)
	r.HandleFunc("/v1/image", api.CreateImageHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyCreateImages),
	)
	r.HandleFunc("/v1/image/{id}", api.UpdateImageHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyUpdateImages),
	)
	r.HandleFunc("/v1/image/{id}", api.DeleteImageHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyDeleteImages),
	)
	r.HandleFunc("/v1/image/{id}/upload", api.UploadImageFileHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyUploadImages),
	)
	r.HandleFunc("/v1/image/{id}/download", api.DownloadImageFileHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyDownloadImages),
	)
}

func (api *apiV1) handleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	switch err {
	case gallery.ErrImageNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case gallery.ErrImageFileNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case gallery.ErrImageFormatNotSupported:
		rest.WriteError(w, http.StatusUnsupportedMediaType, err.Error())
	case gallery.ErrImageDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case gallery.ErrImageDuplicateSlug:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case core.ErrTranslationNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case core.ErrInvalidId:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	default:
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
	}
	return true
}

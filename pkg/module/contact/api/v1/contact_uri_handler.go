package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
)

func (api *apiV1) GetContactUrisHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	paging := rest.GetPaging(r)
	filter := &model.UriFilter{}
	sort := rest.GetSorting(r)

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().DefaultLanguage(ctx)
	}

	result, count, err := api.service.GetContactUris(ctx, contactId, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := ContactUriListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*ContactUriListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, ContactUriToListItemViewModel(item, language, api.service.GetLanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetContactUriByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	id := router.Param(r, "id")
	result, err := api.service.GetContactUriById(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	response := ContactUriToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateContactUriHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	var model *CreateContactUriV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateContactUri(ctx, contactId, ContactUriFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := ContactUriToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateContactUriHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	id := router.Param(r, "id")

	var model *UpdateContactUriV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateContactUri(ctx, contactId, id, ContactUriFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := ContactUriToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteContactUriHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	id := router.Param(r, "id")

	err := api.service.DeleteContactUri(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

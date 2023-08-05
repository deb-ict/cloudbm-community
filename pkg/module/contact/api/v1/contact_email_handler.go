package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
)

func (api *apiV1) GetContactEmailsHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	paging := rest.GetPaging(r)
	filter := &model.EmailFilter{}
	sort := rest.GetSorting(r)

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetContactEmails(ctx, contactId, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := ContactEmailListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*ContactEmailListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, ContactEmailToListItemViewModel(item, language, api.service.GetLanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetContactEmailByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	id := router.Param(r, "id")
	result, err := api.service.GetContactEmailById(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	response := ContactEmailToViewModel(result, language, api.service.GetLanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateContactEmailHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	var model *CreateContactEmailV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateContactEmail(ctx, contactId, ContactEmailFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	response := ContactEmailToViewModel(result, language, api.service.GetLanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateContactEmailHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	id := router.Param(r, "id")

	var model *UpdateContactEmailV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateContactEmail(ctx, contactId, id, ContactEmailFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	response := ContactEmailToViewModel(result, language, api.service.GetLanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteContactEmailHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	id := router.Param(r, "id")

	err := api.service.DeleteContactEmail(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
)

func (api *apiV1) GetContactPhonesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	paging := rest.GetPaging(r)
	filter := &model.PhoneFilter{}
	sort := rest.GetSorting(r)

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetContactPhones(ctx, contactId, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := ContactPhoneListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*ContactPhoneListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, ContactPhoneToListItemViewModel(item, language, api.service.GetLanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetContactPhoneByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	id := router.Param(r, "id")
	result, err := api.service.GetContactPhoneById(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	response := ContactPhoneToViewModel(result, language, api.service.GetLanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateContactPhoneHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	var model *CreateContactPhoneV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateContactPhone(ctx, contactId, ContactPhoneFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	response := ContactPhoneToViewModel(result, language, api.service.GetLanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateContactPhoneHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	id := router.Param(r, "id")

	var model *UpdateContactPhoneV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateContactPhone(ctx, contactId, id, ContactPhoneFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	response := ContactPhoneToViewModel(result, language, api.service.GetLanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteContactPhoneHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	id := router.Param(r, "id")

	err := api.service.DeleteContactPhone(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

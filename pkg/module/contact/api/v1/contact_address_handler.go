package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
)

func (api *apiV1) GetContactAddressesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	paging := rest.GetPaging(r)
	filter := &model.AddressFilter{}
	sort := rest.GetSorting(r)

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().DefaultLanguage(ctx)
	}

	result, count, err := api.service.GetContactAddresses(ctx, contactId, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := ContactAddressListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*ContactAddressListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, ContactAddressToListItemViewModel(item, language, api.service.GetLanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetContactAddressByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	id := router.Param(r, "id")
	result, err := api.service.GetContactAddressById(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	response := ContactAddressToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateContactAddressHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	var model *CreateContactAddressV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateContactAddress(ctx, contactId, ContactAddressFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := ContactAddressToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateContactAddressHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	id := router.Param(r, "id")

	var model *UpdateContactAddressV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateContactAddress(ctx, contactId, id, ContactAddressFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := ContactAddressToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteContactAddressHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := router.Param(r, "contactId")

	id := router.Param(r, "id")

	err := api.service.DeleteContactAddress(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

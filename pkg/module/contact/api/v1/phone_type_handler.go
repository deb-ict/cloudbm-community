package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
)

func (api *apiV1) GetPhoneTypesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paging := rest.GetPaging(r)
	filter := &model.PhoneTypeFilter{}
	sort := rest.GetSorting(r)

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetPhoneTypes(ctx, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := PhoneTypeListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*PhoneTypeListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, PhoneTypeToListItemViewModel(item, language, api.service.GetLanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetPhoneTypeByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")
	result, err := api.service.GetPhoneTypeById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := PhoneTypeToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreatePhoneTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreatePhoneTypeV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreatePhoneType(ctx, PhoneTypeFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := PhoneTypeToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdatePhoneTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	var model *UpdatePhoneTypeV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdatePhoneType(ctx, id, PhoneTypeFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := PhoneTypeToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeletePhoneTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	err := api.service.DeletePhoneType(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

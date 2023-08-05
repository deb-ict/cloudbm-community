package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
)

func (api *apiV1) GetIndustriesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paging := rest.GetPaging(r)
	filter := &model.IndustryFilter{}
	sort := rest.GetSorting(r)

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetIndustries(ctx, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := IndustryListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*IndustryListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, IndustryToListItemViewModel(item, language, api.service.GetLanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetIndustryByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")
	result, err := api.service.GetIndustryById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := IndustryToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateIndustryHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateIndustryV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateIndustry(ctx, IndustryFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := IndustryToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateIndustryHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	var model *UpdateIndustryV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateIndustry(ctx, id, IndustryFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := IndustryToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteIndustryHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	err := api.service.DeleteIndustry(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

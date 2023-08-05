package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
)

func (api *apiV1) GetCompanyTypesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paging := rest.GetPaging(r)
	filter := &model.CompanyTypeFilter{}
	sort := rest.GetSorting(r)

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().DefaultLanguage(ctx)
	}

	result, count, err := api.service.GetCompanyTypes(ctx, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := CompanyTypeListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*CompanyTypeListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, CompanyTypeToListItemViewModel(item, language, api.service.GetLanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetCompanyTypeByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")
	result, err := api.service.GetCompanyTypeById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := CompanyTypeToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateCompanyTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateCompanyTypeV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateCompanyType(ctx, CompanyTypeFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := CompanyTypeToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateCompanyTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	var model *UpdateCompanyTypeV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateCompanyType(ctx, id, CompanyTypeFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := CompanyTypeToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteCompanyTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	err := api.service.DeleteCompanyType(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
)

func (api *apiV1) GetCompaniesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paging := rest.GetPaging(r)
	filter := &model.CompanyFilter{}
	sort := rest.GetSorting(r)

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetCompanies(ctx, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := CompanyListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*CompanyListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, CompanyToListItemViewModel(item, language, api.service.GetLanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetCompanyByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")
	result, err := api.service.GetCompanyById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := CompanyToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateCompanyHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateCompanyV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateCompany(ctx, CompanyFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := CompanyToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateCompanyHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	var model *UpdateCompanyV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateCompany(ctx, id, CompanyFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := CompanyToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteCompanyHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	err := api.service.DeleteCompany(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

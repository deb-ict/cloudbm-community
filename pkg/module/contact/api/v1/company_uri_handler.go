package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
)

func (api *apiV1) GetCompanyUrisHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	paging := rest.GetPaging(r)
	filter := &model.UriFilter{}
	sort := rest.GetSorting(r)

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetCompanyUris(ctx, companyId, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := CompanyUriListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*CompanyUriListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, CompanyUriToListItemViewModel(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetCompanyUriByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	id := router.Param(r, "id")
	result, err := api.service.GetCompanyUriById(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyUriToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateCompanyUriHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	var model *CreateCompanyUriV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateCompanyUri(ctx, companyId, CompanyUriFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyUriToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateCompanyUriHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	id := router.Param(r, "id")

	var model *UpdateCompanyUriV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateCompanyUri(ctx, companyId, id, CompanyUriFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyUriToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteCompanyUriHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	id := router.Param(r, "id")

	err := api.service.DeleteCompanyUri(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

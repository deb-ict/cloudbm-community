package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
)

func (api *apiV1) GetCompanyEmailsHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	paging := rest.GetPaging(r)
	filter := &model.EmailFilter{}
	sort := rest.GetSorting(r)

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetCompanyEmails(ctx, companyId, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := CompanyEmailListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*CompanyEmailListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, CompanyEmailToListItemViewModel(item, language, api.service.GetLanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetCompanyEmailByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	id := router.Param(r, "id")
	result, err := api.service.GetCompanyEmailById(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	response := CompanyEmailToViewModel(result, language, api.service.GetLanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateCompanyEmailHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	var model *CreateCompanyEmailV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateCompanyEmail(ctx, companyId, CompanyEmailFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	response := CompanyEmailToViewModel(result, language, api.service.GetLanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateCompanyEmailHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	id := router.Param(r, "id")

	var model *UpdateCompanyEmailV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateCompanyEmail(ctx, companyId, id, CompanyEmailFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	response := CompanyEmailToViewModel(result, language, api.service.GetLanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteCompanyEmailHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	id := router.Param(r, "id")

	err := api.service.DeleteCompanyEmail(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

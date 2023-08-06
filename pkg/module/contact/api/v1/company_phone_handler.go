package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
)

func (api *apiV1) GetCompanyPhonesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	paging := rest.GetPaging(r)
	filter := &model.PhoneFilter{}
	sort := rest.GetSorting(r)

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetCompanyPhones(ctx, companyId, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := CompanyPhoneListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*CompanyPhoneListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, CompanyPhoneToListItemViewModel(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetCompanyPhoneByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	id := router.Param(r, "id")
	result, err := api.service.GetCompanyPhoneById(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyPhoneToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateCompanyPhoneHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	var model *CreateCompanyPhoneV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateCompanyPhone(ctx, companyId, CompanyPhoneFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyPhoneToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateCompanyPhoneHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	id := router.Param(r, "id")

	var model *UpdateCompanyPhoneV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateCompanyPhone(ctx, companyId, id, CompanyPhoneFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyPhoneToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteCompanyPhoneHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	id := router.Param(r, "id")

	err := api.service.DeleteCompanyPhone(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

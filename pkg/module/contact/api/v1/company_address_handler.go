package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
)

func (api *apiV1) GetCompanyAddressesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	paging := rest.GetPaging(r)
	filter := &model.AddressFilter{}
	sort := rest.GetSorting(r)

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().DefaultLanguage(ctx)
	}

	result, count, err := api.service.GetCompanyAddresses(ctx, companyId, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := CompanyAddressListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*CompanyAddressListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, CompanyAddressToListItemViewModel(item, language, api.service.GetLanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetCompanyAddressByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	id := router.Param(r, "id")
	result, err := api.service.GetCompanyAddressById(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	response := CompanyAddressToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateCompanyAddressHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	var model *CreateCompanyAddressV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateCompanyAddress(ctx, companyId, CompanyAddressFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := CompanyAddressToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateCompanyAddressHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	id := router.Param(r, "id")

	var model *UpdateCompanyAddressV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateCompanyAddress(ctx, companyId, id, CompanyAddressFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := CompanyAddressToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteCompanyAddressHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	id := router.Param(r, "id")

	err := api.service.DeleteCompanyAddress(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/gorilla/mux"
)

type CompanyV1 struct {
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	VatNumber string         `json:"vat_number"`
	Type      *CompanyTypeV1 `json:"type"`
	Industry  *IndustryV1    `json:"industry"`
	IsEnabled bool           `json:"is_enabled"`
	IsSystem  bool           `json:"is_system"`
}

type CompanyListV1 struct {
	rest.PaginatedList
	Items []*CompanyListItemV1 `json:"items"`
}

type CompanyListItemV1 struct {
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	VatNumber string         `json:"vat_number"`
	Type      *CompanyTypeV1 `json:"type"`
	Industry  *IndustryV1    `json:"industry"`
	IsEnabled bool           `json:"is_enabled"`
	IsSystem  bool           `json:"is_system"`
}

type CreateCompanyV1 struct {
	Name       string `json:"name"`
	VatNumber  string `json:"vat_number"`
	TypeId     string `json:"type_id"`
	IndustryId string `json:"industry_id"`
	IsEnabled  bool   `json:"is_enabled"`
}

type UpdateCompanyV1 struct {
	Name       string `json:"name"`
	VatNumber  string `json:"vat_number"`
	TypeId     string `json:"type_id"`
	IndustryId string `json:"industry_id"`
	IsEnabled  bool   `json:"is_enabled"`
}

func (api *apiV1) GetCompaniesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := api.parseCompanyFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	result, count, err := api.service.GetCompanies(ctx, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
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
		response.Items = append(response.Items, CompanyToListItemViewModelV1(item))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetCompanyByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]
	result, err := api.service.GetCompanyById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := CompanyToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateCompanyHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateCompanyV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateCompany(ctx, CompanyFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := CompanyToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateCompanyHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var model *UpdateCompanyV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateCompany(ctx, id, CompanyFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := CompanyToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteCompanyHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	err := api.service.DeleteCompany(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseCompanyFilterV1(r *http.Request) *model.CompanyFilter {
	return &model.CompanyFilter{
		Name: r.URL.Query().Get("name"),
	}
}

func CompanyToViewModelV1(model *model.Company) *CompanyV1 {
	viewModel := &CompanyV1{
		Id:        model.Id,
		Name:      model.Name,
		VatNumber: model.VatNumber,
		IsEnabled: model.IsEnabled,
		IsSystem:  model.IsSystem,
	}
	if model.Type != nil {
		viewModel.Type = CompanyTypeToViewModelV1(model.Type)
	}
	if model.Industry != nil {
		viewModel.Industry = IndustryToViewModelV1(model.Industry)
	}
	return viewModel
}

func CompanyToListItemViewModelV1(model *model.Company) *CompanyListItemV1 {
	viewModel := &CompanyListItemV1{
		Id:        model.Id,
		Name:      model.Name,
		VatNumber: model.VatNumber,
		IsEnabled: model.IsEnabled,
		IsSystem:  model.IsSystem,
	}
	if model.Type != nil {
		viewModel.Type = CompanyTypeToViewModelV1(model.Type)
	}
	if model.Industry != nil {
		viewModel.Industry = IndustryToViewModelV1(model.Industry)
	}
	return viewModel

}

func CompanyFromCreateViewModelV1(viewModel *CreateCompanyV1) *model.Company {
	return &model.Company{
		Name:      viewModel.Name,
		VatNumber: viewModel.VatNumber,
		Type: &model.CompanyType{
			Id: viewModel.TypeId,
		},
		Industry: &model.Industry{
			Id: viewModel.IndustryId,
		},
		Addresses: make([]*model.Address, 0),
		Emails:    make([]*model.Email, 0),
		Phones:    make([]*model.Phone, 0),
		Uris:      make([]*model.Uri, 0),
		IsEnabled: viewModel.IsEnabled,
	}
}

func CompanyFromUpdateViewModelV1(viewModel *UpdateCompanyV1) *model.Company {
	return &model.Company{
		Name:      viewModel.Name,
		VatNumber: viewModel.VatNumber,
		Type: &model.CompanyType{
			Id: viewModel.TypeId,
		},
		Industry: &model.Industry{
			Id: viewModel.IndustryId,
		},
		Addresses: make([]*model.Address, 0),
		Emails:    make([]*model.Email, 0),
		Phones:    make([]*model.Phone, 0),
		Uris:      make([]*model.Uri, 0),
		IsEnabled: viewModel.IsEnabled,
	}
}

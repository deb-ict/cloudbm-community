package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
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

func CompanyToViewModel(model *model.Company) *CompanyV1 {
	viewModel := &CompanyV1{
		Id:        model.Id,
		Name:      model.Name,
		VatNumber: model.VatNumber,
		IsEnabled: model.IsEnabled,
		IsSystem:  model.IsSystem,
	}
	if model.Type != nil {
		viewModel.Type = CompanyTypeToViewModel(model.Type)
	}
	if model.Industry != nil {
		viewModel.Industry = IndustryToViewModel(model.Industry)
	}
	return viewModel
}

func CompanyToListItemViewModel(model *model.Company) *CompanyListItemV1 {
	viewModel := &CompanyListItemV1{
		Id:        model.Id,
		Name:      model.Name,
		VatNumber: model.VatNumber,
		IsEnabled: model.IsEnabled,
		IsSystem:  model.IsSystem,
	}
	if model.Type != nil {
		viewModel.Type = CompanyTypeToViewModel(model.Type)
	}
	if model.Industry != nil {
		viewModel.Industry = IndustryToViewModel(model.Industry)
	}
	return viewModel

}

func CompanyFromCreateViewModel(viewModel *CreateCompanyV1) *model.Company {
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

func CompanyFromUpdateViewModel(viewModel *UpdateCompanyV1) *model.Company {
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

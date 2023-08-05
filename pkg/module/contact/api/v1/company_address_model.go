package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type CompanyAddressV1 struct {
	Id         string               `json:"id"`
	Type       CompanyAddressTypeV1 `json:"type"`
	Street     string               `json:"street"`
	StreetNr   string               `json:"street_nr"`
	Unit       string               `json:"unit"`
	PostalCode string               `json:"postal_code"`
	City       string               `json:"city"`
	State      string               `json:"state"`
	Country    string               `json:"country"`
	IsDefault  bool                 `json:"is_default"`
}

type CompanyAddressTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CompanyAddressListV1 struct {
	rest.PaginatedList
	Items []*CompanyAddressListItemV1 `json:"items"`
}

type CompanyAddressListItemV1 struct {
	Id         string               `json:"id"`
	Type       CompanyAddressTypeV1 `json:"type"`
	Street     string               `json:"street"`
	StreetNr   string               `json:"street_nr"`
	Unit       string               `json:"unit"`
	PostalCode string               `json:"postal_code"`
	City       string               `json:"city"`
	State      string               `json:"state"`
	Country    string               `json:"country"`
	IsDefault  bool                 `json:"is_default"`
}

type CreateCompanyAddressV1 struct {
	TypeId     string `json:"type_id"`
	Street     string `json:"street"`
	StreetNr   string `json:"street_nr"`
	Unit       string `json:"unit"`
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	IsDefault  bool   `json:"is_default"`
}

type UpdateCompanyAddressV1 struct {
	TypeId     string `json:"type_id"`
	Street     string `json:"street"`
	StreetNr   string `json:"street_nr"`
	Unit       string `json:"unit"`
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	IsDefault  bool   `json:"is_default"`
}

func CompanyAddressToViewModel(model *model.Address, language string, defaultLanguage string) *CompanyAddressV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyAddressV1{
		Id: model.Id,
		Type: CompanyAddressTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Street:     model.Street,
		StreetNr:   model.StreetNr,
		Unit:       model.Unit,
		PostalCode: model.PostalCode,
		City:       model.City,
		State:      model.State,
		Country:    model.Country,
		IsDefault:  model.IsDefault,
	}
}

func CompanyAddressToListItemViewModel(model *model.Address, language string, defaultLanguage string) *CompanyAddressListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyAddressListItemV1{
		Id: model.Id,
		Type: CompanyAddressTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Street:     model.Street,
		StreetNr:   model.StreetNr,
		Unit:       model.Unit,
		PostalCode: model.PostalCode,
		City:       model.City,
		State:      model.State,
		Country:    model.Country,
		IsDefault:  model.IsDefault,
	}
}

func CompanyAddressFromCreateViewModel(viewModel *CreateCompanyAddressV1) *model.Address {
	return &model.Address{
		Type: model.AddressType{
			Id: viewModel.TypeId,
		},
		Street:     viewModel.Street,
		StreetNr:   viewModel.StreetNr,
		Unit:       viewModel.Unit,
		PostalCode: viewModel.PostalCode,
		City:       viewModel.City,
		State:      viewModel.State,
		Country:    viewModel.Country,
		IsDefault:  viewModel.IsDefault,
	}
}

func CompanyAddressFromUpdateViewModel(viewModel *UpdateCompanyAddressV1) *model.Address {
	return &model.Address{
		Type: model.AddressType{
			Id: viewModel.TypeId,
		},
		Street:     viewModel.Street,
		StreetNr:   viewModel.StreetNr,
		Unit:       viewModel.Unit,
		PostalCode: viewModel.PostalCode,
		City:       viewModel.City,
		State:      viewModel.State,
		Country:    viewModel.Country,
		IsDefault:  viewModel.IsDefault,
	}
}

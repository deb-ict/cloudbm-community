package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type ContactAddressV1 struct {
	Id         string               `json:"id"`
	Type       ContactAddressTypeV1 `json:"type"`
	Street     string               `json:"street"`
	StreetNr   string               `json:"street_nr"`
	Unit       string               `json:"unit"`
	PostalCode string               `json:"postal_code"`
	City       string               `json:"city"`
	State      string               `json:"state"`
	Country    string               `json:"country"`
	IsDefault  bool                 `json:"is_default"`
}

type ContactAddressTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ContactAddressListV1 struct {
	rest.PaginatedList
	Items []*ContactAddressListItemV1 `json:"items"`
}

type ContactAddressListItemV1 struct {
	Id         string               `json:"id"`
	Type       ContactAddressTypeV1 `json:"type"`
	Street     string               `json:"street"`
	StreetNr   string               `json:"street_nr"`
	Unit       string               `json:"unit"`
	PostalCode string               `json:"postal_code"`
	City       string               `json:"city"`
	State      string               `json:"state"`
	Country    string               `json:"country"`
	IsDefault  bool                 `json:"is_default"`
}

type CreateContactAddressV1 struct {
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

type UpdateContactAddressV1 struct {
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

func ContactAddressToViewModel(model *model.Address, language string, defaultLanguage string) *ContactAddressV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactAddressV1{
		Id: model.Id,
		Type: ContactAddressTypeV1{
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

func ContactAddressToListItemViewModel(model *model.Address, language string, defaultLanguage string) *ContactAddressListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactAddressListItemV1{
		Id: model.Id,
		Type: ContactAddressTypeV1{
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

func ContactAddressFromCreateViewModel(viewModel *CreateContactAddressV1) *model.Address {
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

func ContactAddressFromUpdateViewModel(viewModel *UpdateContactAddressV1) *model.Address {
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

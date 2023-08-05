package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type CompanyPhoneV1 struct {
	Id          string             `json:"id"`
	Type        CompanyPhoneTypeV1 `json:"type"`
	PhoneNumber string             `json:"number"`
	Extension   string             `json:"extension"`
	IsDefault   bool               `json:"is_default"`
}

type CompanyPhoneTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CompanyPhoneListV1 struct {
	rest.PaginatedList
	Items []*CompanyPhoneListItemV1 `json:"items"`
}

type CompanyPhoneListItemV1 struct {
	Id          string             `json:"id"`
	Type        CompanyPhoneTypeV1 `json:"type"`
	PhoneNumber string             `json:"number"`
	Extension   string             `json:"extension"`
	IsDefault   bool               `json:"is_default"`
}

type CreateCompanyPhoneV1 struct {
	TypeId      string `json:"type_id"`
	PhoneNumber string `json:"number"`
	Extension   string `json:"extension"`
	IsDefault   bool   `json:"is_default"`
}

type UpdateCompanyPhoneV1 struct {
	TypeId      string `json:"type_id"`
	PhoneNumber string `json:"number"`
	Extension   string `json:"extension"`
	IsDefault   bool   `json:"is_default"`
}

func CompanyPhoneToViewModel(model *model.Phone, language string, defaultLanguage string) *CompanyPhoneV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyPhoneV1{
		Id: model.Id,
		Type: CompanyPhoneTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		PhoneNumber: model.PhoneNumber,
		Extension:   model.Extension,
		IsDefault:   model.IsDefault,
	}
}

func CompanyPhoneToListItemViewModel(model *model.Phone, language string, defaultLanguage string) *CompanyPhoneListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyPhoneListItemV1{
		Id: model.Id,
		Type: CompanyPhoneTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		PhoneNumber: model.PhoneNumber,
		Extension:   model.Extension,
		IsDefault:   model.IsDefault,
	}
}

func CompanyPhoneFromCreateViewModel(viewModel *CreateCompanyPhoneV1) *model.Phone {
	return &model.Phone{
		Type: model.PhoneType{
			Id: viewModel.TypeId,
		},
		PhoneNumber: viewModel.PhoneNumber,
		Extension:   viewModel.Extension,
		IsDefault:   viewModel.IsDefault,
	}
}

func CompanyPhoneFromUpdateViewModel(viewModel *UpdateCompanyPhoneV1) *model.Phone {
	return &model.Phone{
		Type: model.PhoneType{
			Id: viewModel.TypeId,
		},
		PhoneNumber: viewModel.PhoneNumber,
		Extension:   viewModel.Extension,
		IsDefault:   viewModel.IsDefault,
	}
}

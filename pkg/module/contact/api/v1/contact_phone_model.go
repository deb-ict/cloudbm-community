package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type ContactPhoneV1 struct {
	Id          string             `json:"id"`
	Type        ContactPhoneTypeV1 `json:"type"`
	PhoneNumber string             `json:"number"`
	Extension   string             `json:"extension"`
	IsDefault   bool               `json:"is_default"`
}

type ContactPhoneTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ContactPhoneListV1 struct {
	rest.PaginatedList
	Items []*ContactPhoneListItemV1 `json:"items"`
}

type ContactPhoneListItemV1 struct {
	Id          string             `json:"id"`
	Type        ContactPhoneTypeV1 `json:"type"`
	PhoneNumber string             `json:"number"`
	Extension   string             `json:"extension"`
	IsDefault   bool               `json:"is_default"`
}

type CreateContactPhoneV1 struct {
	TypeId      string `json:"type_id"`
	PhoneNumber string `json:"number"`
	Extension   string `json:"extension"`
	IsDefault   bool   `json:"is_default"`
}

type UpdateContactPhoneV1 struct {
	TypeId      string `json:"type_id"`
	PhoneNumber string `json:"number"`
	Extension   string `json:"extension"`
	IsDefault   bool   `json:"is_default"`
}

func ContactPhoneToViewModel(model *model.Phone, language string, defaultLanguage string) *ContactPhoneV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactPhoneV1{
		Id: model.Id,
		Type: ContactPhoneTypeV1{
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

func ContactPhoneToListItemViewModel(model *model.Phone, language string, defaultLanguage string) *ContactPhoneListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactPhoneListItemV1{
		Id: model.Id,
		Type: ContactPhoneTypeV1{
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

func ContactPhoneFromCreateViewModel(viewModel *CreateContactPhoneV1) *model.Phone {
	return &model.Phone{
		Type: model.PhoneType{
			Id: viewModel.TypeId,
		},
		PhoneNumber: viewModel.PhoneNumber,
		Extension:   viewModel.Extension,
		IsDefault:   viewModel.IsDefault,
	}
}

func ContactPhoneFromUpdateViewModel(viewModel *UpdateContactPhoneV1) *model.Phone {
	return &model.Phone{
		Type: model.PhoneType{
			Id: viewModel.TypeId,
		},
		PhoneNumber: viewModel.PhoneNumber,
		Extension:   viewModel.Extension,
		IsDefault:   viewModel.IsDefault,
	}
}

package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type ContactEmailV1 struct {
	Id        string             `json:"id"`
	Type      ContactEmailTypeV1 `json:"type"`
	Email     string             `json:"email"`
	IsDefault bool               `json:"is_default"`
}

type ContactEmailTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ContactEmailListV1 struct {
	rest.PaginatedList
	Items []*ContactEmailListItemV1 `json:"items"`
}

type ContactEmailListItemV1 struct {
	Id        string             `json:"id"`
	Type      ContactEmailTypeV1 `json:"type"`
	Email     string             `json:"email"`
	IsDefault bool               `json:"is_default"`
}

type CreateContactEmailV1 struct {
	TypeId    string `json:"type_id"`
	Email     string `json:"email"`
	IsDefault bool   `json:"is_default"`
}

type UpdateContactEmailV1 struct {
	TypeId    string `json:"type_id"`
	Email     string `json:"email"`
	IsDefault bool   `json:"is_default"`
}

func ContactEmailToViewModel(model *model.Email, language string, defaultLanguage string) *ContactEmailV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactEmailV1{
		Id: model.Id,
		Type: ContactEmailTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Email:     model.Email,
		IsDefault: model.IsDefault,
	}
}

func ContactEmailToListItemViewModel(model *model.Email, language string, defaultLanguage string) *ContactEmailListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactEmailListItemV1{
		Id: model.Id,
		Type: ContactEmailTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Email:     model.Email,
		IsDefault: model.IsDefault,
	}
}

func ContactEmailFromCreateViewModel(viewModel *CreateContactEmailV1) *model.Email {
	return &model.Email{
		Type: model.EmailType{
			Id: viewModel.TypeId,
		},
		Email:     viewModel.Email,
		IsDefault: viewModel.IsDefault,
	}
}

func ContactEmailFromUpdateViewModel(viewModel *UpdateContactEmailV1) *model.Email {
	return &model.Email{
		Type: model.EmailType{
			Id: viewModel.TypeId,
		},
		Email:     viewModel.Email,
		IsDefault: viewModel.IsDefault,
	}
}

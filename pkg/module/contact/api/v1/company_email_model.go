package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type CompanyEmailV1 struct {
	Id        string             `json:"id"`
	Type      CompanyEmailTypeV1 `json:"type"`
	Email     string             `json:"email"`
	IsDefault bool               `json:"is_default"`
}

type CompanyEmailTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CompanyEmailListV1 struct {
	rest.PaginatedList
	Items []*CompanyEmailListItemV1 `json:"items"`
}

type CompanyEmailListItemV1 struct {
	Id        string             `json:"id"`
	Type      CompanyEmailTypeV1 `json:"type"`
	Email     string             `json:"email"`
	IsDefault bool               `json:"is_default"`
}

type CreateCompanyEmailV1 struct {
	TypeId    string `json:"type_id"`
	Email     string `json:"email"`
	IsDefault bool   `json:"is_default"`
}

type UpdateCompanyEmailV1 struct {
	TypeId    string `json:"type_id"`
	Email     string `json:"email"`
	IsDefault bool   `json:"is_default"`
}

func CompanyEmailToViewModel(model *model.Email, language string, defaultLanguage string) *CompanyEmailV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyEmailV1{
		Id: model.Id,
		Type: CompanyEmailTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Email:     model.Email,
		IsDefault: model.IsDefault,
	}
}

func CompanyEmailToListItemViewModel(model *model.Email, language string, defaultLanguage string) *CompanyEmailListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyEmailListItemV1{
		Id: model.Id,
		Type: CompanyEmailTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Email:     model.Email,
		IsDefault: model.IsDefault,
	}
}

func CompanyEmailFromCreateViewModel(viewModel *CreateCompanyEmailV1) *model.Email {
	return &model.Email{
		Type: model.EmailType{
			Id: viewModel.TypeId,
		},
		Email:     viewModel.Email,
		IsDefault: viewModel.IsDefault,
	}
}

func CompanyEmailFromUpdateViewModel(viewModel *UpdateCompanyEmailV1) *model.Email {
	return &model.Email{
		Type: model.EmailType{
			Id: viewModel.TypeId,
		},
		Email:     viewModel.Email,
		IsDefault: viewModel.IsDefault,
	}
}

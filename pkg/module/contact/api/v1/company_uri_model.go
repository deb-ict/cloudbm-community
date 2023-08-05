package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type CompanyUriV1 struct {
	Id        string           `json:"id"`
	Type      CompanyUriTypeV1 `json:"type"`
	Uri       string           `json:"uri"`
	IsDefault bool             `json:"is_default"`
}

type CompanyUriTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CompanyUriListV1 struct {
	rest.PaginatedList
	Items []*CompanyUriListItemV1 `json:"items"`
}

type CompanyUriListItemV1 struct {
	Id        string           `json:"id"`
	Type      CompanyUriTypeV1 `json:"type"`
	Uri       string           `json:"uri"`
	IsDefault bool             `json:"is_default"`
}

type CreateCompanyUriV1 struct {
	TypeId    string `json:"type_id"`
	Uri       string `json:"uri"`
	IsDefault bool   `json:"is_default"`
}

type UpdateCompanyUriV1 struct {
	TypeId    string `json:"type_id"`
	Uri       string `json:"uri"`
	IsDefault bool   `json:"is_default"`
}

func CompanyUriToViewModel(model *model.Uri, language string, defaultLanguage string) *CompanyUriV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyUriV1{
		Id: model.Id,
		Type: CompanyUriTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Uri:       model.Uri,
		IsDefault: model.IsDefault,
	}
}

func CompanyUriToListItemViewModel(model *model.Uri, language string, defaultLanguage string) *CompanyUriListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyUriListItemV1{
		Id: model.Id,
		Type: CompanyUriTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Uri:       model.Uri,
		IsDefault: model.IsDefault,
	}
}

func CompanyUriFromCreateViewModel(viewModel *CreateCompanyUriV1) *model.Uri {
	return &model.Uri{
		Type: model.UriType{
			Id: viewModel.TypeId,
		},
		Uri:       viewModel.Uri,
		IsDefault: viewModel.IsDefault,
	}
}

func CompanyUriFromUpdateViewModel(viewModel *UpdateCompanyUriV1) *model.Uri {
	return &model.Uri{
		Type: model.UriType{
			Id: viewModel.TypeId,
		},
		Uri:       viewModel.Uri,
		IsDefault: viewModel.IsDefault,
	}
}

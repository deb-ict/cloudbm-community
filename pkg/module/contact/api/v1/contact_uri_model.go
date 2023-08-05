package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type ContactUriV1 struct {
	Id        string           `json:"id"`
	Type      ContactUriTypeV1 `json:"type"`
	Uri       string           `json:"uri"`
	IsDefault bool             `json:"is_default"`
}

type ContactUriTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ContactUriListV1 struct {
	rest.PaginatedList
	Items []*ContactUriListItemV1 `json:"items"`
}

type ContactUriListItemV1 struct {
	Id        string           `json:"id"`
	Type      ContactUriTypeV1 `json:"type"`
	Uri       string           `json:"uri"`
	IsDefault bool             `json:"is_default"`
}

type CreateContactUriV1 struct {
	TypeId    string `json:"type_id"`
	Uri       string `json:"uri"`
	IsDefault bool   `json:"is_default"`
}

type UpdateContactUriV1 struct {
	TypeId    string `json:"type_id"`
	Uri       string `json:"uri"`
	IsDefault bool   `json:"is_default"`
}

func ContactUriToViewModel(model *model.Uri, language string, defaultLanguage string) *ContactUriV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactUriV1{
		Id: model.Id,
		Type: ContactUriTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Uri:       model.Uri,
		IsDefault: model.IsDefault,
	}
}

func ContactUriToListItemViewModel(model *model.Uri, language string, defaultLanguage string) *ContactUriListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactUriListItemV1{
		Id: model.Id,
		Type: ContactUriTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Uri:       model.Uri,
		IsDefault: model.IsDefault,
	}
}

func ContactUriFromCreateViewModel(viewModel *CreateContactUriV1) *model.Uri {
	return &model.Uri{
		Type: model.UriType{
			Id: viewModel.TypeId,
		},
		Uri:       viewModel.Uri,
		IsDefault: viewModel.IsDefault,
	}
}

func ContactUriFromUpdateViewModel(viewModel *UpdateContactUriV1) *model.Uri {
	return &model.Uri{
		Type: model.UriType{
			Id: viewModel.TypeId,
		},
		Uri:       viewModel.Uri,
		IsDefault: viewModel.IsDefault,
	}
}

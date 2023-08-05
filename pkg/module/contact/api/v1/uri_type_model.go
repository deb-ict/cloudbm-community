package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type UriTypeV1 struct {
	Id           string                  `json:"id"`
	Key          string                  `json:"key"`
	Translations []*UriTypeTranslationV1 `json:"translations"`
	IsDefault    bool                    `json:"is_default"`
	IsSystem     bool                    `json:"is_system"`
}

type UriTypeTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UriTypeListV1 struct {
	rest.PaginatedList
	Items []*UriTypeListItemV1 `json:"items"`
}

type UriTypeListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDefault   bool   `json:"is_default"`
	IsSystem    bool   `json:"is_system"`
}

type CreateUriTypeV1 struct {
	Key          string                  `json:"key"`
	Translations []*UriTypeTranslationV1 `json:"translations"`
	IsDefault    bool                    `json:"is_default"`
}

type UpdateUriTypeV1 struct {
	Translations []*UriTypeTranslationV1 `json:"translations"`
	IsDefault    bool                    `json:"is_default"`
}

func UriTypeToViewModel(model *model.UriType) *UriTypeV1 {
	return &UriTypeV1{}
}

func UriTypeToListItemViewModel(model *model.UriType, language string, defaultLanguage string) *UriTypeListItemV1 {
	return &UriTypeListItemV1{}
}

func UriTypeFromCreateViewModel(viewModel *CreateUriTypeV1) *model.UriType {
	return &model.UriType{}
}

func UriTypeFromUpdateViewModel(viewModel *UpdateUriTypeV1) *model.UriType {
	return &model.UriType{}
}

func UriTypeTranslationToViewModel(model *model.UriTypeTranslation) *UriTypeTranslationV1 {
	return &UriTypeTranslationV1{}
}

func UriTypeTranslationFromViewModel(viewModel *UriTypeTranslationV1) *model.UriTypeTranslation {
	return &model.UriTypeTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}

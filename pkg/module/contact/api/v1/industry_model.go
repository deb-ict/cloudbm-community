package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type IndustryV1 struct {
	Id           string                   `json:"id"`
	Key          string                   `json:"key"`
	Translations []*IndustryTranslationV1 `json:"translations"`
	IsSystem     bool                     `json:"is_system"`
}

type IndustryTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type IndustryListV1 struct {
	rest.PaginatedList
	Items []*IndustryListItemV1 `json:"items"`
}

type IndustryListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsSystem    bool   `json:"is_system"`
}

type CreateIndustryV1 struct {
	Key          string                   `json:"key"`
	Translations []*IndustryTranslationV1 `json:"translations"`
}

type UpdateIndustryV1 struct {
	Translations []*IndustryTranslationV1 `json:"translations"`
}

func IndustryToViewModel(model *model.Industry) *IndustryV1 {
	return &IndustryV1{}
}

func IndustryToListItemViewModel(model *model.Industry, language string, defaultLanguage string) *IndustryListItemV1 {
	return &IndustryListItemV1{}
}

func IndustryFromCreateViewModel(viewModel *CreateIndustryV1) *model.Industry {
	return &model.Industry{}
}

func IndustryFromUpdateViewModel(viewModel *UpdateIndustryV1) *model.Industry {
	return &model.Industry{}
}

func IndustryTranslationToViewModel(model *model.IndustryTranslation) *IndustryTranslationV1 {
	return &IndustryTranslationV1{}
}

func IndustryTranslationFromViewModel(viewModel *IndustryTranslationV1) *model.IndustryTranslation {
	return &model.IndustryTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}

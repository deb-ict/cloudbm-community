package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type CompanyTypeV1 struct {
	Id           string                      `json:"id"`
	Key          string                      `json:"key"`
	Translations []*CompanyTypeTranslationV1 `json:"translations"`
	IsSystem     bool                        `json:"is_system"`
}

type CompanyTypeTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CompanyTypeListV1 struct {
	rest.PaginatedList
	Items []*CompanyTypeListItemV1 `json:"items"`
}

type CompanyTypeListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsSystem    bool   `json:"is_system"`
}

type CreateCompanyTypeV1 struct {
	Key          string                      `json:"key"`
	Translations []*CompanyTypeTranslationV1 `json:"translations"`
}

type UpdateCompanyTypeV1 struct {
	Translations []*CompanyTypeTranslationV1 `json:"translations"`
}

func CompanyTypeToViewModel(model *model.CompanyType) *CompanyTypeV1 {
	return &CompanyTypeV1{}
}

func CompanyTypeToListItemViewModel(model *model.CompanyType, language string, defaultLanguage string) *CompanyTypeListItemV1 {
	return &CompanyTypeListItemV1{}
}

func CompanyTypeFromCreateViewModel(viewModel *CreateCompanyTypeV1) *model.CompanyType {
	return &model.CompanyType{}
}

func CompanyTypeFromUpdateViewModel(viewModel *UpdateCompanyTypeV1) *model.CompanyType {
	return &model.CompanyType{}
}

func CompanyTypeTranslationToViewModel(model *model.CompanyTypeTranslation) *CompanyTypeTranslationV1 {
	return &CompanyTypeTranslationV1{}
}

func CompanyTypeTranslationFromViewModel(viewModel *CompanyTypeTranslationV1) *model.CompanyTypeTranslation {
	return &model.CompanyTypeTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}

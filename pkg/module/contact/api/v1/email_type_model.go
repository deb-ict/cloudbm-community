package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type EmailTypeV1 struct {
	Id           string                    `json:"id"`
	Key          string                    `json:"key"`
	Translations []*EmailTypeTranslationV1 `json:"translations"`
	IsDefault    bool                      `json:"is_default"`
	IsSystem     bool                      `json:"is_system"`
}

type EmailTypeTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EmailTypeListV1 struct {
	rest.PaginatedList
	Items []*EmailTypeListItemV1 `json:"items"`
}

type EmailTypeListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDefault   bool   `json:"is_default"`
	IsSystem    bool   `json:"is_system"`
}

type CreateEmailTypeV1 struct {
	Key          string                    `json:"key"`
	Translations []*EmailTypeTranslationV1 `json:"translations"`
	IsDefault    bool                      `json:"is_default"`
}

type UpdateEmailTypeV1 struct {
	Translations []*EmailTypeTranslationV1 `json:"translations"`
	IsDefault    bool                      `json:"is_default"`
}

func EmailTypeToViewModel(model *model.EmailType) *EmailTypeV1 {
	return &EmailTypeV1{}
}

func EmailTypeToListItemViewModel(model *model.EmailType, language string, defaultLanguage string) *EmailTypeListItemV1 {
	return &EmailTypeListItemV1{}
}

func EmailTypeFromCreateViewModel(viewModel *CreateEmailTypeV1) *model.EmailType {
	return &model.EmailType{}
}

func EmailTypeFromUpdateViewModel(viewModel *UpdateEmailTypeV1) *model.EmailType {
	return &model.EmailType{}
}

func EmailTypeTranslationToViewModel(model *model.EmailTypeTranslation) *EmailTypeTranslationV1 {
	return &EmailTypeTranslationV1{}
}

func EmailTypeTranslationFromViewModel(viewModel *EmailTypeTranslationV1) *model.EmailTypeTranslation {
	return &model.EmailTypeTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}

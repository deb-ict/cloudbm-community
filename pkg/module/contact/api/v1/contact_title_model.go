package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type ContactTitleV1 struct {
	Id           string                       `json:"id"`
	Key          string                       `json:"key"`
	Translations []*ContactTitleTranslationV1 `json:"translations"`
	IsSystem     bool                         `json:"is_system"`
}

type ContactTitleTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ContactTitleListV1 struct {
	rest.PaginatedList
	Items []*ContactTitleListItemV1 `json:"items"`
}

type ContactTitleListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsSystem    bool   `json:"is_system"`
}

type CreateContactTitleV1 struct {
	Key          string                       `json:"key"`
	Translations []*ContactTitleTranslationV1 `json:"translations"`
}

type UpdateContactTitleV1 struct {
	Translations []*ContactTitleTranslationV1 `json:"translations"`
}

func ContactTitleToViewModel(model *model.ContactTitle) *ContactTitleV1 {
	return &ContactTitleV1{}
}

func ContactTitleToListItemViewModel(model *model.ContactTitle, language string, defaultLanguage string) *ContactTitleListItemV1 {
	return &ContactTitleListItemV1{}
}

func ContactTitleFromCreateViewModel(viewModel *CreateContactTitleV1) *model.ContactTitle {
	return &model.ContactTitle{}
}

func ContactTitleFromUpdateViewModel(viewModel *UpdateContactTitleV1) *model.ContactTitle {
	return &model.ContactTitle{}
}

func ContactTitleTranslationToViewModel(model *model.ContactTitleTranslation) *ContactTitleTranslationV1 {
	return &ContactTitleTranslationV1{}
}

func ContactTitleTranslationFromViewModel(viewModel *ContactTitleTranslationV1) *model.ContactTitleTranslation {
	return &model.ContactTitleTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}

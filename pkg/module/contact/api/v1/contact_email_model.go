package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type ContactEmailV1 struct {
}

type ContactEmailListV1 struct {
	rest.PaginatedList
	Items []*ContactEmailListItemV1 `json:"items"`
}

type ContactEmailListItemV1 struct {
}

type CreateContactEmailV1 struct {
}

type UpdateContactEmailV1 struct {
}

func ContactEmailToViewModel(model *model.Email) *ContactEmailV1 {
	return &ContactEmailV1{}
}

func ContactEmailToListItemViewModel(model *model.Email, language string, defaultLanguage string) *ContactEmailListItemV1 {
	return &ContactEmailListItemV1{}
}

func ContactEmailFromCreateViewModel(viewModel *CreateContactEmailV1) *model.Email {
	return &model.Email{}
}

func ContactEmailFromUpdateViewModel(viewModel *UpdateContactEmailV1) *model.Email {
	return &model.Email{}
}

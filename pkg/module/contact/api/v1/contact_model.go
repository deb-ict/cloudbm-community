package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type ContactV1 struct {
}

type ContactListV1 struct {
	rest.PaginatedList
	Items []*ContactListItemV1 `json:"items"`
}

type ContactListItemV1 struct {
}

type CreateContactV1 struct {
}

type UpdateContactV1 struct {
}

func ContactToViewModel(model *model.Contact) *ContactV1 {
	return &ContactV1{}
}

func ContactToListItemViewModel(model *model.Contact, language string, defaultLanguage string) *ContactListItemV1 {
	return &ContactListItemV1{}
}

func ContactFromCreateViewModel(viewModel *CreateContactV1) *model.Contact {
	return &model.Contact{}
}

func ContactFromUpdateViewModel(viewModel *UpdateContactV1) *model.Contact {
	return &model.Contact{}
}

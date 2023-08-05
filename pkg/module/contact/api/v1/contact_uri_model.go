package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type ContactUriV1 struct {
}

type ContactUriListV1 struct {
	rest.PaginatedList
	Items []*ContactUriListItemV1 `json:"items"`
}

type ContactUriListItemV1 struct {
}

type CreateContactUriV1 struct {
}

type UpdateContactUriV1 struct {
}

func ContactUriToViewModel(model *model.Uri) *ContactUriV1 {
	return &ContactUriV1{}
}

func ContactUriToListItemViewModel(model *model.Uri, language string, defaultLanguage string) *ContactUriListItemV1 {
	return &ContactUriListItemV1{}
}

func ContactUriFromCreateViewModel(viewModel *CreateContactUriV1) *model.Uri {
	return &model.Uri{}
}

func ContactUriFromUpdateViewModel(viewModel *UpdateContactUriV1) *model.Uri {
	return &model.Uri{}
}

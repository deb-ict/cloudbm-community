package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type ContactPhoneV1 struct {
}

type ContactPhoneListV1 struct {
	rest.PaginatedList
	Items []*ContactPhoneListItemV1 `json:"items"`
}

type ContactPhoneListItemV1 struct {
}

type CreateContactPhoneV1 struct {
}

type UpdateContactPhoneV1 struct {
}

func ContactPhoneToViewModel(model *model.Phone) *ContactPhoneV1 {
	return &ContactPhoneV1{}
}

func ContactPhoneToListItemViewModel(model *model.Phone, language string, defaultLanguage string) *ContactPhoneListItemV1 {
	return &ContactPhoneListItemV1{}
}

func ContactPhoneFromCreateViewModel(viewModel *CreateContactPhoneV1) *model.Phone {
	return &model.Phone{}
}

func ContactPhoneFromUpdateViewModel(viewModel *UpdateContactPhoneV1) *model.Phone {
	return &model.Phone{}
}

package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type ContactAddressV1 struct {
}

type ContactAddressListV1 struct {
	rest.PaginatedList
	Items []*ContactAddressListItemV1 `json:"items"`
}

type ContactAddressListItemV1 struct {
}

type CreateContactAddressV1 struct {
}

type UpdateContactAddressV1 struct {
}

func ContactAddressToViewModel(model *model.Address) *ContactAddressV1 {
	return &ContactAddressV1{}
}

func ContactAddressToListItemViewModel(model *model.Address, language string, defaultLanguage string) *ContactAddressListItemV1 {
	return &ContactAddressListItemV1{}
}

func ContactAddressFromCreateViewModel(viewModel *CreateContactAddressV1) *model.Address {
	return &model.Address{}
}

func ContactAddressFromUpdateViewModel(viewModel *UpdateContactAddressV1) *model.Address {
	return &model.Address{}
}

package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type CompanyPhoneV1 struct {
}

type CompanyPhoneListV1 struct {
	rest.PaginatedList
	Items []*CompanyPhoneListItemV1 `json:"items"`
}

type CompanyPhoneListItemV1 struct {
}

type CreateCompanyPhoneV1 struct {
}

type UpdateCompanyPhoneV1 struct {
}

func CompanyPhoneToViewModel(model *model.Phone) *CompanyPhoneV1 {
	return &CompanyPhoneV1{}
}

func CompanyPhoneToListItemViewModel(model *model.Phone, language string, defaultLanguage string) *CompanyPhoneListItemV1 {
	return &CompanyPhoneListItemV1{}
}

func CompanyPhoneFromCreateViewModel(viewModel *CreateCompanyPhoneV1) *model.Phone {
	return &model.Phone{}
}

func CompanyPhoneFromUpdateViewModel(viewModel *UpdateCompanyPhoneV1) *model.Phone {
	return &model.Phone{}
}

package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type CompanyAddressV1 struct {
}

type CompanyAddressListV1 struct {
	rest.PaginatedList
	Items []*CompanyAddressListItemV1 `json:"items"`
}

type CompanyAddressListItemV1 struct {
}

type CreateCompanyAddressV1 struct {
}

type UpdateCompanyAddressV1 struct {
}

func CompanyAddressToViewModel(model *model.Address) *CompanyAddressV1 {
	return &CompanyAddressV1{}
}

func CompanyAddressToListItemViewModel(model *model.Address, language string, defaultLanguage string) *CompanyAddressListItemV1 {
	return &CompanyAddressListItemV1{}
}

func CompanyAddressFromCreateViewModel(viewModel *CreateCompanyAddressV1) *model.Address {
	return &model.Address{}
}

func CompanyAddressFromUpdateViewModel(viewModel *UpdateCompanyAddressV1) *model.Address {
	return &model.Address{}
}

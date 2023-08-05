package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type CompanyEmailV1 struct {
}

type CompanyEmailListV1 struct {
	rest.PaginatedList
	Items []*CompanyEmailListItemV1 `json:"items"`
}

type CompanyEmailListItemV1 struct {
}

type CreateCompanyEmailV1 struct {
}

type UpdateCompanyEmailV1 struct {
}

func CompanyEmailToViewModel(model *model.Email) *CompanyEmailV1 {
	return &CompanyEmailV1{}
}

func CompanyEmailToListItemViewModel(model *model.Email, language string, defaultLanguage string) *CompanyEmailListItemV1 {
	return &CompanyEmailListItemV1{}
}

func CompanyEmailFromCreateViewModel(viewModel *CreateCompanyEmailV1) *model.Email {
	return &model.Email{}
}

func CompanyEmailFromUpdateViewModel(viewModel *UpdateCompanyEmailV1) *model.Email {
	return &model.Email{}
}

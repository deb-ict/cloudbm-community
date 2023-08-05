package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type CompanyUriV1 struct {
}

type CompanyUriListV1 struct {
	rest.PaginatedList
	Items []*CompanyUriListItemV1 `json:"items"`
}

type CompanyUriListItemV1 struct {
}

type CreateCompanyUriV1 struct {
}

type UpdateCompanyUriV1 struct {
}

func CompanyUriToViewModel(model *model.Uri) *CompanyUriV1 {
	return &CompanyUriV1{}
}

func CompanyUriToListItemViewModel(model *model.Uri, language string, defaultLanguage string) *CompanyUriListItemV1 {
	return &CompanyUriListItemV1{}
}

func CompanyUriFromCreateViewModel(viewModel *CreateCompanyUriV1) *model.Uri {
	return &model.Uri{}
}

func CompanyUriFromUpdateViewModel(viewModel *UpdateCompanyUriV1) *model.Uri {
	return &model.Uri{}
}

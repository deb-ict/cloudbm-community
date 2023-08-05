package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type CompanyV1 struct {
}

type CompanyListV1 struct {
	rest.PaginatedList
	Items []*CompanyListItemV1 `json:"items"`
}

type CompanyListItemV1 struct {
}

type CreateCompanyV1 struct {
}

type UpdateCompanyV1 struct {
}

func CompanyToViewModel(model *model.Company) *CompanyV1 {

}

func CompanyToListItemViewModel(model *model.Company, language string, defaultLanguage string) *CompanyListItemV1 {

}

func CompanyFromCreateViewModel(viewModel *CreateCompanyV1) *model.Company {

}

func CompanyFromUpdateViewModel(viewModel *UpdateCompanyV1) *model.Company {

}

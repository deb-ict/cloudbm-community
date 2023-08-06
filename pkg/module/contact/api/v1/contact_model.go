package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type ContactV1 struct {
	Id         string          `json:"id"`
	Title      *ContactTitleV1 `json:"title"`
	FamilyName string          `json:"family_name"`
	MiddleName string          `json:"middle_name"`
	GivenName  string          `json:"given_name"`
	IsEnabled  bool            `json:"is_enabled"`
	IsSystem   bool            `json:"is_system"`
}

type ContactListV1 struct {
	rest.PaginatedList
	Items []*ContactListItemV1 `json:"items"`
}

type ContactListItemV1 struct {
	Id         string          `json:"id"`
	Title      *ContactTitleV1 `json:"title"`
	FamilyName string          `json:"family_name"`
	MiddleName string          `json:"middle_name"`
	GivenName  string          `json:"given_name"`
	IsEnabled  bool            `json:"is_enabled"`
	IsSystem   bool            `json:"is_system"`
}

type CreateContactV1 struct {
	UserId     string `json:"user_id"`
	TitleId    string `json:"title_id"`
	FamilyName string `json:"family_name"`
	MiddleName string `json:"middle_name"`
	GivenName  string `json:"given_name"`
	IsEnabled  bool   `json:"is_enabled"`
}

type UpdateContactV1 struct {
	UserId     string `json:"user_id"`
	TitleId    string `json:"title_id"`
	FamilyName string `json:"family_name"`
	MiddleName string `json:"middle_name"`
	GivenName  string `json:"given_name"`
	IsEnabled  bool   `json:"is_enabled"`
}

func ContactToViewModel(model *model.Contact) *ContactV1 {
	viewModel := &ContactV1{
		Id:         model.Id,
		FamilyName: model.FamilyName,
		MiddleName: model.MiddleName,
		GivenName:  model.GivenName,
		IsEnabled:  model.IsEnabled,
		IsSystem:   model.IsSystem,
	}
	if model.Title != nil {
		viewModel.Title = ContactTitleToViewModel(model.Title)
	}
	return viewModel
}

func ContactToListItemViewModel(model *model.Contact) *ContactListItemV1 {
	viewModel := &ContactListItemV1{
		Id:         model.Id,
		FamilyName: model.FamilyName,
		MiddleName: model.MiddleName,
		GivenName:  model.GivenName,
		IsEnabled:  model.IsEnabled,
		IsSystem:   model.IsSystem,
	}
	if model.Title != nil {
		viewModel.Title = ContactTitleToViewModel(model.Title)
	}
	return viewModel
}

func ContactFromCreateViewModel(viewModel *CreateContactV1) *model.Contact {
	return &model.Contact{
		UserId: viewModel.UserId,
		Title: &model.ContactTitle{
			Id: viewModel.TitleId,
		},
		FamilyName: viewModel.FamilyName,
		MiddleName: viewModel.MiddleName,
		GivenName:  viewModel.GivenName,
		IsEnabled:  viewModel.IsEnabled,
	}
}

func ContactFromUpdateViewModel(viewModel *UpdateContactV1) *model.Contact {
	return &model.Contact{
		UserId: viewModel.UserId,
		Title: &model.ContactTitle{
			Id: viewModel.TitleId,
		},
		FamilyName: viewModel.FamilyName,
		MiddleName: viewModel.MiddleName,
		GivenName:  viewModel.GivenName,
		IsEnabled:  viewModel.IsEnabled,
	}
}

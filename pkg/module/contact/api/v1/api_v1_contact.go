package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
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

func (api *apiV1) GetContactsHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paging := rest.GetPaging(r)
	filter := &model.ContactFilter{}
	sort := rest.GetSorting(r)

	result, count, err := api.service.GetContacts(ctx, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := ContactListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*ContactListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, ContactToListItemViewModel(item))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetContactByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")
	result, err := api.service.GetContactById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := ContactToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateContactHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateContactV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateContact(ctx, ContactFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := ContactToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateContactHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	var model *UpdateContactV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateContact(ctx, id, ContactFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := ContactToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteContactHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	err := api.service.DeleteContact(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
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

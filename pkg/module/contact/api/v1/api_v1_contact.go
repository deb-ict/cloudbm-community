package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/gorilla/mux"
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

	filter := api.parseContactFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	result, count, err := api.service.GetContacts(ctx, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
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
		response.Items = append(response.Items, ContactToListItemViewModelV1(item))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetContactByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]
	result, err := api.service.GetContactById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := ContactToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateContactHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateContactV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateContact(ctx, ContactFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := ContactToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateContactHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var model *UpdateContactV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateContact(ctx, id, ContactFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := ContactToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteContactHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	err := api.service.DeleteContact(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseContactFilterV1(r *http.Request) *model.ContactFilter {
	return &model.ContactFilter{
		Name: r.URL.Query().Get("name"),
	}
}

func ContactToViewModelV1(model *model.Contact) *ContactV1 {
	viewModel := &ContactV1{
		Id:         model.Id,
		FamilyName: model.FamilyName,
		MiddleName: model.MiddleName,
		GivenName:  model.GivenName,
		IsEnabled:  model.IsEnabled,
		IsSystem:   model.IsSystem,
	}
	if model.Title != nil {
		viewModel.Title = ContactTitleToViewModelV1(model.Title)
	}
	return viewModel
}

func ContactToListItemViewModelV1(model *model.Contact) *ContactListItemV1 {
	viewModel := &ContactListItemV1{
		Id:         model.Id,
		FamilyName: model.FamilyName,
		MiddleName: model.MiddleName,
		GivenName:  model.GivenName,
		IsEnabled:  model.IsEnabled,
		IsSystem:   model.IsSystem,
	}
	if model.Title != nil {
		viewModel.Title = ContactTitleToViewModelV1(model.Title)
	}
	return viewModel
}

func ContactFromCreateViewModelV1(viewModel *CreateContactV1) *model.Contact {
	return &model.Contact{
		Title: &model.ContactTitle{
			Id: viewModel.TitleId,
		},
		FamilyName: viewModel.FamilyName,
		MiddleName: viewModel.MiddleName,
		GivenName:  viewModel.GivenName,
		IsEnabled:  viewModel.IsEnabled,
	}
}

func ContactFromUpdateViewModelV1(viewModel *UpdateContactV1) *model.Contact {
	return &model.Contact{
		Title: &model.ContactTitle{
			Id: viewModel.TitleId,
		},
		FamilyName: viewModel.FamilyName,
		MiddleName: viewModel.MiddleName,
		GivenName:  viewModel.GivenName,
		IsEnabled:  viewModel.IsEnabled,
	}
}

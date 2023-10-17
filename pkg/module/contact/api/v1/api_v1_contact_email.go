package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/gorilla/mux"
)

type ContactEmailV1 struct {
	Id        string             `json:"id"`
	Type      ContactEmailTypeV1 `json:"type"`
	Email     string             `json:"email"`
	IsDefault bool               `json:"is_default"`
}

type ContactEmailTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ContactEmailListV1 struct {
	rest.PaginatedList
	Items []*ContactEmailListItemV1 `json:"items"`
}

type ContactEmailListItemV1 struct {
	Id        string             `json:"id"`
	Type      ContactEmailTypeV1 `json:"type"`
	Email     string             `json:"email"`
	IsDefault bool               `json:"is_default"`
}

type CreateContactEmailV1 struct {
	TypeId    string `json:"type_id"`
	Email     string `json:"email"`
	IsDefault bool   `json:"is_default"`
}

type UpdateContactEmailV1 struct {
	TypeId    string `json:"type_id"`
	Email     string `json:"email"`
	IsDefault bool   `json:"is_default"`
}

func (api *apiV1) GetContactEmailsHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	paging := rest.GetPaging(r)
	filter := &model.EmailFilter{}
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetContactEmails(ctx, contactId, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := ContactEmailListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*ContactEmailListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, ContactEmailToListItemViewModel(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetContactEmailByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	id := mux.Vars(r)["id"]
	result, err := api.service.GetContactEmailById(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ContactEmailToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateContactEmailHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	var model *CreateContactEmailV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateContactEmail(ctx, contactId, ContactEmailFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ContactEmailToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateContactEmailHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	id := mux.Vars(r)["id"]

	var model *UpdateContactEmailV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateContactEmail(ctx, contactId, id, ContactEmailFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ContactEmailToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteContactEmailHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	id := mux.Vars(r)["id"]

	err := api.service.DeleteContactEmail(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func ContactEmailToViewModel(model *model.Email, language string, defaultLanguage string) *ContactEmailV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactEmailV1{
		Id: model.Id,
		Type: ContactEmailTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Email:     model.Email,
		IsDefault: model.IsDefault,
	}
}

func ContactEmailToListItemViewModel(model *model.Email, language string, defaultLanguage string) *ContactEmailListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactEmailListItemV1{
		Id: model.Id,
		Type: ContactEmailTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Email:     model.Email,
		IsDefault: model.IsDefault,
	}
}

func ContactEmailFromCreateViewModel(viewModel *CreateContactEmailV1) *model.Email {
	return &model.Email{
		Type: model.EmailType{
			Id: viewModel.TypeId,
		},
		Email:     viewModel.Email,
		IsDefault: viewModel.IsDefault,
	}
}

func ContactEmailFromUpdateViewModel(viewModel *UpdateContactEmailV1) *model.Email {
	return &model.Email{
		Type: model.EmailType{
			Id: viewModel.TypeId,
		},
		Email:     viewModel.Email,
		IsDefault: viewModel.IsDefault,
	}
}

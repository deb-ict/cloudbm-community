package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/gorilla/mux"
)

type ContactPhoneV1 struct {
	Id          string             `json:"id"`
	Type        ContactPhoneTypeV1 `json:"type"`
	PhoneNumber string             `json:"number"`
	Extension   string             `json:"extension"`
	IsDefault   bool               `json:"is_default"`
}

type ContactPhoneTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ContactPhoneListV1 struct {
	rest.PaginatedList
	Items []*ContactPhoneListItemV1 `json:"items"`
}

type ContactPhoneListItemV1 struct {
	Id          string             `json:"id"`
	Type        ContactPhoneTypeV1 `json:"type"`
	PhoneNumber string             `json:"number"`
	Extension   string             `json:"extension"`
	IsDefault   bool               `json:"is_default"`
}

type CreateContactPhoneV1 struct {
	TypeId      string `json:"type_id"`
	PhoneNumber string `json:"number"`
	Extension   string `json:"extension"`
	IsDefault   bool   `json:"is_default"`
}

type UpdateContactPhoneV1 struct {
	TypeId      string `json:"type_id"`
	PhoneNumber string `json:"number"`
	Extension   string `json:"extension"`
	IsDefault   bool   `json:"is_default"`
}

func (api *apiV1) GetContactPhonesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	paging := rest.GetPaging(r)
	filter := &model.PhoneFilter{}
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetContactPhones(ctx, contactId, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := ContactPhoneListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*ContactPhoneListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, ContactPhoneToListItemViewModel(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetContactPhoneByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	id := mux.Vars(r)["id"]
	result, err := api.service.GetContactPhoneById(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ContactPhoneToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateContactPhoneHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	var model *CreateContactPhoneV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateContactPhone(ctx, contactId, ContactPhoneFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ContactPhoneToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateContactPhoneHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	id := mux.Vars(r)["id"]

	var model *UpdateContactPhoneV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateContactPhone(ctx, contactId, id, ContactPhoneFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ContactPhoneToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteContactPhoneHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	id := mux.Vars(r)["id"]

	err := api.service.DeleteContactPhone(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func ContactPhoneToViewModel(model *model.Phone, language string, defaultLanguage string) *ContactPhoneV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactPhoneV1{
		Id: model.Id,
		Type: ContactPhoneTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		PhoneNumber: model.PhoneNumber,
		Extension:   model.Extension,
		IsDefault:   model.IsDefault,
	}
}

func ContactPhoneToListItemViewModel(model *model.Phone, language string, defaultLanguage string) *ContactPhoneListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactPhoneListItemV1{
		Id: model.Id,
		Type: ContactPhoneTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		PhoneNumber: model.PhoneNumber,
		Extension:   model.Extension,
		IsDefault:   model.IsDefault,
	}
}

func ContactPhoneFromCreateViewModel(viewModel *CreateContactPhoneV1) *model.Phone {
	return &model.Phone{
		Type: model.PhoneType{
			Id: viewModel.TypeId,
		},
		PhoneNumber: viewModel.PhoneNumber,
		Extension:   viewModel.Extension,
		IsDefault:   viewModel.IsDefault,
	}
}

func ContactPhoneFromUpdateViewModel(viewModel *UpdateContactPhoneV1) *model.Phone {
	return &model.Phone{
		Type: model.PhoneType{
			Id: viewModel.TypeId,
		},
		PhoneNumber: viewModel.PhoneNumber,
		Extension:   viewModel.Extension,
		IsDefault:   viewModel.IsDefault,
	}
}

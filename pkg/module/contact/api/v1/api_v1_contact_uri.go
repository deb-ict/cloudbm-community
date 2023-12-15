package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/gorilla/mux"
)

type ContactUriV1 struct {
	Id        string           `json:"id"`
	Type      ContactUriTypeV1 `json:"type"`
	Uri       string           `json:"uri"`
	IsDefault bool             `json:"is_default"`
}

type ContactUriTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ContactUriListV1 struct {
	rest.PaginatedList
	Items []*ContactUriListItemV1 `json:"items"`
}

type ContactUriListItemV1 struct {
	Id        string           `json:"id"`
	Type      ContactUriTypeV1 `json:"type"`
	Uri       string           `json:"uri"`
	IsDefault bool             `json:"is_default"`
}

type CreateContactUriV1 struct {
	TypeId    string `json:"type_id"`
	Uri       string `json:"uri"`
	IsDefault bool   `json:"is_default"`
}

type UpdateContactUriV1 struct {
	TypeId    string `json:"type_id"`
	Uri       string `json:"uri"`
	IsDefault bool   `json:"is_default"`
}

func (api *apiV1) GetContactUrisHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	paging := rest.GetPaging(r)
	filter := &model.UriFilter{}
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetContactUris(ctx, contactId, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := ContactUriListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*ContactUriListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, ContactUriToListItemViewModel(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetContactUriByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	id := mux.Vars(r)["id"]
	result, err := api.service.GetContactUriById(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ContactUriToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateContactUriHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	var model *CreateContactUriV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateContactUri(ctx, contactId, ContactUriFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ContactUriToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateContactUriHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	id := mux.Vars(r)["id"]

	var model *UpdateContactUriV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateContactUri(ctx, contactId, id, ContactUriFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ContactUriToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteContactUriHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	id := mux.Vars(r)["id"]

	err := api.service.DeleteContactUri(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func ContactUriToViewModel(model *model.Uri, language string, defaultLanguage string) *ContactUriV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactUriV1{
		Id: model.Id,
		Type: ContactUriTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Uri:       model.Uri,
		IsDefault: model.IsDefault,
	}
}

func ContactUriToListItemViewModel(model *model.Uri, language string, defaultLanguage string) *ContactUriListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactUriListItemV1{
		Id: model.Id,
		Type: ContactUriTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Uri:       model.Uri,
		IsDefault: model.IsDefault,
	}
}

func ContactUriFromCreateViewModel(viewModel *CreateContactUriV1) *model.Uri {
	return &model.Uri{
		Type: model.UriType{
			Id: viewModel.TypeId,
		},
		Uri:       viewModel.Uri,
		IsDefault: viewModel.IsDefault,
	}
}

func ContactUriFromUpdateViewModel(viewModel *UpdateContactUriV1) *model.Uri {
	return &model.Uri{
		Type: model.UriType{
			Id: viewModel.TypeId,
		},
		Uri:       viewModel.Uri,
		IsDefault: viewModel.IsDefault,
	}
}

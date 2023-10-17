package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/gorilla/mux"
)

type AddressTypeV1 struct {
	Id           string                      `json:"id"`
	Key          string                      `json:"key"`
	Translations []*AddressTypeTranslationV1 `json:"translations"`
	IsDefault    bool                        `json:"is_default"`
	IsSystem     bool                        `json:"is_system"`
}

type AddressTypeTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AddressTypeListV1 struct {
	rest.PaginatedList
	Items []*AddressTypeListItemV1 `json:"items"`
}

type AddressTypeListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDefault   bool   `json:"is_default"`
	IsSystem    bool   `json:"is_system"`
}

type CreateAddressTypeV1 struct {
	Key          string                      `json:"key"`
	Translations []*AddressTypeTranslationV1 `json:"translations"`
	IsDefault    bool                        `json:"is_default"`
}

type UpdateAddressTypeV1 struct {
	Translations []*AddressTypeTranslationV1 `json:"translations"`
	IsDefault    bool                        `json:"is_default"`
}

func (api *apiV1) GetAddressTypesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paging := rest.GetPaging(r)
	filter := &model.AddressTypeFilter{}
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetAddressTypes(ctx, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := AddressTypeListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*AddressTypeListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, AddressTypeToListItemViewModel(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetAddressTypeByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]
	result, err := api.service.GetAddressTypeById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := AddressTypeToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateAddressTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateAddressTypeV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateAddressType(ctx, AddressTypeFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := AddressTypeToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateAddressTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var model *UpdateAddressTypeV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateAddressType(ctx, id, AddressTypeFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := AddressTypeToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteAddressTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	err := api.service.DeleteAddressType(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func AddressTypeToViewModel(model *model.AddressType) *AddressTypeV1 {
	viewModel := &AddressTypeV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*AddressTypeTranslationV1, 0),
		IsDefault:    model.IsDefault,
		IsSystem:     model.IsSystem,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, AddressTypeTranslationToViewModel(translation))
	}
	return viewModel
}

func AddressTypeToListItemViewModel(model *model.AddressType, language string, defaultLanguage string) *AddressTypeListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &AddressTypeListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		IsDefault:   model.IsDefault,
		IsSystem:    model.IsSystem,
	}
}

func AddressTypeFromCreateViewModel(viewModel *CreateAddressTypeV1) *model.AddressType {
	model := &model.AddressType{
		Key:          viewModel.Key,
		Translations: make([]*model.AddressTypeTranslation, 0),
		IsDefault:    viewModel.IsDefault,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, AddressTypeTranslationFromViewModel(translation))
	}
	return model
}

func AddressTypeFromUpdateViewModel(viewModel *UpdateAddressTypeV1) *model.AddressType {
	model := &model.AddressType{
		Translations: make([]*model.AddressTypeTranslation, 0),
		IsDefault:    viewModel.IsDefault,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, AddressTypeTranslationFromViewModel(translation))
	}
	return model
}

func AddressTypeTranslationToViewModel(model *model.AddressTypeTranslation) *AddressTypeTranslationV1 {
	return &AddressTypeTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func AddressTypeTranslationFromViewModel(viewModel *AddressTypeTranslationV1) *model.AddressTypeTranslation {
	return &model.AddressTypeTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}

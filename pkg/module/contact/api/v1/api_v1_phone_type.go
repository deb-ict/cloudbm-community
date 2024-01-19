package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/gorilla/mux"
)

type PhoneTypeV1 struct {
	Id           string                    `json:"id"`
	Key          string                    `json:"key"`
	Translations []*PhoneTypeTranslationV1 `json:"translations"`
	IsDefault    bool                      `json:"is_default"`
	IsSystem     bool                      `json:"is_system"`
}

type PhoneTypeTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PhoneTypeListV1 struct {
	rest.PaginatedList
	Items []*PhoneTypeListItemV1 `json:"items"`
}

type PhoneTypeListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDefault   bool   `json:"is_default"`
	IsSystem    bool   `json:"is_system"`
}

type CreatePhoneTypeV1 struct {
	Key          string                    `json:"key"`
	Translations []*PhoneTypeTranslationV1 `json:"translations"`
	IsDefault    bool                      `json:"is_default"`
}

type UpdatePhoneTypeV1 struct {
	Translations []*PhoneTypeTranslationV1 `json:"translations"`
	IsDefault    bool                      `json:"is_default"`
}

func (api *apiV1) GetPhoneTypesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := api.parsePhoneTypeFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetPhoneTypes(ctx, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := PhoneTypeListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*PhoneTypeListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, PhoneTypeToListItemViewModelV1(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetPhoneTypeByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]
	result, err := api.service.GetPhoneTypeById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := PhoneTypeToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreatePhoneTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreatePhoneTypeV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreatePhoneType(ctx, PhoneTypeFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := PhoneTypeToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdatePhoneTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var model *UpdatePhoneTypeV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdatePhoneType(ctx, id, PhoneTypeFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := PhoneTypeToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeletePhoneTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	err := api.service.DeletePhoneType(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parsePhoneTypeFilterV1(r *http.Request) *model.PhoneTypeFilter {
	return &model.PhoneTypeFilter{
		Language: localization.GetHttpRequestLanguage(r, api.service.LanguageProvider()),
		Name:     r.URL.Query().Get("name"),
	}
}

func PhoneTypeToViewModelV1(model *model.PhoneType) *PhoneTypeV1 {
	viewModel := &PhoneTypeV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*PhoneTypeTranslationV1, 0),
		IsDefault:    model.IsDefault,
		IsSystem:     model.IsSystem,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, PhoneTypeTranslationToViewModelV1(translation))
	}
	return viewModel
}

func PhoneTypeToListItemViewModelV1(model *model.PhoneType, language string, defaultLanguage string) *PhoneTypeListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &PhoneTypeListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		IsDefault:   model.IsDefault,
		IsSystem:    model.IsSystem,
	}
}

func PhoneTypeFromCreateViewModelV1(viewModel *CreatePhoneTypeV1) *model.PhoneType {
	model := &model.PhoneType{
		Key:          viewModel.Key,
		Translations: make([]*model.PhoneTypeTranslation, 0),
		IsDefault:    viewModel.IsDefault,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, PhoneTypeTranslationFromViewModelV1(translation))
	}
	return model
}

func PhoneTypeFromUpdateViewModelV1(viewModel *UpdatePhoneTypeV1) *model.PhoneType {
	model := &model.PhoneType{
		Translations: make([]*model.PhoneTypeTranslation, 0),
		IsDefault:    viewModel.IsDefault,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, PhoneTypeTranslationFromViewModelV1(translation))
	}
	return model
}

func PhoneTypeTranslationToViewModelV1(model *model.PhoneTypeTranslation) *PhoneTypeTranslationV1 {
	return &PhoneTypeTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func PhoneTypeTranslationFromViewModelV1(viewModel *PhoneTypeTranslationV1) *model.PhoneTypeTranslation {
	return &model.PhoneTypeTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}

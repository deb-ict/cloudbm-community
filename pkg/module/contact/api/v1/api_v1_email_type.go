package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
)

type EmailTypeV1 struct {
	Id           string                    `json:"id"`
	Key          string                    `json:"key"`
	Translations []*EmailTypeTranslationV1 `json:"translations"`
	IsDefault    bool                      `json:"is_default"`
	IsSystem     bool                      `json:"is_system"`
}

type EmailTypeTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EmailTypeListV1 struct {
	rest.PaginatedList
	Items []*EmailTypeListItemV1 `json:"items"`
}

type EmailTypeListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDefault   bool   `json:"is_default"`
	IsSystem    bool   `json:"is_system"`
}

type CreateEmailTypeV1 struct {
	Key          string                    `json:"key"`
	Translations []*EmailTypeTranslationV1 `json:"translations"`
	IsDefault    bool                      `json:"is_default"`
}

type UpdateEmailTypeV1 struct {
	Translations []*EmailTypeTranslationV1 `json:"translations"`
	IsDefault    bool                      `json:"is_default"`
}

func (api *apiV1) GetEmailTypesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := api.parseEmailTypeFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetEmailTypes(ctx, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := EmailTypeListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*EmailTypeListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, EmailTypeToListItemViewModelV1(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetEmailTypeByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")
	result, err := api.service.GetEmailTypeById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := EmailTypeToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateEmailTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateEmailTypeV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateEmailType(ctx, EmailTypeFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := EmailTypeToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateEmailTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	var model *UpdateEmailTypeV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateEmailType(ctx, id, EmailTypeFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := EmailTypeToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteEmailTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	err := api.service.DeleteEmailType(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseEmailTypeFilterV1(r *http.Request) *model.EmailTypeFilter {
	return &model.EmailTypeFilter{
		Language: localization.GetHttpRequestLanguage(r, api.service.LanguageProvider()),
		Name:     r.URL.Query().Get("name"),
	}
}

func EmailTypeToViewModelV1(model *model.EmailType) *EmailTypeV1 {
	viewModel := &EmailTypeV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*EmailTypeTranslationV1, 0),
		IsDefault:    model.IsDefault,
		IsSystem:     model.IsSystem,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, EmailTypeTranslationToViewModelV1(translation))
	}
	return viewModel
}

func EmailTypeToListItemViewModelV1(model *model.EmailType, language string, defaultLanguage string) *EmailTypeListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &EmailTypeListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		IsDefault:   model.IsDefault,
		IsSystem:    model.IsSystem,
	}
}

func EmailTypeFromCreateViewModelV1(viewModel *CreateEmailTypeV1) *model.EmailType {
	model := &model.EmailType{
		Key:          viewModel.Key,
		Translations: make([]*model.EmailTypeTranslation, 0),
		IsDefault:    viewModel.IsDefault,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, EmailTypeTranslationFromViewModelV1(translation))
	}
	return model
}

func EmailTypeFromUpdateViewModelV1(viewModel *UpdateEmailTypeV1) *model.EmailType {
	model := &model.EmailType{
		Translations: make([]*model.EmailTypeTranslation, 0),
		IsDefault:    viewModel.IsDefault,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, EmailTypeTranslationFromViewModelV1(translation))
	}
	return model
}

func EmailTypeTranslationToViewModelV1(model *model.EmailTypeTranslation) *EmailTypeTranslationV1 {
	return &EmailTypeTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func EmailTypeTranslationFromViewModelV1(viewModel *EmailTypeTranslationV1) *model.EmailTypeTranslation {
	return &model.EmailTypeTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}

package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
)

type ContactTitleV1 struct {
	Id           string                       `json:"id"`
	Key          string                       `json:"key"`
	Translations []*ContactTitleTranslationV1 `json:"translations"`
	IsSystem     bool                         `json:"is_system"`
}

type ContactTitleTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ContactTitleListV1 struct {
	rest.PaginatedList
	Items []*ContactTitleListItemV1 `json:"items"`
}

type ContactTitleListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsSystem    bool   `json:"is_system"`
}

type CreateContactTitleV1 struct {
	Key          string                       `json:"key"`
	Translations []*ContactTitleTranslationV1 `json:"translations"`
}

type UpdateContactTitleV1 struct {
	Translations []*ContactTitleTranslationV1 `json:"translations"`
}

func (api *apiV1) GetContactTitlesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := api.parseContactTitleFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetContactTitles(ctx, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := ContactTitleListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*ContactTitleListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, ContactTitleToListItemViewModelV1(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetContactTitleByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")
	result, err := api.service.GetContactTitleById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := ContactTitleToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateContactTitleHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateContactTitleV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateContactTitle(ctx, ContactTitleFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := ContactTitleToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateContactTitleHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	var model *UpdateContactTitleV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateContactTitle(ctx, id, ContactTitleFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := ContactTitleToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteContactTitleHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	err := api.service.DeleteContactTitle(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseContactTitleFilterV1(r *http.Request) *model.ContactTitleFilter {
	return &model.ContactTitleFilter{
		Language: localization.GetHttpRequestLanguage(r, api.service.LanguageProvider()),
		Name:     r.URL.Query().Get("name"),
	}
}

func ContactTitleToViewModelV1(model *model.ContactTitle) *ContactTitleV1 {
	viewModel := &ContactTitleV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*ContactTitleTranslationV1, 0),
		IsSystem:     model.IsSystem,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, ContactTitleTranslationToViewModelV1(translation))
	}
	return viewModel
}

func ContactTitleToListItemViewModelV1(model *model.ContactTitle, language string, defaultLanguage string) *ContactTitleListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &ContactTitleListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		IsSystem:    model.IsSystem,
	}
}

func ContactTitleFromCreateViewModelV1(viewModel *CreateContactTitleV1) *model.ContactTitle {
	model := &model.ContactTitle{
		Key:          viewModel.Key,
		Translations: make([]*model.ContactTitleTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, ContactTitleTranslationFromViewModelV1(translation))
	}
	return model
}

func ContactTitleFromUpdateViewModelV1(viewModel *UpdateContactTitleV1) *model.ContactTitle {
	model := &model.ContactTitle{
		Translations: make([]*model.ContactTitleTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, ContactTitleTranslationFromViewModelV1(translation))
	}
	return model
}

func ContactTitleTranslationToViewModelV1(model *model.ContactTitleTranslation) *ContactTitleTranslationV1 {
	return &ContactTitleTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func ContactTitleTranslationFromViewModelV1(viewModel *ContactTitleTranslationV1) *model.ContactTitleTranslation {
	return &model.ContactTitleTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}

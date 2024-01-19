package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/gorilla/mux"
)

type IndustryV1 struct {
	Id           string                   `json:"id"`
	Key          string                   `json:"key"`
	Translations []*IndustryTranslationV1 `json:"translations"`
	IsSystem     bool                     `json:"is_system"`
}

type IndustryTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type IndustryListV1 struct {
	rest.PaginatedList
	Items []*IndustryListItemV1 `json:"items"`
}

type IndustryListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsSystem    bool   `json:"is_system"`
}

type CreateIndustryV1 struct {
	Key          string                   `json:"key"`
	Translations []*IndustryTranslationV1 `json:"translations"`
}

type UpdateIndustryV1 struct {
	Translations []*IndustryTranslationV1 `json:"translations"`
}

func (api *apiV1) GetIndustriesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := api.parseIndustryFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetIndustries(ctx, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := IndustryListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*IndustryListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, IndustryToListItemViewModelV1(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetIndustryByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]
	result, err := api.service.GetIndustryById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := IndustryToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateIndustryHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateIndustryV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateIndustry(ctx, IndustryFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := IndustryToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateIndustryHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var model *UpdateIndustryV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateIndustry(ctx, id, IndustryFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := IndustryToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteIndustryHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	err := api.service.DeleteIndustry(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseIndustryFilterV1(r *http.Request) *model.IndustryFilter {
	return &model.IndustryFilter{
		Language: localization.GetHttpRequestLanguage(r, api.service.LanguageProvider()),
		Name:     r.URL.Query().Get("name"),
	}
}

func IndustryToViewModelV1(model *model.Industry) *IndustryV1 {
	viewModel := &IndustryV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*IndustryTranslationV1, 0),
		IsSystem:     model.IsSystem,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, IndustryTranslationToViewModelV1(translation))
	}
	return viewModel
}

func IndustryToListItemViewModelV1(model *model.Industry, language string, defaultLanguage string) *IndustryListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &IndustryListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		IsSystem:    model.IsSystem,
	}
}

func IndustryFromCreateViewModelV1(viewModel *CreateIndustryV1) *model.Industry {
	model := &model.Industry{
		Key:          viewModel.Key,
		Translations: make([]*model.IndustryTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, IndustryTranslationFromViewModelV1(translation))
	}
	return model
}

func IndustryFromUpdateViewModelV1(viewModel *UpdateIndustryV1) *model.Industry {
	model := &model.Industry{
		Translations: make([]*model.IndustryTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, IndustryTranslationFromViewModelV1(translation))
	}
	return model
}

func IndustryTranslationToViewModelV1(model *model.IndustryTranslation) *IndustryTranslationV1 {
	return &IndustryTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func IndustryTranslationFromViewModelV1(viewModel *IndustryTranslationV1) *model.IndustryTranslation {
	return &model.IndustryTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}

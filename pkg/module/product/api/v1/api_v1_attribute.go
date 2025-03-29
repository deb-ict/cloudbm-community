package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
	"github.com/gorilla/mux"
)

type AttributeV1 struct {
	Id           string                    `json:"id"`
	Translations []*AttributeTranslationV1 `json:"translations"`
	IsEnabled    bool                      `json:"is_enabled"`
}

type AttributeTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

type AttributeListV1 struct {
	rest.PaginatedList
	Items []*AttributeListItemV1 `json:"items"`
}

type AttributeListItemV1 struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	IsEnabled   bool   `json:"is_enabled"`
}

type CreateAttributeV1 struct {
	Translations []*AttributeTranslationV1 `json:"translations"`
	IsEnabled    bool                      `json:"is_enabled"`
}

type UpdateAttributeV1 struct {
	Translations []*AttributeTranslationV1 `json:"translations"`
	IsEnabled    bool                      `json:"is_enabled"`
}

func (api *apiV1) GetAttributesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := api.parseAttributeFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	result, count, err := api.service.GetAttributes(ctx, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := AttributeListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*AttributeListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, AttributeToListItemViewModelV1(item, filter.Language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetAttributeByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]
	result, err := api.service.GetAttributeById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := AttributeToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateAttributeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateAttributeV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateAttribute(ctx, AttributeFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := AttributeToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateAttributeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var model *UpdateAttributeV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateAttribute(ctx, id, AttributeFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := AttributeToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteAttributeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	err := api.service.DeleteAttribute(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseAttributeFilterV1(r *http.Request) *model.AttributeFilter {
	filter := &model.AttributeFilter{}

	filter.Language = r.URL.Query().Get("language")
	if filter.Language == "" {
		filter.Language = api.service.LanguageProvider().DefaultLanguage(r.Context())
	}
	filter.Name = r.URL.Query().Get("name")

	return filter
}

func AttributeToViewModelV1(model *model.Attribute) *AttributeV1 {
	viewModel := &AttributeV1{
		Id:           model.Id,
		Translations: make([]*AttributeTranslationV1, len(model.Translations)),
		IsEnabled:    model.IsEnabled,
	}
	for i, translation := range model.Translations {
		viewModel.Translations[i] = AttributeTranslationToViewModelV1(translation)
	}
	return viewModel
}

func AttributeToListItemViewModelV1(model *model.Attribute, language string, defaultLanguage string) *AttributeListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &AttributeListItemV1{
		Id:          model.Id,
		Name:        translation.Name,
		Slug:        translation.Slug,
		Description: translation.Description,
		IsEnabled:   model.IsEnabled,
	}
}

func AttributeFromCreateViewModelV1(viewModel *CreateAttributeV1) *model.Attribute {
	model := &model.Attribute{
		Translations: make([]*model.AttributeTranslation, len(viewModel.Translations)),
		IsEnabled:    viewModel.IsEnabled,
	}
	for i, translation := range viewModel.Translations {
		model.Translations[i] = AttributeTranslationFromViewModelV1(translation)
	}
	return model
}

func AttributeFromUpdateViewModelV1(viewModel *UpdateAttributeV1) *model.Attribute {
	model := &model.Attribute{
		Translations: make([]*model.AttributeTranslation, len(viewModel.Translations)),
		IsEnabled:    viewModel.IsEnabled,
	}
	for i, translation := range viewModel.Translations {
		model.Translations[i] = AttributeTranslationFromViewModelV1(translation)
	}
	return model
}

func AttributeTranslationToViewModelV1(model *model.AttributeTranslation) *AttributeTranslationV1 {
	return &AttributeTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Slug:        model.Slug,
		Description: model.Description,
	}
}

func AttributeTranslationFromViewModelV1(viewModel *AttributeTranslationV1) *model.AttributeTranslation {
	return &model.AttributeTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Slug:        viewModel.Slug,
		Description: viewModel.Description,
	}
}

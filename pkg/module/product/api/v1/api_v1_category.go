package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
	"github.com/gorilla/mux"
)

type CategoryV1 struct {
	Id           string                   `json:"id"`
	ParentId     string                   `json:"parent_id"`
	Translations []*CategoryTranslationV1 `json:"translations"`
	ThumbnailId  string                   `json:"thumbnail_id"`
	IsEnabled    bool                     `json:"is_enabled"`
}

type CategoryTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

type CategoryListV1 struct {
	rest.PaginatedList
	Items []*CategoryListItemV1 `json:"items"`
}

type CategoryListItemV1 struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Summary     string `json:"summary"`
	ThumbnailId string `json:"thumbnail_id"`
	IsEnabled   bool   `json:"is_enabled"`
}

type CreateCategoryV1 struct {
	ParentId     string                   `json:"parent_id"`
	Translations []*CategoryTranslationV1 `json:"translations"`
	ThumbnailId  string                   `json:"thumbnail_id"`
	IsEnabled    bool                     `json:"is_enabled"`
}

type UpdateCategoryV1 struct {
	ParentId     string                   `json:"parent_id"`
	Translations []*CategoryTranslationV1 `json:"translations"`
	ThumbnailId  string                   `json:"thumbnail_id"`
	IsEnabled    bool                     `json:"is_enabled"`
}

func (api *apiV1) GetCategoriesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := api.parseCategoryFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	result, count, err := api.service.GetCategories(ctx, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := CategoryListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*CategoryListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, CategoryToListItemViewModelV1(item, filter.Language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetCategoryByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]
	result, err := api.service.GetCategoryById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := CategoryToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateCategoryHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateCategoryV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateCategory(ctx, CategoryFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := CategoryToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateCategoryHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var model *UpdateCategoryV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateCategory(ctx, id, CategoryFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := CategoryToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteCategoryHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	err := api.service.DeleteCategory(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseCategoryFilterV1(r *http.Request) *model.CategoryFilter {
	filter := &model.CategoryFilter{
		AllLevels: true,
	}

	allLevels := r.URL.Query().Get("allLevels")
	if allLevels == "false" || allLevels == "0" {
		filter.AllLevels = false
	}
	filter.Language = r.URL.Query().Get("language")
	if filter.Language == "" {
		filter.Language = api.service.LanguageProvider().UserLanguage(r.Context())
	}
	filter.Name = r.URL.Query().Get("name")
	filter.ParentId = r.URL.Query().Get("parent")

	return filter
}

func CategoryToViewModelV1(model *model.Category) *CategoryV1 {
	viewModel := &CategoryV1{
		Id:           model.Id,
		ParentId:     model.ParentId,
		Translations: make([]*CategoryTranslationV1, 0),
		ThumbnailId:  model.ThumbnailId,
		IsEnabled:    model.IsEnabled,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, CategoryTranslationToViewModelV1(translation))
	}
	return viewModel
}

func CategoryToListItemViewModelV1(model *model.Category, language string, defaultLanguage string) *CategoryListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &CategoryListItemV1{
		Id:          model.Id,
		Name:        translation.Name,
		Slug:        translation.Slug,
		Summary:     translation.Summary,
		ThumbnailId: model.ThumbnailId,
		IsEnabled:   model.IsEnabled,
	}
}

func CategoryFromCreateViewModelV1(viewModel *CreateCategoryV1) *model.Category {
	model := &model.Category{
		ParentId:     viewModel.ParentId,
		Translations: make([]*model.CategoryTranslation, 0),
		ThumbnailId:  viewModel.ThumbnailId,
		IsEnabled:    viewModel.IsEnabled,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, CategoryTranslationFromViewModelV1(translation))
	}
	return model
}

func CategoryFromUpdateViewModelV1(viewModel *UpdateCategoryV1) *model.Category {
	model := &model.Category{
		ParentId:     viewModel.ParentId,
		Translations: make([]*model.CategoryTranslation, 0),
		ThumbnailId:  viewModel.ThumbnailId,
		IsEnabled:    viewModel.IsEnabled,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, CategoryTranslationFromViewModelV1(translation))
	}
	return model
}

func CategoryTranslationToViewModelV1(model *model.CategoryTranslation) *CategoryTranslationV1 {
	return &CategoryTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Slug:        model.Slug,
		Summary:     model.Summary,
		Description: model.Description,
	}
}

func CategoryTranslationFromViewModelV1(viewModel *CategoryTranslationV1) *model.CategoryTranslation {
	return &model.CategoryTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Slug:        viewModel.Slug,
		Summary:     viewModel.Summary,
		Description: viewModel.Description,
	}
}

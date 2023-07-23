package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
	"github.com/deb-ict/go-router"
)

type CategoryV1 struct {
	Id           string                  `json:"id"`
	ParentId     string                  `json:"parent_id"`
	Translations []CategoryTranslationV1 `json:"translations"`
	ThumbnailId  string                  `json:"thumbnail_id"`
	ThumbnailUri string                  `json:"thumbnail_uri"`
	SortOrder    int64                   `json:"sort_order"`
	IsEnabled    bool                    `json:"is_enabled"`
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
	Items []CategoryListItemV1 `json:"items"`
}

type CategoryListItemV1 struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Summary      string `json:"summary"`
	ThumbnailId  string `json:"thumbnail_id"`
	ThumbnailUri string `json:"thumbnail_uri"`
	SortOrder    int64  `json:"sort_order"`
	IsEnabled    bool   `json:"is_enabled"`
}

type CreateCategoryV1 struct {
	ParentId     string                  `json:"parent_id"`
	Translations []CategoryTranslationV1 `json:"translations"`
	ThumbnailId  string                  `json:"thumbnail_id"`
	ThumbnailUri string                  `json:"thumbnail_uri"`
}

type UpdateCategoryV1 struct {
	ParentId     string                  `json:"parent_id"`
	Translations []CategoryTranslationV1 `json:"translations"`
	ThumbnailId  string                  `json:"thumbnail_id"`
	ThumbnailUri string                  `json:"thumbnail_uri"`
	SortOrder    int64                   `json:"sort_order"`
	IsEnabled    bool                    `json:"is_enabled"`
}

func (api *apiV1) GetCateogiesHandlerV1(w http.ResponseWriter, r *http.Request) {
	paging := rest.GetPaging(r)
	filter := &model.CategoryFilter{}
	sort := rest.GetSorting(r)

	parentId := router.QueryValue(r, "parentId")
	if parentId != "" {
		filter.ParentId = parentId
	}

	result, count, err := api.service.GetCategories(r.Context(), paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := CategoryListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]CategoryListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, api.createCategoryListItemViewModel(item))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetCategoryByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	id := router.Param(r, "id")
	result, err := api.service.GetCategoryById(r.Context(), id)
	if api.handleError(w, err) {
		return
	}

	response := api.createCategoryViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateCategoryHandlerV1(w http.ResponseWriter, r *http.Request) {
	var model CreateCategoryV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateCategory(r.Context(), model.toDomainModel())
	if api.handleError(w, err) {
		return
	}

	response := api.createCategoryViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateCategoryHandlerV1(w http.ResponseWriter, r *http.Request) {
	id := router.Param(r, "id")

	var model UpdateCategoryV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateCategory(r.Context(), id, model.toDomainModel())
	if api.handleError(w, err) {
		return
	}

	response := api.createCategoryViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteCategoryHandlerV1(w http.ResponseWriter, r *http.Request) {
	id := router.Param(r, "id")

	err := api.service.DeleteCategory(r.Context(), id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) registerCategoryRoutes(r *router.Router) {
	r.HandleFunc(
		"/v1/category",
		api.GetCateogiesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized("product.read"),
	)
	r.HandleFunc(
		"/v1/category/{id}",
		api.GetCategoryByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized("product.read"),
	)
	r.HandleFunc(
		"/v1/category",
		api.CreateCategoryHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized("product.create"),
	)
	r.HandleFunc(
		"/v1/category/{id}",
		api.UpdateCategoryHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized("product.update"),
	)
	r.HandleFunc(
		"/v1/category/{id}",
		api.DeleteCategoryHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized("product.delete"),
	)
}

func (api *apiV1) createCategoryViewModel(model *model.Category) CategoryV1 {
	viewModel := CategoryV1{
		Id:           model.Id,
		ParentId:     model.ParentId,
		Translations: make([]CategoryTranslationV1, 0),
		ThumbnailId:  model.ThumbnailId,
		ThumbnailUri: model.ThumbnailUri,
		SortOrder:    model.SortOrder,
		IsEnabled:    model.IsEnabled,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, api.createCategoryTranslationViewModel(translation))
	}

	return viewModel
}

func (api *apiV1) createCategoryTranslationViewModel(model *model.CategoryTranslation) CategoryTranslationV1 {
	return CategoryTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Slug:        model.Slug,
		Summary:     model.Summary,
		Description: model.Description,
	}
}

func (api *apiV1) createCategoryListItemViewModel(model *model.Category) CategoryListItemV1 {
	defaultTranslation := model.GetTranslation("") //TODO: we need to get the current language from cookie?
	return CategoryListItemV1{
		Id:           model.Id,
		Name:         defaultTranslation.Name,
		Slug:         defaultTranslation.Slug,
		Summary:      defaultTranslation.Summary,
		ThumbnailId:  model.ThumbnailId,
		ThumbnailUri: model.ThumbnailUri,
		SortOrder:    model.SortOrder,
		IsEnabled:    model.IsEnabled,
	}
}
func (vm *CreateCategoryV1) toDomainModel() *model.Category {
	model := &model.Category{
		ParentId:     vm.ParentId,
		Translations: make([]*model.CategoryTranslation, 0),
		ThumbnailId:  vm.ThumbnailId,
		ThumbnailUri: vm.ThumbnailUri,
		SortOrder:    -1,
		IsEnabled:    true,
	}
	for _, translation := range vm.Translations {
		model.Translations = append(model.Translations, translation.toDomainModel())
	}
	return model
}

func (vm *UpdateCategoryV1) toDomainModel() *model.Category {
	model := &model.Category{
		ParentId:     vm.ParentId,
		Translations: make([]*model.CategoryTranslation, 0),
		ThumbnailId:  vm.ThumbnailId,
		ThumbnailUri: vm.ThumbnailUri,
		SortOrder:    vm.SortOrder,
		IsEnabled:    vm.IsEnabled,
	}
	for _, translation := range vm.Translations {
		model.Translations = append(model.Translations, translation.toDomainModel())
	}
	return model
}

func (vm *CategoryTranslationV1) toDomainModel() *model.CategoryTranslation {
	return &model.CategoryTranslation{
		Language:    vm.Language,
		Name:        vm.Name,
		Slug:        vm.Slug,
		Summary:     vm.Summary,
		Description: vm.Description,
	}
}

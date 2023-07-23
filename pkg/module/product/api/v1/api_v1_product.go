package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
	"github.com/deb-ict/go-router"
)

type ProductV1 struct {
	Id           string                 `json:"id"`
	CategoryIds  []string               `json:"category_ids"`
	Translations []ProductTranslationV1 `json:"translations"`
	ThumbnailId  string                 `json:"thumbnail_id"`
	ThumbnailUri string                 `json:"thumbnail_uri"`
	Price        string                 `json:"price"`
	IsEnabled    bool                   `json:"is_enabled"`
}

type ProductTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

type ProductListV1 struct {
	rest.PaginatedList
	Items []ProductListItemV1 `json:"items"`
}

type ProductListItemV1 struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Summary      string `json:"summary"`
	ThumbnailId  string `json:"thumbnail_id"`
	ThumbnailUri string `json:"thumbnail_uri"`
	Price        string `json:"price"`
	IsEnabled    bool   `json:"is_enabled"`
}

type CreateProductV1 struct {
	CategoryIds  []string               `json:"category_ids"`
	Translations []ProductTranslationV1 `json:"translations"`
	ThumbnailId  string                 `json:"thumbnail_id"`
	ThumbnailUri string                 `json:"thumbnail_uri"`
	Price        string                 `json:"price"`
}

type UpdateProductV1 struct {
	CategoryIds  []string               `json:"category_ids"`
	Translations []ProductTranslationV1 `json:"translations"`
	ThumbnailId  string                 `json:"thumbnail_id"`
	ThumbnailUri string                 `json:"thumbnail_uri"`
	Price        string                 `json:"price"`
	IsEnabled    bool                   `json:"is_enabled"`
}

func (api *apiV1) GetProductsHandlerV1(w http.ResponseWriter, r *http.Request) {
	paging := rest.GetPaging(r)
	filter := &model.ProductFilter{}
	sort := rest.GetSorting(r)

	categoryId := router.QueryValue(r, "categoryId")
	if categoryId != "" {
		filter.CategoryId = categoryId
	}

	result, count, err := api.service.GetProducts(r.Context(), paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := ProductListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]ProductListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, api.createProductListItemViewModel(item))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetProductByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	id := router.Param(r, "id")
	result, err := api.service.GetProductById(r.Context(), id)
	if api.handleError(w, err) {
		return
	}

	response := api.createProductViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateProductHandlerV1(w http.ResponseWriter, r *http.Request) {
	var model CreateProductV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateProduct(r.Context(), model.toDomainModel())
	if api.handleError(w, err) {
		return
	}

	response := api.createProductViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateProductHandlerV1(w http.ResponseWriter, r *http.Request) {
	id := router.Param(r, "id")

	var model UpdateProductV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateProduct(r.Context(), id, model.toDomainModel())
	if api.handleError(w, err) {
		return
	}

	response := api.createProductViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteProductHandlerV1(w http.ResponseWriter, r *http.Request) {
	id := router.Param(r, "id")

	err := api.service.DeleteProduct(r.Context(), id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) registerProductRoutes(r *router.Router) {
	r.HandleFunc(
		"/v1/product",
		api.GetProductsHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized("product.read"),
	)
	r.HandleFunc(
		"/v1/product/{id}",
		api.GetProductByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized("product.read"),
	)
	r.HandleFunc(
		"/v1/product",
		api.CreateProductHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized("product.create"),
	)
	r.HandleFunc(
		"/v1/product/{id}",
		api.UpdateProductHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized("product.update"),
	)
	r.HandleFunc(
		"/v1/product/{id}",
		api.DeleteProductHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized("product.delete"),
	)
}

func (api *apiV1) createProductViewModel(model *model.Product) ProductV1 {
	viewModel := ProductV1{
		Id:           model.Id,
		CategoryIds:  model.CategoryIds,
		Translations: make([]ProductTranslationV1, 0),
		ThumbnailId:  model.ThumbnailId,
		ThumbnailUri: model.ThumbnailUri,
		IsEnabled:    model.IsEnabled,
		Price:        "TODO: FORMAT PRICE",
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, api.createProductTranslationViewModel(translation))
	}

	return viewModel
}

func (api *apiV1) createProductTranslationViewModel(model *model.ProductTranslation) ProductTranslationV1 {
	return ProductTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Slug:        model.Slug,
		Summary:     model.Summary,
		Description: model.Description,
	}
}

func (api *apiV1) createProductListItemViewModel(model *model.Product) ProductListItemV1 {
	defaultTranslation := model.GetTranslation("") //TODO: we need to get the current language from cookie?
	return ProductListItemV1{
		Id:           model.Id,
		Name:         defaultTranslation.Name,
		Slug:         defaultTranslation.Slug,
		Summary:      defaultTranslation.Summary,
		ThumbnailId:  model.ThumbnailId,
		ThumbnailUri: model.ThumbnailUri,
		Price:        "TODO: FORMAT PRICE",
		IsEnabled:    model.IsEnabled,
	}
}
func (vm *CreateProductV1) toDomainModel() *model.Product {
	model := &model.Product{
		CategoryIds:     vm.CategoryIds,
		Translations:    make([]*model.ProductTranslation, 0),
		ThumbnailId:     vm.ThumbnailId,
		ThumbnailUri:    vm.ThumbnailUri,
		Price:           0, //TODO: Parse price
		PriceMultiplier: 0, //TODO: Parse price
		IsEnabled:       true,
	}
	for _, translation := range vm.Translations {
		model.Translations = append(model.Translations, translation.toDomainModel())
	}
	return model
}

func (vm *UpdateProductV1) toDomainModel() *model.Product {
	model := &model.Product{
		CategoryIds:     vm.CategoryIds,
		Translations:    make([]*model.ProductTranslation, 0),
		ThumbnailId:     vm.ThumbnailId,
		ThumbnailUri:    vm.ThumbnailUri,
		Price:           0, //TODO: Parse price
		PriceMultiplier: 0, //TODO: Parse price
		IsEnabled:       vm.IsEnabled,
	}
	for _, translation := range vm.Translations {
		model.Translations = append(model.Translations, translation.toDomainModel())
	}
	return model
}

func (vm *ProductTranslationV1) toDomainModel() *model.ProductTranslation {
	return &model.ProductTranslation{
		Language:    vm.Language,
		Name:        vm.Name,
		Slug:        vm.Slug,
		Summary:     vm.Summary,
		Description: vm.Description,
	}
}

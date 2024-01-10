package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
	"github.com/gorilla/mux"
)

type ProductV1 struct {
	Id           string                  `json:"id"`
	CategoryIds  []string                `json:"category_ids"`
	Translations []*ProductTranslationV1 `json:"translations"`
	ThumbnailId  string                  `json:"thumbnail_id"`
	Ean          string                  `json:"ean"`
	Sku          string                  `json:"sku"`
	Mpn          string                  `json:"mpn"`
	RegularPrice string                  `json:"regular_price"`
	SalesPrice   string                  `json:"sales_price"`
	IsEnabled    bool                    `json:"is_enabled"`
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
	Items []*ProductListItemV1 `json:"items"`
}

type ProductListItemV1 struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Summary      string `json:"summary"`
	ThumbnailId  string `json:"thumbnail_id"`
	Ean          string `json:"ean"`
	Sku          string `json:"sku"`
	Mpn          string `json:"mpn"`
	RegularPrice string `json:"regular_price"`
	SalesPrice   string `json:"sales_price"`
	IsEnabled    bool   `json:"is_enabled"`
}

type CreateProductV1 struct {
	CategoryIds  []string                `json:"category_ids"`
	Translations []*ProductTranslationV1 `json:"translations"`
	ThumbnailId  string                  `json:"thumbnail_id"`
	Ean          string                  `json:"ean"`
	Sku          string                  `json:"sku"`
	Mpn          string                  `json:"mpn"`
	RegularPrice string                  `json:"regular_price"`
	SalesPrice   string                  `json:"sales_price"`
	IsEnabled    bool                    `json:"is_enabled"`
}

type UpdateProductV1 struct {
	CategoryIds  []string                `json:"category_ids"`
	Translations []*ProductTranslationV1 `json:"translations"`
	ThumbnailId  string                  `json:"thumbnail_id"`
	Ean          string                  `json:"ean"`
	Sku          string                  `json:"sku"`
	Mpn          string                  `json:"mpn"`
	RegularPrice string                  `json:"regular_price"`
	SalesPrice   string                  `json:"sales_price"`
	IsEnabled    bool                    `json:"is_enabled"`
}

func (api *apiV1) GetProductsHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := api.parseProductFilter(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	result, count, err := api.service.GetProducts(ctx, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := ProductListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*ProductListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, ProductToListItemViewModel(item, filter.Language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetProductByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]
	result, err := api.service.GetProductById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ProductToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateProductHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateProductV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateProduct(ctx, ProductFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ProductToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateProductHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var model *UpdateProductV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateProduct(ctx, id, ProductFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ProductToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteProductHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	err := api.service.DeleteProduct(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseProductFilter(r *http.Request) *model.ProductFilter {
	filter := &model.ProductFilter{}

	filter.Language = r.URL.Query().Get("language")
	if filter.Language == "" {
		filter.Language = api.service.LanguageProvider().UserLanguage(r.Context())
	}
	filter.Name = r.URL.Query().Get("name")
	filter.CategoryId = r.URL.Query().Get("category")

	return filter
}

func ProductToViewModel(model *model.Product, language string, defaultLanguage string) *ProductV1 {
	viewModel := &ProductV1{
		Id:           model.Id,
		CategoryIds:  model.CategoryIds,
		Translations: make([]*ProductTranslationV1, 0),
		ThumbnailId:  model.ThumbnailId,
		Ean:          model.Ean,
		Sku:          model.Sku,
		Mpn:          model.Mpn,
		RegularPrice: model.RegularPrice.String(),
		SalesPrice:   model.SalesPrice.String(),
		IsEnabled:    model.IsEnabled,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, ProductTranslationToViewModel(translation))
	}
	return viewModel
}

func ProductToListItemViewModel(model *model.Product, language string, defaultLanguage string) *ProductListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &ProductListItemV1{
		Id:           model.Id,
		Name:         translation.Name,
		Slug:         translation.Slug,
		Summary:      translation.Summary,
		ThumbnailId:  model.ThumbnailId,
		Ean:          model.Ean,
		Sku:          model.Sku,
		Mpn:          model.Mpn,
		RegularPrice: model.RegularPrice.String(),
		SalesPrice:   model.SalesPrice.String(),
		IsEnabled:    model.IsEnabled,
	}
}

func ProductFromCreateViewModel(viewModel *CreateProductV1) *model.Product {
	model := &model.Product{
		CategoryIds:  viewModel.CategoryIds,
		Translations: make([]*model.ProductTranslation, 0),
		ThumbnailId:  viewModel.ThumbnailId,
		Ean:          viewModel.Ean,
		Sku:          viewModel.Sku,
		Mpn:          viewModel.Mpn,
		RegularPrice: core.TryGetDecimalFromString(viewModel.RegularPrice),
		SalesPrice:   core.TryGetDecimalFromString(viewModel.SalesPrice),
		IsEnabled:    viewModel.IsEnabled,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, ProductTranslationFromViewModel(translation))
	}
	return model
}

func ProductFromUpdateViewModel(viewModel *UpdateProductV1) *model.Product {
	model := &model.Product{
		CategoryIds:  viewModel.CategoryIds,
		Translations: make([]*model.ProductTranslation, 0),
		ThumbnailId:  viewModel.ThumbnailId,
		Ean:          viewModel.Ean,
		Sku:          viewModel.Sku,
		Mpn:          viewModel.Mpn,
		RegularPrice: core.TryGetDecimalFromString(viewModel.RegularPrice),
		SalesPrice:   core.TryGetDecimalFromString(viewModel.SalesPrice),
		IsEnabled:    viewModel.IsEnabled,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, ProductTranslationFromViewModel(translation))
	}
	return model
}

func ProductTranslationToViewModel(model *model.ProductTranslation) *ProductTranslationV1 {
	return &ProductTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Slug:        model.Slug,
		Summary:     model.Summary,
		Description: model.Description,
	}
}

func ProductTranslationFromViewModel(viewModel *ProductTranslationV1) *model.ProductTranslation {
	return &model.ProductTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Slug:        viewModel.Slug,
		Summary:     viewModel.Summary,
		Description: viewModel.Description,
	}
}

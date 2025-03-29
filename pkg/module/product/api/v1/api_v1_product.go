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
	Type         string                  `json:"type"`
	TemplateId   string                  `json:"template_id,omitempty"`
	CategoryIds  []string                `json:"category_ids"`
	Translations []*ProductTranslationV1 `json:"translations"`
	Attributes   []*ProductAttributeV1   `json:"attributes,omitempty"`
	ThumbnailId  string                  `json:"thumbnail_id"`
	Gtin         string                  `json:"gtin"`
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

type ProductAttributeV1 struct {
	AttributeId string `json:"attribute_id"`
	ValueId     string `json:"value_id"`
}

type ProductListV1 struct {
	rest.PaginatedList
	Items []*ProductListItemV1 `json:"items"`
}

type ProductListItemV1 struct {
	Id           string `json:"id"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Summary      string `json:"summary"`
	ThumbnailId  string `json:"thumbnail_id"`
	Gtin         string `json:"gtin"`
	Sku          string `json:"sku"`
	Mpn          string `json:"mpn"`
	RegularPrice string `json:"regular_price"`
	SalesPrice   string `json:"sales_price"`
	IsEnabled    bool   `json:"is_enabled"`
}

type CreateProductV1 struct {
	Type         string                  `json:"type"`
	TemplateId   string                  `json:"template_id,omitempty"`
	CategoryIds  []string                `json:"category_ids"`
	Translations []*ProductTranslationV1 `json:"translations"`
	Attributes   []*ProductAttributeV1   `json:"attributes,omitempty"`
	ThumbnailId  string                  `json:"thumbnail_id"`
	Gtin         string                  `json:"gtin"`
	Sku          string                  `json:"sku"`
	Mpn          string                  `json:"mpn"`
	RegularPrice string                  `json:"regular_price"`
	SalesPrice   string                  `json:"sales_price"`
	IsEnabled    bool                    `json:"is_enabled"`
}

type UpdateProductV1 struct {
	Type         string                  `json:"type"`
	TemplateId   string                  `json:"template_id,omitempty"`
	CategoryIds  []string                `json:"category_ids"`
	Translations []*ProductTranslationV1 `json:"translations"`
	Attributes   []*ProductAttributeV1   `json:"attributes,omitempty"`
	ThumbnailId  string                  `json:"thumbnail_id"`
	Gtin         string                  `json:"gtin"`
	Sku          string                  `json:"sku"`
	Mpn          string                  `json:"mpn"`
	RegularPrice string                  `json:"regular_price"`
	SalesPrice   string                  `json:"sales_price"`
	IsEnabled    bool                    `json:"is_enabled"`
}

func (api *apiV1) GetProductsHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := api.parseProductFilterV1(r)
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
		response.Items = append(response.Items, ProductToListItemViewModelV1(item, filter.Language, api.service.LanguageProvider().DefaultLanguage(ctx)))
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

	response := ProductToViewModelV1(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateProductHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateProductV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateProduct(ctx, ProductFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ProductToViewModelV1(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
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

	result, err := api.service.UpdateProduct(ctx, id, ProductFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ProductToViewModelV1(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
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

func (api *apiV1) parseProductFilterV1(r *http.Request) *model.ProductFilter {
	filter := &model.ProductFilter{}

	productType := r.URL.Query().Get("type")
	if productType != "" {
		productTypeValue := model.ParseProductType(productType)
		if productTypeValue.IsValid() {
			filter.Type = &productTypeValue
		}
	}

	filter.Language = r.URL.Query().Get("language")
	if filter.Language == "" {
		filter.Language = api.service.LanguageProvider().UserLanguage(r.Context())
	}
	filter.Name = r.URL.Query().Get("name")
	filter.CategoryId = r.URL.Query().Get("category")

	//TODO: Include templates
	//TODO: Include variants

	return filter
}

func ProductToViewModelV1(model *model.Product, language string, defaultLanguage string) *ProductV1 {
	viewModel := &ProductV1{
		Id:           model.Id,
		Type:         model.Type.String(),
		TemplateId:   model.TemplateId,
		CategoryIds:  make([]string, len(model.CategoryIds)),
		Translations: make([]*ProductTranslationV1, len(model.Translations)),
		Attributes:   make([]*ProductAttributeV1, len(model.Attributes)),
		ThumbnailId:  model.ThumbnailId,
		Gtin:         model.Gtin,
		Sku:          model.Sku,
		Mpn:          model.Mpn,
		RegularPrice: model.RegularPrice.String(),
		SalesPrice:   model.SalesPrice.String(),
		IsEnabled:    model.IsEnabled,
	}
	copy(viewModel.CategoryIds, model.CategoryIds)
	for i, translation := range model.Translations {
		viewModel.Translations[i] = ProductTranslationToViewModelV1(translation)
	}
	for i, attribute := range model.Attributes {
		viewModel.Attributes[i] = ProductAttributeToViewModelV1(attribute)
	}
	return viewModel
}

func ProductToListItemViewModelV1(model *model.Product, language string, defaultLanguage string) *ProductListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &ProductListItemV1{
		Id:           model.Id,
		Type:         model.Type.String(),
		Name:         translation.Name,
		Slug:         translation.Slug,
		Summary:      translation.Summary,
		ThumbnailId:  model.ThumbnailId,
		Gtin:         model.Gtin,
		Sku:          model.Sku,
		Mpn:          model.Mpn,
		RegularPrice: model.RegularPrice.String(),
		SalesPrice:   model.SalesPrice.String(),
		IsEnabled:    model.IsEnabled,
	}
}

func ProductFromCreateViewModelV1(viewModel *CreateProductV1) *model.Product {
	model := &model.Product{
		Type:         model.ParseProductType(viewModel.Type),
		TemplateId:   viewModel.TemplateId,
		CategoryIds:  make([]string, len(viewModel.CategoryIds)),
		Translations: make([]*model.ProductTranslation, len(viewModel.Translations)),
		Attributes:   make([]*model.ProductAttribute, len(viewModel.Attributes)),
		ThumbnailId:  viewModel.ThumbnailId,
		Gtin:         viewModel.Gtin,
		Sku:          viewModel.Sku,
		Mpn:          viewModel.Mpn,
		RegularPrice: core.TryGetDecimalFromString(viewModel.RegularPrice),
		SalesPrice:   core.TryGetDecimalFromString(viewModel.SalesPrice),
		IsEnabled:    viewModel.IsEnabled,
	}
	copy(model.CategoryIds, viewModel.CategoryIds)
	for i, translation := range viewModel.Translations {
		model.Translations[i] = ProductTranslationFromViewModelV1(translation)
	}
	for i, attribute := range viewModel.Attributes {
		model.Attributes[i] = ProductAttributeFromViewModelV1(attribute)
	}
	return model
}

func ProductFromUpdateViewModelV1(viewModel *UpdateProductV1) *model.Product {
	model := &model.Product{
		Type:         model.ParseProductType(viewModel.Type),
		TemplateId:   viewModel.TemplateId,
		CategoryIds:  make([]string, len(viewModel.CategoryIds)),
		Translations: make([]*model.ProductTranslation, len(viewModel.Translations)),
		Attributes:   make([]*model.ProductAttribute, len(viewModel.Attributes)),
		ThumbnailId:  viewModel.ThumbnailId,
		Gtin:         viewModel.Gtin,
		Sku:          viewModel.Sku,
		Mpn:          viewModel.Mpn,
		RegularPrice: core.TryGetDecimalFromString(viewModel.RegularPrice),
		SalesPrice:   core.TryGetDecimalFromString(viewModel.SalesPrice),
		IsEnabled:    viewModel.IsEnabled,
	}
	copy(model.CategoryIds, viewModel.CategoryIds)
	for i, translation := range viewModel.Translations {
		model.Translations[i] = ProductTranslationFromViewModelV1(translation)
	}
	for i, attribute := range viewModel.Attributes {
		model.Attributes[i] = ProductAttributeFromViewModelV1(attribute)
	}
	return model
}

func ProductTranslationToViewModelV1(model *model.ProductTranslation) *ProductTranslationV1 {
	return &ProductTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Slug:        model.Slug,
		Summary:     model.Summary,
		Description: model.Description,
	}
}

func ProductTranslationFromViewModelV1(viewModel *ProductTranslationV1) *model.ProductTranslation {
	return &model.ProductTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Slug:        viewModel.Slug,
		Summary:     viewModel.Summary,
		Description: viewModel.Description,
	}
}

func ProductAttributeToViewModelV1(model *model.ProductAttribute) *ProductAttributeV1 {
	return &ProductAttributeV1{
		AttributeId: model.AttributeId,
		ValueId:     model.ValueId,
	}
}

func ProductAttributeFromViewModelV1(viewModel *ProductAttributeV1) *model.ProductAttribute {
	return &model.ProductAttribute{
		AttributeId: viewModel.AttributeId,
		ValueId:     viewModel.ValueId,
	}
}

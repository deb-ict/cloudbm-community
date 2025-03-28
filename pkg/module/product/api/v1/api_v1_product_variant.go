package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
	"github.com/gorilla/mux"
)

type ProductVariantV1 struct {
	Id           string                         `json:"id"`
	Attributes   []*ProductVariantValueV1       `json:"attributes"`
	Translations []*ProductVariantTranslationV1 `json:"translations"`
	ThumbnailId  string                         `json:"thumbnail_id"`
	Gtin         string                         `json:"gtin"`
	Sku          string                         `json:"sku"`
	Mpn          string                         `json:"mpn"`
	RegularPrice string                         `json:"regular_price"`
	SalesPrice   string                         `json:"sales_price"`
	IsEnabled    bool                           `json:"is_enabled"`
}

type ProductVariantValueV1 struct {
	AttributeId string `json:"attribute_id"`
	ValueId     string `json:"value_id"`
}

type ProductVariantTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

type ProductVariantListV1 struct {
	rest.PaginatedList
	Items []*ProductVariantListItemV1 `json:"items"`
}

type ProductVariantListItemV1 struct {
	Id           string                   `json:"id"`
	Name         string                   `json:"name"`
	Slug         string                   `json:"slug"`
	Summary      string                   `json:"summary"`
	Attributes   []*ProductVariantValueV1 `json:"attributes"`
	ThumbnailId  string                   `json:"thumbnail_id"`
	Gtin         string                   `json:"gtin"`
	Sku          string                   `json:"sku"`
	Mpn          string                   `json:"mpn"`
	RegularPrice string                   `json:"regular_price"`
	SalesPrice   string                   `json:"sales_price"`
	IsEnabled    bool                     `json:"is_enabled"`
}

type CreateProductVariantV1 struct {
	Attributes   []*ProductVariantValueV1       `json:"attributes"`
	Translations []*ProductVariantTranslationV1 `json:"translations"`
	ThumbnailId  string                         `json:"thumbnail_id"`
	Gtin         string                         `json:"gtin"`
	Sku          string                         `json:"sku"`
	Mpn          string                         `json:"mpn"`
	RegularPrice string                         `json:"regular_price"`
	SalesPrice   string                         `json:"sales_price"`
	IsEnabled    bool                           `json:"is_enabled"`
}

type UpdateProductVariantV1 struct {
	Attributes   []*ProductVariantValueV1       `json:"attributes"`
	Translations []*ProductVariantTranslationV1 `json:"translations"`
	ThumbnailId  string                         `json:"thumbnail_id"`
	Gtin         string                         `json:"gtin"`
	Sku          string                         `json:"sku"`
	Mpn          string                         `json:"mpn"`
	RegularPrice string                         `json:"regular_price"`
	SalesPrice   string                         `json:"sales_price"`
	IsEnabled    bool                           `json:"is_enabled"`
}

func (api *apiV1) GetProductVariantsHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	productId := mux.Vars(r)["productId"]
	filter := api.parseProductVariantFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	result, count, err := api.service.GetProductVariants(ctx, productId, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := ProductVariantListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*ProductVariantListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, ProductVariantToListItemViewModelV1(item, filter.Language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetProductVariantByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	productId := mux.Vars(r)["productId"]
	id := mux.Vars(r)["id"]
	result, err := api.service.GetProductVariantById(ctx, productId, id)
	if api.handleError(w, err) {
		return
	}

	response := ProductVariantToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateProductVariantHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	productId := mux.Vars(r)["productId"]

	var model *CreateProductVariantV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateProductVariant(ctx, productId, ProductVariantFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := ProductVariantToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateProductVariantHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	productId := mux.Vars(r)["productId"]
	id := mux.Vars(r)["id"]

	var model *UpdateProductVariantV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateProductVariant(ctx, productId, id, ProductVariantFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := ProductVariantToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteProductVariantHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	productId := mux.Vars(r)["productId"]
	id := mux.Vars(r)["id"]

	err := api.service.DeleteProductVariant(ctx, productId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseProductVariantFilterV1(r *http.Request) *model.ProductVariantFilter {
	filter := &model.ProductVariantFilter{}

	filter.Language = r.URL.Query().Get("language")
	if filter.Language == "" {
		filter.Language = api.service.LanguageProvider().UserLanguage(r.Context())
	}
	filter.Name = r.URL.Query().Get("name")

	return filter
}

func ProductVariantToViewModelV1(model *model.ProductVariant) *ProductVariantV1 {
	viewModel := &ProductVariantV1{
		Id:           model.Id,
		Attributes:   make([]*ProductVariantValueV1, 0),
		Translations: make([]*ProductVariantTranslationV1, 0),
		ThumbnailId:  model.Details.ThumbnailId,
		Gtin:         model.Details.Gtin,
		Sku:          model.Details.Sku,
		Mpn:          model.Details.Mpn,
		RegularPrice: model.Details.RegularPrice.String(),
		SalesPrice:   model.Details.SalesPrice.String(),
		IsEnabled:    model.Details.IsEnabled,
	}
	for _, attribute := range model.Attributes {
		viewModel.Attributes = append(viewModel.Attributes, ProductVariantValueToViewModelV1(attribute))
	}
	for _, translation := range model.Details.Translations {
		viewModel.Translations = append(viewModel.Translations, ProductVariantTranslationToViewModelV1(translation))
	}
	return viewModel
}

func ProductVariantToListItemViewModelV1(model *model.ProductVariant, language string, defaultLanguage string) *ProductVariantListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	viewModel := &ProductVariantListItemV1{
		Id:           model.Id,
		Name:         translation.Name,
		Slug:         translation.Slug,
		Summary:      translation.Summary,
		Attributes:   make([]*ProductVariantValueV1, 0),
		ThumbnailId:  model.Details.ThumbnailId,
		Gtin:         model.Details.Gtin,
		Sku:          model.Details.Sku,
		Mpn:          model.Details.Mpn,
		RegularPrice: model.Details.RegularPrice.String(),
		SalesPrice:   model.Details.SalesPrice.String(),
		IsEnabled:    model.Details.IsEnabled,
	}
	for _, attribute := range model.Attributes {
		viewModel.Attributes = append(viewModel.Attributes, &ProductVariantValueV1{
			AttributeId: attribute.AttributeId,
			ValueId:     attribute.ValueId,
		})
	}
	return viewModel
}

func ProductVariantFromCreateViewModelV1(viewModel *CreateProductVariantV1) *model.ProductVariant {
	model := &model.ProductVariant{
		Attributes: make([]*model.ProductVariantValue, 0),
		Details: &model.ProductDetail{
			Translations: make([]*model.ProductTranslation, 0),
			ThumbnailId:  viewModel.ThumbnailId,
			Gtin:         viewModel.Gtin,
			Sku:          viewModel.Sku,
			Mpn:          viewModel.Mpn,
			RegularPrice: core.TryGetDecimalFromString(viewModel.RegularPrice),
			SalesPrice:   core.TryGetDecimalFromString(viewModel.SalesPrice),
			IsEnabled:    viewModel.IsEnabled,
		},
	}
	for _, attribute := range viewModel.Attributes {
		model.Attributes = append(model.Attributes, ProductVariantValueFromViewModelV1(attribute))
	}
	for _, translation := range viewModel.Translations {
		model.Details.Translations = append(model.Details.Translations, ProductVariantTranslationFromViewModelV1(translation))
	}
	return model
}

func ProductVariantFromUpdateViewModelV1(viewModel *UpdateProductVariantV1) *model.ProductVariant {
	model := &model.ProductVariant{
		Attributes: make([]*model.ProductVariantValue, 0),
		Details: &model.ProductDetail{
			Translations: make([]*model.ProductTranslation, 0),
			ThumbnailId:  viewModel.ThumbnailId,
			Gtin:         viewModel.Gtin,
			Sku:          viewModel.Sku,
			Mpn:          viewModel.Mpn,
			RegularPrice: core.TryGetDecimalFromString(viewModel.RegularPrice),
			SalesPrice:   core.TryGetDecimalFromString(viewModel.SalesPrice),
			IsEnabled:    viewModel.IsEnabled,
		},
	}
	for _, attribute := range viewModel.Attributes {
		model.Attributes = append(model.Attributes, ProductVariantValueFromViewModelV1(attribute))
	}
	for _, translation := range viewModel.Translations {
		model.Details.Translations = append(model.Details.Translations, ProductVariantTranslationFromViewModelV1(translation))
	}
	return model
}

func ProductVariantTranslationToViewModelV1(model *model.ProductTranslation) *ProductVariantTranslationV1 {
	return &ProductVariantTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Slug:        model.Slug,
		Summary:     model.Summary,
		Description: model.Description,
	}
}

func ProductVariantTranslationFromViewModelV1(viewModel *ProductVariantTranslationV1) *model.ProductTranslation {
	return &model.ProductTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Slug:        viewModel.Slug,
		Summary:     viewModel.Summary,
		Description: viewModel.Description,
	}
}

func ProductVariantValueToViewModelV1(model *model.ProductVariantValue) *ProductVariantValueV1 {
	return &ProductVariantValueV1{
		AttributeId: model.AttributeId,
		ValueId:     model.ValueId,
	}
}

func ProductVariantValueFromViewModelV1(viewModel *ProductVariantValueV1) *model.ProductVariantValue {
	return &model.ProductVariantValue{
		AttributeId: viewModel.AttributeId,
		ValueId:     viewModel.AttributeId,
	}
}

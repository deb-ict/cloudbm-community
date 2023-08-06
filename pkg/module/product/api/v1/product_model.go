package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	globalV1 "github.com/deb-ict/cloudbm-community/pkg/module/global/api/v1"
	globalModel "github.com/deb-ict/cloudbm-community/pkg/module/global/model"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
	"github.com/shopspring/decimal"
)

type ProductV1 struct {
	Id           string                          `json:"id"`
	CategoryIds  []string                        `json:"category_ids"`
	Translations []*ProductTranslationV1         `json:"translations"`
	ThumbnailId  string                          `json:"thumbnail_id"`
	ThumbnailUri string                          `json:"thumbnail_uri"`
	Price        decimal.Decimal                 `json:"price"`
	TaxProfile   *globalV1.LocalizedTaxProfileV1 `json:"tax_profile"`
	IsEnabled    bool                            `json:"is_enabled"`
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
	Id           string                          `json:"id"`
	Name         string                          `json:"name"`
	Slug         string                          `json:"slug"`
	Summary      string                          `json:"summary"`
	ThumbnailId  string                          `json:"thumbnail_id"`
	ThumbnailUri string                          `json:"thumbnail_uri"`
	Price        decimal.Decimal                 `json:"price"`
	TaxProfile   *globalV1.LocalizedTaxProfileV1 `json:"tax_profile"`
	IsEnabled    bool                            `json:"is_enabled"`
}

type CreateProductV1 struct {
	CategoryIds  []string                `json:"category_ids"`
	Translations []*ProductTranslationV1 `json:"translations"`
	ThumbnailId  string                  `json:"thumbnail_id"`
	ThumbnailUri string                  `json:"thumbnail_uri"`
	Price        decimal.Decimal         `json:"price"`
	TaxProfileId string                  `json:"tax_profile_id"`
	IsEnabled    bool                    `json:"is_enabled"`
}

type UpdateProductV1 struct {
	CategoryIds  []string                `json:"category_ids"`
	Translations []*ProductTranslationV1 `json:"translations"`
	ThumbnailId  string                  `json:"thumbnail_id"`
	ThumbnailUri string                  `json:"thumbnail_uri"`
	Price        decimal.Decimal         `json:"price"`
	TaxProfileId string                  `json:"tax_profile_id"`
	IsEnabled    bool                    `json:"is_enabled"`
}

func ProductToViewModel(model *model.Product, language string, defaultLanguage string) *ProductV1 {
	viewModel := &ProductV1{
		Id:           model.Id,
		CategoryIds:  model.CategoryIds,
		Translations: make([]*ProductTranslationV1, 0),
		ThumbnailId:  model.ThumbnailId,
		ThumbnailUri: model.ThumbnailUri,
		Price:        model.Price,
		TaxProfile:   globalV1.TaxProfileToLocalizedViewModel(model.TaxProfile, language, defaultLanguage),
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
		ThumbnailUri: model.ThumbnailUri,
		Price:        model.Price,
		TaxProfile:   globalV1.TaxProfileToLocalizedViewModel(model.TaxProfile, language, defaultLanguage),
		IsEnabled:    model.IsEnabled,
	}
}

func ProductFromCreateViewModel(viewModel *CreateProductV1) *model.Product {
	model := &model.Product{
		CategoryIds:  viewModel.CategoryIds,
		Translations: make([]*model.ProductTranslation, 0),
		ThumbnailId:  viewModel.ThumbnailId,
		ThumbnailUri: viewModel.ThumbnailUri,
		Price:        viewModel.Price,
		TaxProfile: &globalModel.TaxProfile{
			Id: viewModel.TaxProfileId,
		},
		IsEnabled: viewModel.IsEnabled,
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
		ThumbnailUri: viewModel.ThumbnailUri,
		Price:        viewModel.Price,
		TaxProfile: &globalModel.TaxProfile{
			Id: viewModel.TaxProfileId,
		},
		IsEnabled: viewModel.IsEnabled,
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

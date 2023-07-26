package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
)

type CategoryV1 struct {
	Id           string                   `json:"id"`
	ParentId     string                   `json:"parent_id"`
	Translations []*CategoryTranslationV1 `json:"translations"`
	ThumbnailId  string                   `json:"thumbnail_id"`
	ThumbnailUri string                   `json:"thumbnail_uri"`
	SortOrder    int64                    `json:"sort_order"`
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
	ParentId     string                   `json:"parent_id"`
	Translations []*CategoryTranslationV1 `json:"translations"`
	ThumbnailId  string                   `json:"thumbnail_id"`
	ThumbnailUri string                   `json:"thumbnail_uri"`
}

type UpdateCategoryV1 struct {
	ParentId     string                   `json:"parent_id"`
	Translations []*CategoryTranslationV1 `json:"translations"`
	ThumbnailId  string                   `json:"thumbnail_id"`
	ThumbnailUri string                   `json:"thumbnail_uri"`
	SortOrder    int64                    `json:"sort_order"`
	IsEnabled    bool                     `json:"is_enabled"`
}

func CategoryToViewModel(model *model.Category) *CategoryV1 {
	viewModel := &CategoryV1{
		Id:           model.Id,
		ParentId:     model.ParentId,
		Translations: make([]*CategoryTranslationV1, 0),
		ThumbnailId:  model.ThumbnailId,
		ThumbnailUri: model.ThumbnailUri,
		SortOrder:    model.SortOrder,
		IsEnabled:    model.IsEnabled,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, CategoryTranslationToViewModel(translation))
	}

	return viewModel
}

func CategoryToListItemViewModel(model *model.Category, language string, defaultLanguage string) *CategoryListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &CategoryListItemV1{
		Id:           model.Id,
		Name:         translation.Name,
		Slug:         translation.Slug,
		Summary:      translation.Summary,
		ThumbnailId:  model.ThumbnailId,
		ThumbnailUri: model.ThumbnailUri,
		SortOrder:    model.SortOrder,
		IsEnabled:    model.IsEnabled,
	}
}

func CategoryFromCreateViewModel(viewModel *CreateCategoryV1) *model.Category {
	model := &model.Category{
		ParentId:     viewModel.ParentId,
		Translations: make([]*model.CategoryTranslation, 0),
		ThumbnailId:  viewModel.ThumbnailId,
		ThumbnailUri: viewModel.ThumbnailUri,
		SortOrder:    -1,
		IsEnabled:    true,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, CategoryTranslationFromViewModel(translation))
	}
	return model
}

func CategoryFromUpdateViewModel(viewModel *UpdateCategoryV1) *model.Category {
	model := &model.Category{
		ParentId:     viewModel.ParentId,
		Translations: make([]*model.CategoryTranslation, 0),
		ThumbnailId:  viewModel.ThumbnailId,
		ThumbnailUri: viewModel.ThumbnailUri,
		SortOrder:    viewModel.SortOrder,
		IsEnabled:    viewModel.IsEnabled,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, CategoryTranslationFromViewModel(translation))
	}
	return model
}

func CategoryTranslationToViewModel(model *model.CategoryTranslation) *CategoryTranslationV1 {
	return &CategoryTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Slug:        model.Slug,
		Summary:     model.Summary,
		Description: model.Description,
	}
}

func CategoryTranslationFromViewModel(viewModel *CategoryTranslationV1) *model.CategoryTranslation {
	return &model.CategoryTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Slug:        viewModel.Slug,
		Summary:     viewModel.Summary,
		Description: viewModel.Description,
	}
}

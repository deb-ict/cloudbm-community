package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
)

func (svc *service) GetCategories(ctx context.Context, offset int64, limit int64, filter *model.CategoryFilter, sort *core.Sort) ([]*model.Category, int64, error) {
	data, count, err := svc.database.Categories().GetCategories(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetCategoryById(ctx context.Context, id string) (*model.Category, error) {
	data, err := svc.database.Categories().GetCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrCategoryNotFound
	}

	return data, nil
}

func (svc *service) GetCategoryByName(ctx context.Context, language string, name string) (*model.Category, error) {
	normalizedLanguage := svc.stringNormalizer.NormalizeString(language)

	data, err := svc.database.Categories().GetCategoryByName(ctx, normalizedLanguage, name)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrCategoryNotFound
	}

	return data, nil
}

func (svc *service) GetCategoryBySlug(ctx context.Context, language string, slug string) (*model.Category, error) {
	normalizedLanguage := svc.stringNormalizer.NormalizeString(language)
	normalizedSlug := svc.stringNormalizer.NormalizeString(slug)

	data, err := svc.database.Categories().GetCategoryBySlug(ctx, normalizedLanguage, normalizedSlug)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrCategoryNotFound
	}

	return data, nil
}

func (svc *service) CreateCategory(ctx context.Context, model *model.Category) (*model.Category, error) {
	//TODO: Check for duplicas

	// Normalize translations
	for _, translation := range model.Translations {
		translation.Language = svc.stringNormalizer.NormalizeString(translation.Language)
		translation.Slug = svc.stringNormalizer.NormalizeString(translation.Slug)
	}

	newId, err := svc.database.Categories().CreateCategory(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetCategoryById(ctx, newId)
}

func (svc *service) UpdateCategory(ctx context.Context, id string, model *model.Category) (*model.Category, error) {
	data, err := svc.database.Categories().GetCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrCategoryNotFound
	}

	//TODO: Set fields

	err = svc.database.Categories().UpdateCategory(ctx, data)
	if err != nil {
		return nil, err
	}

	return svc.GetCategoryById(ctx, id)
}

func (svc *service) DeleteCategory(ctx context.Context, id string) error {
	data, err := svc.database.Categories().GetCategoryById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return product.ErrCategoryNotFound
	}

	//TODO: Check dependencies

	err = svc.database.Categories().DeleteCategory(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

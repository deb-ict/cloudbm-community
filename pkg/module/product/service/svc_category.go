package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
)

func (svc *service) GetCategories(ctx context.Context, offset int64, limit int64, filter *model.CategoryFilter, sort *core.Sort) ([]*model.Category, int64, error) {
	filter.Language = localization.NormalizeLanguage(filter.Language)
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
	normalizedLanguage := localization.NormalizeLanguage(language)
	normalizedName := svc.stringNormalizer.NormalizeString(name)

	data, err := svc.database.Categories().GetCategoryByName(ctx, normalizedLanguage, normalizedName)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrCategoryNotFound
	}

	return data, nil
}

func (svc *service) GetCategoryBySlug(ctx context.Context, language string, slug string) (*model.Category, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
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
	model.Normalize(svc.stringNormalizer)
	model.Id = ""

	err := svc.checkDuplicateCategory(ctx, model)
	if err != nil {
		return nil, err
	}

	newId, err := svc.database.Categories().CreateCategory(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetCategoryById(ctx, newId)
}

func (svc *service) UpdateCategory(ctx context.Context, id string, model *model.Category) (*model.Category, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = id

	err := svc.checkDuplicateCategory(ctx, model)
	if err != nil {
		return nil, err
	}

	data, err := svc.database.Categories().GetCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrCategoryNotFound
	}
	data.UpdateModel(model)

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

	err = svc.database.Categories().DeleteCategory(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (svc *service) checkDuplicateCategory(ctx context.Context, model *model.Category) error {
	for _, translation := range model.Translations {
		if err := svc.checkDuplicateCategoryName(ctx, model, translation); err != nil {
			return err
		}
		if err := svc.checkDuplicateCategorySlug(ctx, model, translation); err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) checkDuplicateCategoryName(ctx context.Context, model *model.Category, translation *model.CategoryTranslation) error {
	duplicate, err := svc.database.Categories().GetCategoryByName(ctx, translation.Language, translation.NormalizedName)
	if err != nil {
		return err
	}
	if duplicate != nil && duplicate.Id != model.Id {
		return product.ErrCategoryDuplicateName
	}
	return nil
}

func (svc *service) checkDuplicateCategorySlug(ctx context.Context, model *model.Category, translation *model.CategoryTranslation) error {
	duplicate, err := svc.database.Categories().GetCategoryBySlug(ctx, translation.Language, translation.Slug)
	if err != nil {
		return err
	}
	if duplicate != nil && duplicate.Id != model.Id {
		return product.ErrCategoryDuplicateSlug
	}
	return nil
}

package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
)

func (svc *service) GetCategories(ctx context.Context, offset int64, limit int64, filter *model.CategoryFilter, sort *core.Sort) ([]*model.Category, int64, error) {
	data, count, err := svc.database.CategoryRepository().GetCategories(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetCategoryById(ctx context.Context, id string) (*model.Category, error) {
	data, err := svc.database.CategoryRepository().GetCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrCategoryNotFound
	}

	return data, nil
}

func (svc *service) GetCategoryByName(ctx context.Context, language string, name string) (*model.Category, error) {
	data, err := svc.database.CategoryRepository().GetCategoryByName(ctx, language, name)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrCategoryNotFound
	}

	return data, nil
}

func (svc *service) GetCategoryBySlug(ctx context.Context, language string, slug string) (*model.Category, error) {
	data, err := svc.database.CategoryRepository().GetCategoryBySlug(ctx, language, slug)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrCategoryNotFound
	}

	return data, nil
}

func (svc *service) CreateCategory(ctx context.Context, model *model.Category) (*model.Category, error) {
	newId, err := svc.database.CategoryRepository().CreateCategory(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetCategoryById(ctx, newId)
}

func (svc *service) UpdateCategory(ctx context.Context, id string, model *model.Category) (*model.Category, error) {
	data, err := svc.database.CategoryRepository().GetCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrCategoryNotFound
	}

	model.Id = id
	err = svc.database.CategoryRepository().UpdateCategory(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetCategoryById(ctx, id)
}

func (svc *service) DeleteCategory(ctx context.Context, id string) error {
	data, err := svc.database.CategoryRepository().GetCategoryById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return product.ErrCategoryNotFound
	}

	err = svc.database.CategoryRepository().DeleteCategory(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
)

func (svc *service) GetProducts(ctx context.Context, offset int64, limit int64, filter *model.ProductFilter, sort *core.Sort) ([]*model.Product, int64, error) {
	data, count, err := svc.database.ProductRepository().GetProducts(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetProductById(ctx context.Context, id string) (*model.Product, error) {
	data, err := svc.database.ProductRepository().GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrProductNotFound
	}

	return data, nil
}

func (svc *service) GetProductByName(ctx context.Context, language string, name string) (*model.Product, error) {
	data, err := svc.database.ProductRepository().GetProductByName(ctx, language, name)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrProductNotFound
	}

	return data, nil
}

func (svc *service) GetProductBySlug(ctx context.Context, language string, slug string) (*model.Product, error) {
	data, err := svc.database.ProductRepository().GetProductBySlug(ctx, language, slug)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrProductNotFound
	}

	return data, nil
}

func (svc *service) CreateProduct(ctx context.Context, model *model.Product) (*model.Product, error) {
	//TODO: Check for duplicas

	newId, err := svc.database.ProductRepository().CreateProduct(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetProductById(ctx, newId)
}

func (svc *service) UpdateProduct(ctx context.Context, id string, model *model.Product) (*model.Product, error) {
	data, err := svc.database.ProductRepository().GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrProductNotFound
	}

	//TODO: Set fields

	err = svc.database.ProductRepository().UpdateProduct(ctx, data)
	if err != nil {
		return nil, err
	}

	return svc.GetProductById(ctx, id)
}

func (svc *service) DeleteProduct(ctx context.Context, id string) error {
	data, err := svc.database.ProductRepository().GetProductById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return product.ErrProductNotFound
	}

	err = svc.database.ProductRepository().DeleteProduct(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

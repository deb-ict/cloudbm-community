package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
)

func (svc *service) GetProducts(ctx context.Context, offset int64, limit int64, filter *model.ProductFilter, sort *core.Sort) ([]*model.Product, int64, error) {
	filter.Language = localization.NormalizeLanguage(filter.Language)
	data, count, err := svc.database.Products().GetProducts(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetProductById(ctx context.Context, id string) (*model.Product, error) {
	data, err := svc.database.Products().GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrProductNotFound
	}

	return data, nil
}

func (svc *service) GetProductByName(ctx context.Context, language string, name string) (*model.Product, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	normalizedName := svc.stringNormalizer.NormalizeString(name)

	data, err := svc.database.Products().GetProductByName(ctx, normalizedLanguage, normalizedName)
	if err != nil {
		return nil, err
	}
	if data == nil {
		if !svc.languageProvider.IsDefaultLanguage(ctx, normalizedLanguage) {
			return svc.GetProductByName(ctx, svc.languageProvider.DefaultLanguage(ctx), name)
		}
		return nil, product.ErrProductNotFound
	}

	return data, nil
}

func (svc *service) GetProductBySlug(ctx context.Context, language string, slug string) (*model.Product, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	normalizedSlug := svc.stringNormalizer.NormalizeString(slug)

	data, err := svc.database.Products().GetProductBySlug(ctx, normalizedLanguage, normalizedSlug)
	if err != nil {
		return nil, err
	}
	if data == nil {
		if !svc.languageProvider.IsDefaultLanguage(ctx, normalizedLanguage) {
			return svc.GetProductBySlug(ctx, svc.languageProvider.DefaultLanguage(ctx), slug)
		}
		return nil, product.ErrProductNotFound
	}

	return data, nil
}

func (svc *service) CreateProduct(ctx context.Context, model *model.Product) (*model.Product, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = ""

	err := svc.checkDuplicateProduct(ctx, model)
	if err != nil {
		return nil, err
	}

	newId, err := svc.database.Products().CreateProduct(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetProductById(ctx, newId)
}

func (svc *service) UpdateProduct(ctx context.Context, id string, model *model.Product) (*model.Product, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = id

	err := svc.checkDuplicateProduct(ctx, model)
	if err != nil {
		return nil, err
	}

	data, err := svc.database.Products().GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrProductNotFound
	}
	data.UpdateModel(model)

	err = svc.database.Products().UpdateProduct(ctx, data)
	if err != nil {
		return nil, err
	}

	return svc.GetProductById(ctx, id)
}

func (svc *service) DeleteProduct(ctx context.Context, id string) error {
	data, err := svc.database.Products().GetProductById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return product.ErrProductNotFound
	}

	err = svc.database.Products().DeleteProduct(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (svc *service) checkDuplicateProduct(ctx context.Context, model *model.Product) error {
	for _, translation := range model.Translations {
		if err := svc.checkDuplicateProductName(ctx, model, translation); err != nil {
			return err
		}
		if err := svc.checkDuplicateProductSlug(ctx, model, translation); err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) checkDuplicateProductName(ctx context.Context, model *model.Product, translation *model.ProductTranslation) error {
	duplicate, err := svc.database.Products().GetProductByName(ctx, translation.Language, translation.NormalizedName)
	if err != nil {
		return err
	}
	if duplicate != nil && duplicate.Id != model.Id {
		return product.ErrProductDuplicateName
	}
	return nil
}

func (svc *service) checkDuplicateProductSlug(ctx context.Context, model *model.Product, translation *model.ProductTranslation) error {
	duplicate, err := svc.database.Products().GetProductBySlug(ctx, translation.Language, translation.Slug)
	if err != nil {
		return err
	}
	if duplicate != nil && duplicate.Id != model.Id {
		return product.ErrProductDuplicateSlug
	}
	return nil
}

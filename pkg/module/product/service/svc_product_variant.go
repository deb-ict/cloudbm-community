package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
)

func (svc *service) GetProductVariants(ctx context.Context, productId string, offset int64, limit int64, filter *model.ProductVariantFilter, sort *core.Sort) ([]*model.ProductVariant, int64, error) {
	filter.Language = localization.NormalizeLanguage(filter.Language)
	data, count, err := svc.database.ProductVariants().GetProductVariants(ctx, productId, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetProductVariantById(ctx context.Context, productId string, id string) (*model.ProductVariant, error) {
	data, err := svc.database.ProductVariants().GetProductVariantById(ctx, productId, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrProductVariantNotFound
	}

	return data, nil
}

func (svc *service) GetProductVariantByName(ctx context.Context, productId string, language string, name string) (*model.ProductVariant, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	normalizedName := svc.stringNormalizer.NormalizeString(name)

	data, err := svc.database.ProductVariants().GetProductVariantByName(ctx, productId, normalizedLanguage, normalizedName)
	if err != nil {
		return nil, err
	}
	if data == nil {
		if !svc.languageProvider.IsDefaultLanguage(ctx, normalizedLanguage) {
			return svc.GetProductVariantByName(ctx, productId, svc.languageProvider.DefaultLanguage(ctx), name)
		}
		return nil, product.ErrProductVariantNotFound
	}

	return data, nil
}

func (svc *service) GetProductVariantBySlug(ctx context.Context, productId string, language string, slug string) (*model.ProductVariant, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	normalizedSlug := svc.stringNormalizer.NormalizeString(slug)

	data, err := svc.database.ProductVariants().GetProductVariantBySlug(ctx, productId, normalizedLanguage, normalizedSlug)
	if err != nil {
		return nil, err
	}
	if data == nil {
		if !svc.languageProvider.IsDefaultLanguage(ctx, normalizedLanguage) {
			return svc.GetProductVariantBySlug(ctx, productId, svc.languageProvider.DefaultLanguage(ctx), slug)
		}
		return nil, product.ErrProductVariantNotFound
	}

	return data, nil
}

func (svc *service) CreateProductVariant(ctx context.Context, productId string, model *model.ProductVariant) (*model.ProductVariant, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = ""
	model.ProductId = productId

	err := svc.checkDuplicateProductVariant(ctx, model)
	if err != nil {
		return nil, err
	}

	newId, err := svc.database.ProductVariants().CreateProductVariant(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetProductVariantById(ctx, productId, newId)
}

func (svc *service) UpdateProductVariant(ctx context.Context, productId string, id string, model *model.ProductVariant) (*model.ProductVariant, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = id
	model.ProductId = productId

	err := svc.checkDuplicateProductVariant(ctx, model)
	if err != nil {
		return nil, err
	}

	data, err := svc.database.ProductVariants().GetProductVariantById(ctx, productId, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrProductVariantNotFound
	}
	data.UpdateModel(model)

	err = svc.database.ProductVariants().UpdateProductVariant(ctx, data)
	if err != nil {
		return nil, err
	}

	return svc.GetProductVariantById(ctx, productId, id)
}

func (svc *service) DeleteProductVariant(ctx context.Context, productId string, id string) error {
	data, err := svc.database.ProductVariants().GetProductVariantById(ctx, productId, id)
	if err != nil {
		return err
	}
	if data == nil {
		return product.ErrProductVariantNotFound
	}

	err = svc.database.ProductVariants().DeleteProductVariant(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (svc *service) checkDuplicateProductVariant(ctx context.Context, model *model.ProductVariant) error {
	for _, translation := range model.Details.Translations {
		if err := svc.checkDuplicateProductVariantByName(ctx, model, translation); err != nil {
			return err
		}
		if err := svc.checkDuplicateProductVariantBySlug(ctx, model, translation); err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) checkDuplicateProductVariantByName(ctx context.Context, model *model.ProductVariant, translation *model.ProductTranslation) error {
	duplicate, err := svc.database.ProductVariants().GetProductVariantByName(ctx, model.ProductId, translation.Language, translation.NormalizedName)
	if err != nil {
		return err
	}
	if duplicate != nil && duplicate.Id != model.Id {
		return product.ErrProductVariantDuplicateName
	}
	return nil
}

func (svc *service) checkDuplicateProductVariantBySlug(ctx context.Context, model *model.ProductVariant, translation *model.ProductTranslation) error {
	duplicate, err := svc.database.ProductVariants().GetProductVariantBySlug(ctx, model.ProductId, translation.Language, translation.Slug)
	if err != nil {
		return err
	}
	if duplicate != nil && duplicate.Id != model.Id {
		return product.ErrProductVariantDuplicateSlug
	}
	return nil
}

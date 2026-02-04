package service

import (
	"context"
	"log/slog"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
)

func (svc *service) GetAttributeValues(ctx context.Context, attributeId string, offset int64, limit int64, filter *model.AttributeValueFilter, sort *core.Sort) ([]*model.AttributeValue, int64, error) {
	parent, err := svc.GetAttributeById(ctx, attributeId)
	if err != nil {
		return nil, 0, err
	}

	filter.Language = localization.NormalizeLanguage(filter.Language)
	data, count, err := svc.database.AttributeValues().GetAttributeValues(ctx, parent, offset, limit, filter, sort)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get attribute values from the database",
			slog.String("attributeId", attributeId),
			slog.Any("error", err),
		)
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetAttributeValueById(ctx context.Context, attributeId string, id string) (*model.AttributeValue, error) {
	parent, err := svc.GetAttributeById(ctx, attributeId)
	if err != nil {
		return nil, err
	}

	data, err := svc.database.AttributeValues().GetAttributeValueById(ctx, parent, id)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get attribute value from database by id",
			slog.String("attributeId", attributeId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, product.ErrAttributeValueNotFound
	}

	return data, nil
}

func (svc *service) GetAttributeValueByName(ctx context.Context, attributeId string, language string, name string) (*model.AttributeValue, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	normalizedName := svc.stringNormalizer.NormalizeString(name)

	parent, err := svc.GetAttributeById(ctx, attributeId)
	if err != nil {
		return nil, err
	}

	data, err := svc.database.AttributeValues().GetAttributeValueByName(ctx, parent, normalizedLanguage, normalizedName)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get attribute value from database by name",
			slog.String("attributeId", attributeId),
			slog.String("language", normalizedLanguage),
			slog.String("name", normalizedName),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		if !svc.languageProvider.IsDefaultLanguage(ctx, normalizedLanguage) {
			return svc.GetAttributeValueByName(ctx, attributeId, svc.languageProvider.DefaultLanguage(ctx), name)
		}
		return nil, product.ErrAttributeValueNotFound
	}

	return data, nil
}

func (svc *service) GetAttributeValueBySlug(ctx context.Context, attributeId string, language string, slug string) (*model.AttributeValue, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	normalizedSlug := svc.stringNormalizer.NormalizeString(slug)

	parent, err := svc.GetAttributeById(ctx, attributeId)
	if err != nil {
		return nil, err
	}

	data, err := svc.database.AttributeValues().GetAttributeValueBySlug(ctx, parent, normalizedLanguage, normalizedSlug)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get attribute value from database by slug",
			slog.String("attributeId", attributeId),
			slog.String("language", normalizedLanguage),
			slog.String("slug", normalizedSlug),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		if !svc.languageProvider.IsDefaultLanguage(ctx, normalizedLanguage) {
			return svc.GetAttributeValueBySlug(ctx, attributeId, svc.languageProvider.DefaultLanguage(ctx), slug)
		}
		return nil, product.ErrAttributeValueNotFound
	}

	return data, nil
}

func (svc *service) CreateAttributeValue(ctx context.Context, attributeId string, model *model.AttributeValue) (*model.AttributeValue, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = ""
	model.AttributeId = attributeId

	parent, err := svc.GetAttributeById(ctx, attributeId)
	if err != nil {
		return nil, err
	}

	err = svc.checkDuplicateAttributeValue(ctx, parent, model)
	if err != nil {
		return nil, err
	}

	newId, err := svc.database.AttributeValues().CreateAttributeValue(ctx, parent, model)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to create attribute value in database",
			slog.String("attributeId", attributeId),
			slog.Any("error", err),
		)
		return nil, err
	}

	return svc.GetAttributeValueById(ctx, attributeId, newId)
}

func (svc *service) UpdateAttributeValue(ctx context.Context, attributeId string, id string, model *model.AttributeValue) (*model.AttributeValue, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = id
	model.AttributeId = attributeId

	parent, err := svc.GetAttributeById(ctx, attributeId)
	if err != nil {
		return nil, err
	}

	err = svc.checkDuplicateAttributeValue(ctx, parent, model)
	if err != nil {
		return nil, err
	}

	data, err := svc.database.AttributeValues().GetAttributeValueById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, product.ErrAttributeValueNotFound
	}
	data.UpdateModel(model)

	err = svc.database.AttributeValues().UpdateAttributeValue(ctx, parent, data)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to update attribute value in database",
			slog.String("attributeId", attributeId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}

	return svc.GetAttributeValueById(ctx, attributeId, id)
}

func (svc *service) DeleteAttributeValue(ctx context.Context, attributeId string, id string) error {
	parent, err := svc.GetAttributeById(ctx, attributeId)
	if err != nil {
		return err
	}

	data, err := svc.database.AttributeValues().GetAttributeValueById(ctx, parent, id)
	if err != nil {
		return err
	}
	if data == nil {
		return product.ErrAttributeValueNotFound
	}

	err = svc.database.AttributeValues().DeleteAttributeValue(ctx, parent, data)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to delete attribute value in database",
			slog.String("attributeId", attributeId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}

	return nil
}

func (svc *service) checkDuplicateAttributeValue(ctx context.Context, parent *model.Attribute, model *model.AttributeValue) error {
	for _, translation := range model.Translations {
		if err := svc.checkDuplicateAttributeValueByName(ctx, parent, model, translation); err != nil {
			return err
		}
		if err := svc.checkDuplicateAttributeValueBySlug(ctx, parent, model, translation); err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) checkDuplicateAttributeValueByName(ctx context.Context, parent *model.Attribute, model *model.AttributeValue, translation *model.AttributeValueTranslation) error {
	duplicate, err := svc.database.AttributeValues().GetAttributeValueByName(ctx, parent, translation.Language, translation.NormalizedName)
	if err != nil {
		return err
	}
	if duplicate != nil && duplicate.Id != model.Id {
		return product.ErrAttributeValueDuplicateName
	}
	return nil
}

func (svc *service) checkDuplicateAttributeValueBySlug(ctx context.Context, parent *model.Attribute, model *model.AttributeValue, translation *model.AttributeValueTranslation) error {
	duplicate, err := svc.database.AttributeValues().GetAttributeValueBySlug(ctx, parent, translation.Language, translation.Slug)
	if err != nil {
		return err
	}
	if duplicate != nil && duplicate.Id != model.Id {
		return product.ErrAttributeValueDuplicateSlug
	}
	return nil
}

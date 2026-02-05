package service

import (
	"context"
	"log/slog"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
)

func (svc *service) GetAttributes(ctx context.Context, offset int64, limit int64, filter *model.AttributeFilter, sort *core.Sort) ([]*model.Attribute, int64, error) {
	filter.Language = localization.NormalizeLanguage(filter.Language)
	data, count, err := svc.database.Attributes().GetAttributes(ctx, offset, limit, filter, sort)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get attributes from database",
			slog.Any("error", err),
		)
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetAttributeById(ctx context.Context, id string) (*model.Attribute, error) {
	data, err := svc.database.Attributes().GetAttributeById(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get attribute from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, product.ErrAttributeNotFound
	}

	return data, nil
}

func (svc *service) GetAttributeByName(ctx context.Context, language string, name string) (*model.Attribute, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	normalizedName := svc.stringNormalizer.NormalizeString(name)

	data, err := svc.database.Attributes().GetAttributeByName(ctx, normalizedLanguage, normalizedName)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get attribute from database by name",
			slog.String("language", normalizedLanguage),
			slog.String("name", normalizedName),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		if !svc.languageProvider.IsDefaultLanguage(ctx, normalizedLanguage) {
			return svc.GetAttributeByName(ctx, svc.languageProvider.DefaultLanguage(ctx), name)
		}
		return nil, product.ErrAttributeNotFound
	}

	return data, nil
}

func (svc *service) GetAttributeBySlug(ctx context.Context, language string, slug string) (*model.Attribute, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	normalizedSlug := svc.stringNormalizer.NormalizeString(slug)

	data, err := svc.database.Attributes().GetAttributeBySlug(ctx, normalizedLanguage, normalizedSlug)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get attribute from database by slug",
			slog.String("language", normalizedLanguage),
			slog.String("slug", normalizedSlug),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		if !svc.languageProvider.IsDefaultLanguage(ctx, normalizedLanguage) {
			return svc.GetAttributeBySlug(ctx, svc.languageProvider.DefaultLanguage(ctx), slug)
		}
		return nil, product.ErrAttributeNotFound
	}

	return data, nil
}

func (svc *service) CreateAttribute(ctx context.Context, model *model.Attribute) (*model.Attribute, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = ""

	err := svc.checkDuplicateAttribute(ctx, model)
	if err != nil {
		return nil, err
	}

	newId, err := svc.database.Attributes().CreateAttribute(ctx, model)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to create attribute in database",
			slog.Any("error", err),
		)
		return nil, err
	}

	return svc.GetAttributeById(ctx, newId)
}

func (svc *service) UpdateAttribute(ctx context.Context, id string, model *model.Attribute) (*model.Attribute, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = id

	err := svc.checkDuplicateAttribute(ctx, model)
	if err != nil {
		return nil, err
	}

	data, err := svc.database.Attributes().GetAttributeById(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get attribute from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, product.ErrAttributeNotFound
	}
	data.UpdateModel(model)

	err = svc.database.Attributes().UpdateAttribute(ctx, data)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to update attribute in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}

	return svc.GetAttributeById(ctx, id)
}

func (svc *service) DeleteAttribute(ctx context.Context, id string) error {
	data, err := svc.database.Attributes().GetAttributeById(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get attribute from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	if data == nil {
		return product.ErrAttributeNotFound
	}

	err = svc.database.Attributes().DeleteAttribute(ctx, data)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to delete attribute in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}

	return nil
}

func (svc *service) checkDuplicateAttribute(ctx context.Context, model *model.Attribute) error {
	for _, translation := range model.Translations {
		if err := svc.checkDuplicateAttributeByName(ctx, model, translation); err != nil {
			return err
		}
		if err := svc.checkDuplicateAttributeBySlug(ctx, model, translation); err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) checkDuplicateAttributeByName(ctx context.Context, model *model.Attribute, translation *model.AttributeTranslation) error {
	duplicate, err := svc.database.Attributes().GetAttributeByName(ctx, translation.Language, translation.NormalizedName)
	if err != nil {
		return err
	}
	if duplicate != nil && duplicate.Id != model.Id {
		return product.ErrAttributeDuplicateName
	}
	return nil
}

func (svc *service) checkDuplicateAttributeBySlug(ctx context.Context, model *model.Attribute, translation *model.AttributeTranslation) error {
	duplicate, err := svc.database.Attributes().GetAttributeBySlug(ctx, translation.Language, translation.Slug)
	if err != nil {
		return err
	}
	if duplicate != nil && duplicate.Id != model.Id {
		return product.ErrAttributeDuplicateSlug
	}
	return nil
}

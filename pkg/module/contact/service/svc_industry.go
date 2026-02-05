package service

import (
	"context"
	"log/slog"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/logging"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetIndustries(ctx context.Context, offset int64, limit int64, filter *model.IndustryFilter, sort *core.Sort) ([]*model.Industry, int64, error) {
	data, count, err := svc.database.Industries().GetIndustries(ctx, offset, limit, filter, sort)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get industries from database",
			slog.Any("error", err),
		)
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetIndustryById(ctx context.Context, id string) (*model.Industry, error) {
	data, err := svc.database.Industries().GetIndustryById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get industry from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrIndustryNotFound
	}

	return data, nil
}

func (svc *service) CreateIndustry(ctx context.Context, model *model.Industry) (*model.Industry, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = ""

	err := svc.validateIndustryName(ctx, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate industry",
			slog.Any("error", err),
		)
		return nil, err
	}

	newId, err := svc.database.Industries().CreateIndustry(ctx, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to create industry in database",
			slog.Any("error", err),
		)
		return nil, err
	}

	return svc.GetIndustryById(ctx, newId)
}

func (svc *service) UpdateIndustry(ctx context.Context, id string, model *model.Industry) (*model.Industry, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = id

	data, err := svc.database.Industries().GetIndustryById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get industry from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrIndustryNotFound
	}
	data.UpdateModel(model)

	err = svc.validateIndustryName(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate industry",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}

	err = svc.database.Industries().UpdateIndustry(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to update industry in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	return svc.GetIndustryById(ctx, id)
}

func (svc *service) DeleteIndustry(ctx context.Context, id string) error {
	data, err := svc.database.Industries().GetIndustryById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get industry from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	if data == nil {
		return contact.ErrIndustryNotFound
	}
	if data.IsSystem {
		return contact.ErrIndustryReadOnly
	}

	err = svc.database.Industries().DeleteIndustry(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to delete industry in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	return nil
}

func (svc *service) validateIndustryName(ctx context.Context, model *model.Industry) error {
	if model.IsTransient() {
		existing, err := svc.database.Industries().GetIndustryByKey(ctx, model.Key)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrIndustryDuplicateKey
		}
	}

	for _, translation := range model.Translations {
		existing, err := svc.database.Industries().GetIndustryByName(ctx, translation.Language, translation.Name)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrIndustryDuplicateName
		}
	}
	return nil
}

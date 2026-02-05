package service

import (
	"context"
	"log/slog"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/logging"
	"github.com/deb-ict/cloudbm-community/pkg/module/metadata"
	"github.com/deb-ict/cloudbm-community/pkg/module/metadata/model"
)

func (svc *service) GetTaxRates(ctx context.Context, offset int64, limit int64, filter *model.TaxRateFilter, sort *core.Sort) ([]*model.TaxRate, int64, error) {
	data, count, err := svc.database.TaxRates().GetTaxRates(ctx, offset, limit, filter, sort)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get tax rates from database",
			slog.Any("error", err),
		)
		return nil, 0, err
	}
	return data, count, nil
}

func (svc *service) GetTaxRateById(ctx context.Context, id string) (*model.TaxRate, error) {
	data, err := svc.database.TaxRates().GetTaxRateById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get tax rate from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, metadata.ErrTaxRateNotFound
	}
	return data, nil
}

func (svc *service) CreateTaxRate(ctx context.Context, model *model.TaxRate) (*model.TaxRate, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = ""

	err := svc.validateTaxRate(ctx, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate tax rate",
			slog.Any("error", err),
		)
		return nil, err
	}

	newId, err := svc.database.TaxRates().CreateTaxRate(ctx, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to create tax rate in database",
			slog.Any("error", err),
		)
		return nil, err
	}

	return svc.GetTaxRateById(ctx, newId)
}

func (svc *service) UpdateTaxRate(ctx context.Context, id string, model *model.TaxRate) (*model.TaxRate, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = id

	data, err := svc.database.TaxRates().GetTaxRateById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get tax rate from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, metadata.ErrTaxRateNotFound
	}
	data.UpdateModel(model)

	err = svc.validateTaxRate(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate tax rate",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}

	err = svc.database.TaxRates().UpdateTaxRate(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to update tax rate in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}

	return svc.GetTaxRateById(ctx, id)
}

func (svc *service) DeleteTaxRate(ctx context.Context, id string) error {
	data, err := svc.database.TaxRates().GetTaxRateById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get tax rate from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	if data == nil {
		return metadata.ErrTaxRateNotFound
	}

	err = svc.database.TaxRates().DeleteTaxRate(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to delete tax rate in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}

	return nil
}

func (svc *service) validateTaxRate(ctx context.Context, model *model.TaxRate) error {
	if model.IsTransient() {
		existing, err := svc.database.TaxRates().GetTaxRateByKey(ctx, model.Key)
		if err != nil {
			return err
		}
		if existing != nil {
			return metadata.ErrTaxRateDuplicateKey
		}
	}

	for _, translation := range model.Translations {
		existing, err := svc.database.TaxRates().GetTaxRateByName(ctx, translation.Language, translation.Name)
		if err != nil {
			return err
		}
		if existing != nil {
			return metadata.ErrTaxRateDuplicateName
		}
	}

	return nil
}

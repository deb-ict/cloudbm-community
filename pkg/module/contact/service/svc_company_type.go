package service

import (
	"context"
	"log/slog"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/logging"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetCompanyTypes(ctx context.Context, offset int64, limit int64, filter *model.CompanyTypeFilter, sort *core.Sort) ([]*model.CompanyType, int64, error) {
	data, count, err := svc.database.CompanyTypes().GetCompanyTypes(ctx, offset, limit, filter, sort)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company types from database",
			slog.Any("error", err),
		)
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetCompanyTypeById(ctx context.Context, id string) (*model.CompanyType, error) {
	data, err := svc.database.CompanyTypes().GetCompanyTypeById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company type from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyTypeNotFound
	}

	return data, nil
}

func (svc *service) CreateCompanyType(ctx context.Context, model *model.CompanyType) (*model.CompanyType, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = ""

	err := svc.validateCompanyType(ctx, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate company type",
			slog.Any("error", err),
		)
		return nil, err
	}

	newId, err := svc.database.CompanyTypes().CreateCompanyType(ctx, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to create company type in database",
			slog.Any("error", err),
		)
		return nil, err
	}

	return svc.GetCompanyTypeById(ctx, newId)
}

func (svc *service) UpdateCompanyType(ctx context.Context, id string, model *model.CompanyType) (*model.CompanyType, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = id

	data, err := svc.database.CompanyTypes().GetCompanyTypeById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company type from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyTypeNotFound
	}
	data.UpdateModel(model)

	err = svc.validateCompanyType(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate company type",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}

	err = svc.database.CompanyTypes().UpdateCompanyType(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to update company type in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	return svc.GetCompanyTypeById(ctx, id)
}

func (svc *service) DeleteCompanyType(ctx context.Context, id string) error {
	data, err := svc.database.CompanyTypes().GetCompanyTypeById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company type from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	if data == nil {
		return contact.ErrCompanyTypeNotFound
	}
	if data.IsSystem {
		return contact.ErrCompanyTypeReadOnly
	}

	err = svc.database.CompanyTypes().DeleteCompanyType(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to delete company type in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	return nil
}

func (svc *service) validateCompanyType(ctx context.Context, model *model.CompanyType) error {
	if model.IsTransient() {
		existing, err := svc.database.CompanyTypes().GetCompanyTypeByKey(ctx, model.Key)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrCompanyTypeDuplicateKey
		}
	}

	for _, translation := range model.Translations {
		existing, err := svc.database.CompanyTypes().GetCompanyTypeByName(ctx, translation.Language, translation.Name)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrCompanyTypeDuplicateName
		}
	}
	return nil
}

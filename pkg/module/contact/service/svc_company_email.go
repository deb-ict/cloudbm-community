package service

import (
	"context"
	"log/slog"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/logging"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetCompanyEmails(ctx context.Context, companyId string, offset int64, limit int64, filter *model.EmailFilter, sort *core.Sort) ([]*model.Email, int64, error) {
	parent, err := svc.database.Companies().GetCompanyById(ctx, companyId)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company from database by id",
			slog.String("companyId", companyId),
			slog.Any("error", err),
		)
		return nil, 0, err
	}
	if parent == nil {
		return nil, 0, contact.ErrCompanyNotFound
	}

	data, count, err := svc.database.CompanyEmails().GetCompanyEmails(ctx, parent, offset, limit, filter, sort)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company emails from database",
			slog.String("companyId", companyId),
			slog.Any("error", err),
		)
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetCompanyEmailById(ctx context.Context, companyId string, id string) (*model.Email, error) {
	parent, err := svc.database.Companies().GetCompanyById(ctx, companyId)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company from database by id",
			slog.String("companyId", companyId),
			slog.Any("error", err),
		)
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyEmails().GetCompanyEmailById(ctx, parent, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company email from database by id",
			slog.String("companyId", companyId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyAddressNotFound
	}

	return data, nil
}

func (svc *service) CreateCompanyEmail(ctx context.Context, companyId string, model *model.Email) (*model.Email, error) {
	model.Id = ""

	parent, err := svc.database.Companies().GetCompanyById(ctx, companyId)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company from database by id",
			slog.String("companyId", companyId),
			slog.Any("error", err),
		)
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	err = svc.validateCompanyEmail(ctx, parent, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate company email",
			slog.String("companyId", companyId),
			slog.Any("error", err),
		)
		return nil, err
	}

	if model.IsDefault {
		err = svc.resetDefaultCompanyEmail(ctx, parent, model)
		if err != nil {
			logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to reset default company email",
				slog.String("companyId", companyId),
				slog.Any("error", err),
			)
			return nil, err
		}
	}

	newId, err := svc.database.CompanyEmails().CreateCompanyEmail(ctx, parent, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to create company email in database",
			slog.String("companyId", companyId),
			slog.Any("error", err),
		)
		return nil, err
	}
	return svc.GetCompanyEmailById(ctx, companyId, newId)
}

func (svc *service) UpdateCompanyEmail(ctx context.Context, companyId string, id string, model *model.Email) (*model.Email, error) {
	model.Id = id

	parent, err := svc.database.Companies().GetCompanyById(ctx, companyId)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company from database by id",
			slog.String("companyId", companyId),
			slog.Any("error", err),
		)
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyEmails().GetCompanyEmailById(ctx, parent, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company email from database by id",
			slog.String("companyId", companyId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyEmailNotFound
	}
	data.UpdateModel(model)

	err = svc.validateCompanyEmail(ctx, parent, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate company email",
			slog.String("companyId", companyId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}

	if data.IsDefault {
		err = svc.resetDefaultCompanyEmail(ctx, parent, data)
		if err != nil {
			logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to reset default company email",
				slog.String("companyId", companyId),
				slog.String("id", id),
				slog.Any("error", err),
			)
			return nil, err
		}
	}

	err = svc.database.CompanyEmails().UpdateCompanyEmail(ctx, parent, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to update company email in database",
			slog.String("companyId", companyId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	return svc.GetCompanyEmailById(ctx, companyId, id)
}

func (svc *service) DeleteCompanyEmail(ctx context.Context, companyId string, id string) error {
	parent, err := svc.database.Companies().GetCompanyById(ctx, companyId)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company from database by id",
			slog.String("companyId", companyId),
			slog.Any("error", err),
		)
		return err
	}
	if parent == nil {
		return contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyEmails().GetCompanyEmailById(ctx, parent, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company email from database by id",
			slog.String("companyId", companyId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	if data == nil {
		return contact.ErrCompanyEmailNotFound
	}
	if data.IsDefault {
		return contact.ErrCompanyEmailIsDefault
	}

	err = svc.database.CompanyEmails().DeleteCompanyEmail(ctx, parent, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to delete company email in database",
			slog.String("companyId", companyId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	return nil
}

func (svc *service) resetDefaultCompanyEmail(ctx context.Context, parent *model.Company, model *model.Email) error {
	current, err := svc.database.CompanyEmails().GetDefaultCompanyEmail(ctx, parent)
	if err != nil {
		return err
	}
	if current != nil && current.Id != model.Id {
		current.IsDefault = false
		err = svc.database.CompanyEmails().UpdateCompanyEmail(ctx, parent, current)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) validateCompanyEmail(ctx context.Context, parent *model.Company, model *model.Email) error {
	modelType, err := svc.database.EmailTypes().GetEmailTypeById(ctx, model.Type.Id)
	if err != nil {
		return err
	}
	if modelType == nil {
		return contact.ErrEmailTypeNotFound
	}

	return nil
}

package service

import (
	"context"
	"log/slog"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/logging"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetCompanies(ctx context.Context, offset int64, limit int64, filter *model.CompanyFilter, sort *core.Sort) ([]*model.Company, int64, error) {
	data, count, err := svc.database.Companies().GetCompanies(ctx, offset, limit, filter, sort)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get companies from database",
			slog.Any("error", err),
		)
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetCompanyById(ctx context.Context, id string) (*model.Company, error) {
	data, err := svc.database.Companies().GetCompanyById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyNotFound
	}

	return data, nil
}

func (svc *service) CreateCompany(ctx context.Context, model *model.Company) (*model.Company, error) {
	model.Id = ""

	err := svc.validateCompany(ctx, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate company",
			slog.Any("error", err),
		)
		return nil, err
	}

	newId, err := svc.database.Companies().CreateCompany(ctx, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to create company in database",
			slog.Any("error", err),
		)
		return nil, err
	}

	return svc.GetCompanyById(ctx, newId)
}

func (svc *service) UpdateCompany(ctx context.Context, id string, model *model.Company) (*model.Company, error) {
	model.Id = id

	data, err := svc.database.Companies().GetCompanyById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyNotFound
	}
	data.UpdateModel(model)

	err = svc.validateCompany(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate company",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}

	err = svc.database.Companies().UpdateCompany(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to update company in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	return svc.GetCompanyById(ctx, id)
}

func (svc *service) DeleteCompany(ctx context.Context, id string) error {
	data, err := svc.database.Companies().GetCompanyById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get company from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	if data == nil {
		return contact.ErrCompanyNotFound
	}
	if data.IsSystem {
		return contact.ErrCompanyReadOnly
	}

	err = svc.database.Companies().DeleteCompany(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to delete company in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	return nil
}

func (svc *service) validateCompany(ctx context.Context, model *model.Company) error {
	if model.Type != nil {
		companyType, err := svc.database.CompanyTypes().GetCompanyTypeById(ctx, model.Type.Id)
		if err != nil {
			return err
		}
		if companyType == nil {
			return contact.ErrCompanyTypeNotFound
		}
	}

	if model.Industry != nil {
		industry, err := svc.database.Industries().GetIndustryById(ctx, model.Industry.Id)
		if err != nil {
			return err
		}
		if industry == nil {
			return contact.ErrIndustryNotFound
		}
	}

	for _, address := range model.Addresses {
		err := svc.validateCompanyAddress(ctx, model, address)
		if err != nil {
			return err
		}
	}
	for _, email := range model.Emails {
		err := svc.validateCompanyEmail(ctx, model, email)
		if err != nil {
			return err
		}
	}
	for _, phone := range model.Phones {
		err := svc.validateCompanyPhone(ctx, model, phone)
		if err != nil {
			return err
		}
	}
	for _, uri := range model.Uris {
		err := svc.validateCompanyUri(ctx, model, uri)
		if err != nil {
			return err
		}
	}

	return nil
}

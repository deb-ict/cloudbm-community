package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetCompanyEmails(ctx context.Context, companyId string, offset int64, limit int64, filter *model.EmailFilter, sort *core.Sort) ([]*model.Email, int64, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, 0, err
	}
	if parent == nil {
		return nil, 0, contact.ErrCompanyNotFound
	}

	data, count, err := svc.database.CompanyEmailRepository().GetCompanyEmails(ctx, parent, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetCompanyEmailById(ctx context.Context, companyId string, id string) (*model.Email, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyEmailRepository().GetCompanyEmailById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyAddressNotFound
	}

	return data, nil
}

func (svc *service) CreateCompanyEmail(ctx context.Context, companyId string, model *model.Email) (*model.Email, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	newId, err := svc.database.CompanyEmailRepository().CreateCompanyEmail(ctx, parent, model)
	if err != nil {
		return nil, err
	}
	return svc.GetCompanyEmailById(ctx, companyId, newId)
}

func (svc *service) UpdateCompanyEmail(ctx context.Context, companyId string, id string, model *model.Email) (*model.Email, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyEmailRepository().GetCompanyEmailById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyEmailNotFound
	}

	err = svc.database.CompanyEmailRepository().UpdateCompanyEmail(ctx, parent, data)
	if err != nil {
		return nil, err
	}
	return svc.GetCompanyEmailById(ctx, companyId, id)
}

func (svc *service) DeleteCompanyEmail(ctx context.Context, companyId string, id string) error {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return err
	}
	if parent == nil {
		return contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyEmailRepository().GetCompanyEmailById(ctx, parent, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrCompanyEmailNotFound
	}

	err = svc.database.CompanyEmailRepository().DeleteCompanyEmail(ctx, parent, data)
	if err != nil {
		return err
	}
	return nil
}

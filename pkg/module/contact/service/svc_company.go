package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetCompanies(ctx context.Context, offset int64, limit int64, filter *model.CompanyFilter, sort *core.Sort) ([]*model.Company, int64, error) {
	data, count, err := svc.database.CompanyRepository().GetCompanies(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetCompanyById(ctx context.Context, id string) (*model.Company, error) {
	data, err := svc.database.CompanyRepository().GetCompanyById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyNotFound
	}

	return data, nil
}

func (svc *service) CreateCompany(ctx context.Context, model *model.Company) (*model.Company, error) {
	newId, err := svc.database.CompanyRepository().CreateCompany(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetCompanyById(ctx, newId)
}

func (svc *service) UpdateCompany(ctx context.Context, id string, model *model.Company) (*model.Company, error) {
	data, err := svc.GetCompanyById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return contact.ErrXXXNotFound
	}

	//TODO: Set fields

	err = svc.database.CompanyRepository().UpdateCompany(ctx, data)
	if err != nil {
		return nil, err
	}
	return svc.GetCompanyById(ctx, id)
}

func (svc *service) DeleteCompany(ctx context.Context, id string) error {
	data, err := svc.GetCompanyById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrXXXNotFound
	}

	err = svc.database.CompanyRepository().DeleteCompany(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

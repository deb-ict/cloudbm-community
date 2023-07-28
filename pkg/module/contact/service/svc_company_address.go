package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetCompanyAddresses(ctx context.Context, companyId string, offset int64, limit int64, filter *model.AddressFilter, sort *core.Sort) ([]*model.Address, int64, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, 0, err
	}
	if parent == nil {
		return nil, 0, contact.ErrCompanyNotFound
	}
}

func (svc *service) GetCompanyAddressById(ctx context.Context, companyId string, id string) (*model.Address, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}
}

func (svc *service) CreateCompanyAddress(ctx context.Context, companyId string, model *model.Address) (*model.Address, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}
}

func (svc *service) UpdateCompanyAddress(ctx context.Context, companyId string, id string, model *model.Address) (*model.Address, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}
}

func (svc *service) DeleteCompanyAddress(ctx context.Context, companyId string, id string) error {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return err
	}
	if parent == nil {
		return contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyAddressRepository().GetCompanyAddressById(ctx, parent, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrContactXXXNotFound
	}

	err = svc.database.CompanyAddressRepository().DeleteCompanyAddress(ctx, parent, data)
	if err != nil {
		return err
	}
	return nil
}

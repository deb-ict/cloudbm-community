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

	data, count, err := svc.database.CompanyAddressRepository().GetCompanyAddresses(ctx, parent, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetCompanyAddressById(ctx context.Context, companyId string, id string) (*model.Address, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyAddressRepository().GetCompanyAddressById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyAddressNotFound
	}

	return data, nil
}

func (svc *service) CreateCompanyAddress(ctx context.Context, companyId string, model *model.Address) (*model.Address, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	newId, err := svc.database.CompanyAddressRepository().CreateCompanyAddress(ctx, parent, model)
	if err != nil {
		return nil, err
	}
	return svc.GetCompanyAddressById(ctx, companyId, newId)
}

func (svc *service) UpdateCompanyAddress(ctx context.Context, companyId string, id string, model *model.Address) (*model.Address, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyAddressRepository().GetCompanyAddressById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyAddressNotFound
	}

	err = svc.database.CompanyAddressRepository().UpdateCompanyAddress(ctx, parent, data)
	if err != nil {
		return nil, err
	}
	return svc.GetCompanyAddressById(ctx, companyId, id)
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
		return contact.ErrCompanyAddressNotFound
	}

	err = svc.database.CompanyAddressRepository().DeleteCompanyAddress(ctx, parent, data)
	if err != nil {
		return err
	}
	return nil
}

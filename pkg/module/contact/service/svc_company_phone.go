package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetCompanyPhones(ctx context.Context, companyId string, offset int64, limit int64, filter *model.PhoneFilter, sort *core.Sort) ([]*model.Phone, int64, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, 0, err
	}
	if parent == nil {
		return nil, 0, contact.ErrCompanyNotFound
	}

	data, count, err := svc.database.CompanyPhoneRepository().GetCompanyPhones(ctx, parent, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetCompanyPhoneById(ctx context.Context, companyId string, id string) (*model.Phone, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyPhoneRepository().GetCompanyPhoneById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyAddressNotFound
	}

	return data, nil
}

func (svc *service) CreateCompanyPhone(ctx context.Context, companyId string, model *model.Phone) (*model.Phone, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	newId, err := svc.database.CompanyPhoneRepository().CreateCompanyPhone(ctx, parent, model)
	if err != nil {
		return nil, err
	}
	return svc.GetCompanyPhoneById(ctx, companyId, newId)
}

func (svc *service) UpdateCompanyPhone(ctx context.Context, companyId string, id string, model *model.Phone) (*model.Phone, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyPhoneRepository().GetCompanyPhoneById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyPhoneNotFound
	}

	err = svc.database.CompanyPhoneRepository().UpdateCompanyPhone(ctx, parent, data)
	if err != nil {
		return nil, err
	}
	return svc.GetCompanyPhoneById(ctx, companyId, id)
}

func (svc *service) DeleteCompanyPhone(ctx context.Context, companyId string, id string) error {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return err
	}
	if parent == nil {
		return contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyPhoneRepository().GetCompanyPhoneById(ctx, parent, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrCompanyPhoneNotFound
	}

	err = svc.database.CompanyPhoneRepository().DeleteCompanyPhone(ctx, parent, data)
	if err != nil {
		return err
	}
	return nil
}

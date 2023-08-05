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

	err = svc.validateCompanyAddress(ctx, parent, model)
	if err != nil {
		return nil, err
	}

	if model.IsDefault {
		err = svc.resetDefaultCompanyAddress(ctx, parent, model)
		if err != nil {
			return nil, err
		}
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

	data.Type = model.Type
	data.Street = model.Street
	data.StreetNr = model.StreetNr
	data.Unit = model.Unit
	data.PostalCode = model.PostalCode
	data.City = model.City
	data.State = model.State
	data.Country = model.Country
	data.IsDefault = model.IsDefault

	err = svc.validateCompanyAddress(ctx, parent, data)
	if err != nil {
		return nil, err
	}

	if data.IsDefault {
		err = svc.resetDefaultCompanyAddress(ctx, parent, data)
		if err != nil {
			return nil, err
		}
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
	if data.IsDefault {
		return contact.ErrCompanyAddressIsDefault
	}

	err = svc.database.CompanyAddressRepository().DeleteCompanyAddress(ctx, parent, data)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) resetDefaultCompanyAddress(ctx context.Context, parent *model.Company, model *model.Address) error {
	current, err := svc.database.CompanyAddressRepository().GetDefaultCompanyAddress(ctx, parent)
	if err != nil {
		return err
	}
	if current != nil && current.Id != model.Id {
		current.IsDefault = false
		err = svc.database.CompanyAddressRepository().UpdateCompanyAddress(ctx, parent, current)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) validateCompanyAddress(ctx context.Context, parent *model.Company, model *model.Address) error {
	modelType, err := svc.database.AddressTypeRepository().GetAddressTypeById(ctx, model.Type.Id)
	if err != nil {
		return err
	}
	if modelType == nil {
		return contact.ErrAddressTypeNotFound
	}

	return nil
}

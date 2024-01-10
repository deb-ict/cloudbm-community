package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetCompanyPhones(ctx context.Context, companyId string, offset int64, limit int64, filter *model.PhoneFilter, sort *core.Sort) ([]*model.Phone, int64, error) {
	parent, err := svc.database.Companies().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, 0, err
	}
	if parent == nil {
		return nil, 0, contact.ErrCompanyNotFound
	}

	data, count, err := svc.database.CompanyPhones().GetCompanyPhones(ctx, parent, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetCompanyPhoneById(ctx context.Context, companyId string, id string) (*model.Phone, error) {
	parent, err := svc.database.Companies().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyPhones().GetCompanyPhoneById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyAddressNotFound
	}

	return data, nil
}

func (svc *service) CreateCompanyPhone(ctx context.Context, companyId string, model *model.Phone) (*model.Phone, error) {
	model.Id = ""

	parent, err := svc.database.Companies().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	err = svc.validateCompanyPhone(ctx, parent, model)
	if err != nil {
		return nil, err
	}

	if model.IsDefault {
		err = svc.resetDefaultCompanyPhone(ctx, parent, model)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.CompanyPhones().CreateCompanyPhone(ctx, parent, model)
	if err != nil {
		return nil, err
	}
	return svc.GetCompanyPhoneById(ctx, companyId, newId)
}

func (svc *service) UpdateCompanyPhone(ctx context.Context, companyId string, id string, model *model.Phone) (*model.Phone, error) {
	model.Id = id

	parent, err := svc.database.Companies().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyPhones().GetCompanyPhoneById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyPhoneNotFound
	}
	data.UpdateModel(model)

	err = svc.validateCompanyPhone(ctx, parent, data)
	if err != nil {
		return nil, err
	}

	if data.IsDefault {
		err = svc.resetDefaultCompanyPhone(ctx, parent, data)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.CompanyPhones().UpdateCompanyPhone(ctx, parent, data)
	if err != nil {
		return nil, err
	}
	return svc.GetCompanyPhoneById(ctx, companyId, id)
}

func (svc *service) DeleteCompanyPhone(ctx context.Context, companyId string, id string) error {
	parent, err := svc.database.Companies().GetCompanyById(ctx, companyId)
	if err != nil {
		return err
	}
	if parent == nil {
		return contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyPhones().GetCompanyPhoneById(ctx, parent, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrCompanyPhoneNotFound
	}
	if data.IsDefault {
		return contact.ErrCompanyPhoneIsDefault
	}

	err = svc.database.CompanyPhones().DeleteCompanyPhone(ctx, parent, data)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) resetDefaultCompanyPhone(ctx context.Context, parent *model.Company, model *model.Phone) error {
	current, err := svc.database.CompanyPhones().GetDefaultCompanyPhone(ctx, parent)
	if err != nil {
		return err
	}
	if current != nil && current.Id != model.Id {
		current.IsDefault = false
		err = svc.database.CompanyPhones().UpdateCompanyPhone(ctx, parent, current)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) validateCompanyPhone(ctx context.Context, parent *model.Company, model *model.Phone) error {
	modelType, err := svc.database.PhoneTypes().GetPhoneTypeById(ctx, model.Type.Id)
	if err != nil {
		return err
	}
	if modelType == nil {
		return contact.ErrPhoneTypeNotFound
	}

	return nil
}

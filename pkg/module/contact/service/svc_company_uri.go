package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetCompanyUris(ctx context.Context, companyId string, offset int64, limit int64, filter *model.UriFilter, sort *core.Sort) ([]*model.Uri, int64, error) {
	parent, err := svc.database.Companies().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, 0, err
	}
	if parent == nil {
		return nil, 0, contact.ErrCompanyNotFound
	}

	data, count, err := svc.database.CompanyUris().GetCompanyUris(ctx, parent, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetCompanyUriById(ctx context.Context, companyId string, id string) (*model.Uri, error) {
	parent, err := svc.database.Companies().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyUris().GetCompanyUriById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyAddressNotFound
	}

	return data, nil
}

func (svc *service) CreateCompanyUri(ctx context.Context, companyId string, model *model.Uri) (*model.Uri, error) {
	model.Id = ""

	parent, err := svc.database.Companies().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	err = svc.validateCompanyUri(ctx, parent, model)
	if err != nil {
		return nil, err
	}

	if model.IsDefault {
		err = svc.resetDefaultCompanyUri(ctx, parent, model)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.CompanyUris().CreateCompanyUri(ctx, parent, model)
	if err != nil {
		return nil, err
	}
	return svc.GetCompanyUriById(ctx, companyId, newId)
}

func (svc *service) UpdateCompanyUri(ctx context.Context, companyId string, id string, model *model.Uri) (*model.Uri, error) {
	model.Id = id

	parent, err := svc.database.Companies().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyUris().GetCompanyUriById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyUriNotFound
	}
	data.UpdateModel(model)

	err = svc.validateCompanyUri(ctx, parent, data)
	if err != nil {
		return nil, err
	}

	if data.IsDefault {
		err = svc.resetDefaultCompanyUri(ctx, parent, data)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.CompanyUris().UpdateCompanyUri(ctx, parent, data)
	if err != nil {
		return nil, err
	}
	return svc.GetCompanyUriById(ctx, companyId, id)
}

func (svc *service) DeleteCompanyUri(ctx context.Context, companyId string, id string) error {
	parent, err := svc.database.Companies().GetCompanyById(ctx, companyId)
	if err != nil {
		return err
	}
	if parent == nil {
		return contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyUris().GetCompanyUriById(ctx, parent, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrCompanyUriNotFound
	}
	if data.IsDefault {
		return contact.ErrCompanyUriIsDefault
	}

	err = svc.database.CompanyUris().DeleteCompanyUri(ctx, parent, data)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) resetDefaultCompanyUri(ctx context.Context, parent *model.Company, model *model.Uri) error {
	current, err := svc.database.CompanyUris().GetDefaultCompanyUri(ctx, parent)
	if err != nil {
		return err
	}
	if current != nil && current.Id != model.Id {
		current.IsDefault = false
		err = svc.database.CompanyUris().UpdateCompanyUri(ctx, parent, current)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) validateCompanyUri(ctx context.Context, parent *model.Company, model *model.Uri) error {
	modelType, err := svc.database.UriTypes().GetUriTypeById(ctx, model.Type.Id)
	if err != nil {
		return err
	}
	if modelType == nil {
		return contact.ErrUriTypeNotFound
	}

	return nil
}

package service

import (
	"context"
	"strings"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetCompanyTypes(ctx context.Context, offset int64, limit int64, filter *model.CompanyTypeFilter, sort *core.Sort) ([]*model.CompanyType, int64, error) {
	data, count, err := svc.database.CompanyTypeRepository().GetCompanyTypes(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetCompanyTypeById(ctx context.Context, id string) (*model.CompanyType, error) {
	data, err := svc.database.CompanyTypeRepository().GetCompanyTypeById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyTypeNotFound
	}

	return data, nil
}

func (svc *service) CreateCompanyType(ctx context.Context, model *model.CompanyType) (*model.CompanyType, error) {
	model.Key = strings.ToLower(model.Key)
	existing, err := svc.database.CompanyTypeRepository().GetCompanyTypeByKey(ctx, model.Key)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, contact.ErrCompanyTypeDuplicateKey
	}

	//TODO: Check for duplicates on name

	newId, err := svc.database.CompanyTypeRepository().CreateCompanyType(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetCompanyTypeById(ctx, newId)
}

func (svc *service) UpdateCompanyType(ctx context.Context, id string, model *model.CompanyType) (*model.CompanyType, error) {
	data, err := svc.GetCompanyTypeById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyTypeNotFound
	}

	//TODO: Check for duplicates on name

	//TODO: Set fields

	err = svc.database.CompanyTypeRepository().UpdateCompanyType(ctx, data)
	if err != nil {
		return nil, err
	}
	return svc.GetCompanyTypeById(ctx, id)
}

func (svc *service) DeleteCompanyType(ctx context.Context, id string) error {
	data, err := svc.GetCompanyTypeById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrCompanyTypeNotFound
	}

	//TODO: Check dependencies

	err = svc.database.CompanyTypeRepository().DeleteCompanyType(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetCompanyUris(ctx context.Context, companyId string, offset int64, limit int64, filter *model.UriFilter, sort *core.Sort) ([]*model.Uri, int64, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, 0, err
	}
	if parent == nil {
		return nil, 0, contact.ErrCompanyNotFound
	}
}

func (svc *service) GetCompanyUriById(ctx context.Context, companyId string, id string) (*model.Uri, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}
}

func (svc *service) CreateCompanyUri(ctx context.Context, companyId string, model *model.Uri) (*model.Uri, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}
}

func (svc *service) UpdateCompanyUri(ctx context.Context, companyId string, id string, model *model.Uri) (*model.Uri, error) {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrCompanyNotFound
	}
}

func (svc *service) DeleteCompanyUri(ctx context.Context, companyId string, id string) error {
	parent, err := svc.database.CompanyRepository().GetCompanyById(ctx, companyId)
	if err != nil {
		return err
	}
	if parent == nil {
		return contact.ErrCompanyNotFound
	}

	data, err := svc.database.CompanyUriRepository().GetCompanyUriById(ctx, parent, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrContactXXXNotFound
	}

	err = svc.database.CompanyUriRepository().DeleteCompanyUri(ctx, parent, data)
	if err != nil {
		return err
	}
	return nil
}

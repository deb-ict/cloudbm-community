package contact

import (
	"context"
)

func (svc *service) GetCompanies(ctx context.Context, pageIndex int, pageSize int) (*CompanyList, error) {
	return svc.database.GetCompanyStore().GetCompanies(ctx, pageIndex, pageSize)
}

func (svc *service) GetCompanyById(ctx context.Context, id string) (*Company, error) {
	//TODO: Validate id

	return svc.database.GetCompanyStore().GetCompanyById(ctx, id)
}

func (svc *service) CreateCompany(ctx context.Context, company Company) (*Company, error) {
	var err error

	//TODO: Validate model

	duplicate, err := svc.database.GetCompanyStore().GetCompanyByName(ctx, company.Name)
	if duplicate != nil {
		return nil, ErrCompanyDuplicateName
	}
	if err != nil {
		return nil, err
	}

	newId, err := svc.database.GetCompanyStore().CreateCompany(ctx, company)
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyStore().GetCompanyById(ctx, newId)
}

func (svc *service) UpdateCompany(ctx context.Context, id string, company Company) (*Company, error) {
	var err error

	//TODO: Validate id
	//TODO: Validate model

	existing, err := svc.database.GetCompanyStore().GetCompanyById(ctx, id)
	if existing == nil && err != nil {
		return nil, ErrCompanyNotFound
	}
	if err != nil {
		return nil, err
	}

	err = svc.database.GetCompanyStore().UpdateCompany(ctx, id, company)
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyStore().GetCompanyById(ctx, id)
}

func (svc *service) DeleteCompany(ctx context.Context, id string) error {
	var err error

	//TODO: Validate id

	existing, err := svc.database.GetCompanyStore().GetCompanyById(ctx, id)
	if existing == nil && err != nil {
		return ErrCompanyNotFound
	}
	if err != nil {
		return err
	}

	return svc.database.GetCompanyStore().DeleteCompany(ctx, id)
}

package contact

import (
	"context"
)

func (svc *service) GetCompanyEmails(ctx context.Context, companyId string, pageIndex int, pageSize int) (*EmailList, error) {
	var err error

	//TODO: Validate companyId

	company, err := svc.database.GetCompanyStore().GetCompanyById(ctx, companyId)
	if company == nil && err == nil {
		return nil, ErrCompanyNotFound
	}
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyEmailStore().GetCompanyEmails(ctx, companyId, pageIndex, pageSize)
}

func (svc *service) GetCompanyEmailById(ctx context.Context, companyId string, id string) (*Email, error) {
	var err error

	//TODO: Validate companyId
	//TODO: Validate id

	company, err := svc.database.GetCompanyStore().GetCompanyById(ctx, companyId)
	if company == nil && err == nil {
		return nil, ErrCompanyNotFound
	}
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyEmailStore().GetCompanyEmailById(ctx, companyId, id)
}

func (svc *service) CreateCompanyEmail(ctx context.Context, companyId string, email Email) (*Email, error) {
	var err error

	//TODO: Validate companyId
	//TODO: Validate id

	company, err := svc.database.GetCompanyStore().GetCompanyById(ctx, companyId)
	if company == nil && err == nil {
		return nil, ErrCompanyNotFound
	}
	if err != nil {
		return nil, err
	}

	emailType, err := svc.database.GetEmailTypeStore().GetEmailTypeById(ctx, email.TypeId)
	if emailType == nil && err == nil {
		return nil, ErrEmailTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	duplicate, err := svc.database.GetCompanyEmailStore().GetCompanyEmailByType(ctx, companyId, email.TypeId)
	if duplicate != nil {
		return nil, ErrCompanyEmailDuplicate
	}
	if err != nil && err != ErrCompanyEmailNotFound {
		return nil, err
	}

	if email.IsPrimary {
		err = svc.ResetPrimaryCompanyEmail(ctx, companyId)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.GetCompanyEmailStore().CreateCompanyEmail(ctx, companyId, email)
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyEmailStore().GetCompanyEmailById(ctx, companyId, newId)
}

func (svc *service) UpdateCompanyEmail(ctx context.Context, companyId string, id string, email Email) (*Email, error) {
	var err error

	//TODO: Validate companyId
	//TODO: Validate id

	company, err := svc.database.GetCompanyStore().GetCompanyById(ctx, companyId)
	if company == nil && err == nil {
		return nil, ErrCompanyNotFound
	}
	if err != nil {
		return nil, err
	}

	existing, err := svc.database.GetCompanyEmailStore().GetCompanyEmailById(ctx, companyId, id)
	if existing == nil && err == nil {
		return nil, ErrCompanyEmailNotFound
	}
	if err != nil {
		return nil, err
	}

	emailType, err := svc.database.GetEmailTypeStore().GetEmailTypeById(ctx, email.TypeId)
	if emailType == nil && err == nil {
		return nil, ErrEmailTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	duplicate, err := svc.database.GetCompanyEmailStore().GetCompanyEmailByType(ctx, companyId, email.TypeId)
	if duplicate != nil {
		return nil, ErrCompanyEmailDuplicate
	}
	if err != nil && err != ErrCompanyEmailNotFound {
		return nil, err
	}

	if email.IsPrimary {
		err = svc.ResetPrimaryCompanyEmail(ctx, companyId)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.GetCompanyEmailStore().UpdateCompanyEmail(ctx, companyId, id, email)
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyEmailStore().GetCompanyEmailById(ctx, companyId, id)
}

func (svc *service) DeleteCompanyEmail(ctx context.Context, companyId string, id string) error {
	var err error

	//TODO: Validate companyId
	//TODO: Validate id

	company, err := svc.database.GetCompanyStore().GetCompanyById(ctx, companyId)
	if company == nil && err == nil {
		return ErrCompanyNotFound
	}
	if err != nil {
		return err
	}

	existing, err := svc.database.GetCompanyEmailStore().GetCompanyEmailById(ctx, companyId, id)
	if existing == nil && err == nil {
		return ErrCompanyEmailNotFound
	}
	if err != nil {
		return err
	}

	return svc.database.GetCompanyEmailStore().DeleteCompanyEmail(ctx, companyId, id)
}

func (svc *service) ResetPrimaryCompanyEmail(ctx context.Context, companyId string) error {
	var err error

	primary, err := svc.database.GetCompanyEmailStore().GetCompanyPrimaryEmail(ctx, companyId)
	if err != nil && err != ErrCompanyEmailNotFound {
		return err
	}

	if primary != nil {
		primary.IsPrimary = false
		err = svc.database.GetCompanyEmailStore().UpdateCompanyEmail(ctx, companyId, primary.Id, *primary)
		if err != nil {
			return err
		}
	}

	return nil
}

package contact

import (
	"context"
)

func (svc *service) GetCompanyPhones(ctx context.Context, companyId string, pageIndex int, pageSize int) (*PhoneList, error) {
	var err error

	//TODO: Validate companyId

	company, err := svc.database.GetCompanyStore().GetCompanyById(ctx, companyId)
	if company == nil && err == nil {
		return nil, ErrCompanyNotFound
	}
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyPhoneStore().GetCompanyPhones(ctx, companyId, pageIndex, pageSize)
}

func (svc *service) GetCompanyPhoneById(ctx context.Context, companyId string, id string) (*Phone, error) {
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

	return svc.database.GetCompanyPhoneStore().GetCompanyPhoneById(ctx, companyId, id)
}

func (svc *service) CreateCompanyPhone(ctx context.Context, companyId string, phone Phone) (*Phone, error) {
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

	phoneType, err := svc.database.GetPhoneTypeStore().GetPhoneTypeById(ctx, phone.TypeId)
	if phoneType == nil && err == nil {
		return nil, ErrPhoneTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	duplicate, err := svc.database.GetCompanyPhoneStore().GetCompanyPhoneByType(ctx, companyId, phone.TypeId)
	if duplicate != nil {
		return nil, ErrCompanyPhoneDuplicate
	}
	if err != nil {
		return nil, err
	}

	if phone.IsPrimary {
		err = svc.ResetPrimaryCompanyPhone(ctx, companyId)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.GetCompanyPhoneStore().CreateCompanyPhone(ctx, companyId, phone)
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyPhoneStore().GetCompanyPhoneById(ctx, companyId, newId)
}

func (svc *service) UpdateCompanyPhone(ctx context.Context, companyId string, id string, phone Phone) (*Phone, error) {
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

	existing, err := svc.database.GetCompanyPhoneStore().GetCompanyPhoneById(ctx, companyId, id)
	if existing == nil {
		return nil, ErrCompanyPhoneNotFound
	}
	if err != nil {
		return nil, err
	}

	phoneType, err := svc.database.GetPhoneTypeStore().GetPhoneTypeById(ctx, phone.TypeId)
	if phoneType == nil && err == nil {
		return nil, ErrPhoneTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	duplicate, err := svc.database.GetCompanyPhoneStore().GetCompanyPhoneByType(ctx, companyId, phone.TypeId)
	if duplicate != nil {
		return nil, ErrCompanyPhoneDuplicate
	}
	if err != nil {
		return nil, err
	}

	if phone.IsPrimary {
		err = svc.ResetPrimaryCompanyPhone(ctx, companyId)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.GetCompanyPhoneStore().UpdateCompanyPhone(ctx, companyId, id, phone)
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyPhoneStore().GetCompanyPhoneById(ctx, companyId, id)
}

func (svc *service) DeleteCompanyPhone(ctx context.Context, companyId string, id string) error {
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

	existing, err := svc.database.GetCompanyPhoneStore().GetCompanyPhoneById(ctx, companyId, id)
	if existing == nil {
		return ErrCompanyPhoneNotFound
	}
	if err != nil {
		return err
	}

	return svc.database.GetCompanyPhoneStore().DeleteCompanyPhone(ctx, companyId, id)
}

func (svc *service) ResetPrimaryCompanyPhone(ctx context.Context, companyId string) error {
	var err error

	primary, err := svc.database.GetCompanyPhoneStore().GetCompanyPrimaryPhone(ctx, companyId)
	if err != nil {
		return err
	}

	if primary != nil {
		primary.IsPrimary = false
		err = svc.database.GetCompanyPhoneStore().UpdateCompanyPhone(ctx, companyId, primary.Id, *primary)
		if err != nil {
			return err
		}
	}

	return nil
}

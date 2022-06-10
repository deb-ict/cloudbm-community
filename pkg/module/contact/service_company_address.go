package contact

import (
	"context"
)

func (svc *service) GetCompanyAddresses(ctx context.Context, companyId string, pageIndex int, pageSize int) (*AddressList, error) {
	var err error

	//TODO: Validate companyId

	company, err := svc.database.GetCompanyStore().GetCompanyById(ctx, companyId)
	if company == nil && err == nil {
		return nil, ErrCompanyNotFound
	}
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyAddressStore().GetCompanyAddresses(ctx, companyId, pageIndex, pageSize)
}

func (svc *service) GetCompanyAddressById(ctx context.Context, companyId string, id string) (*Address, error) {
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

	return svc.database.GetCompanyAddressStore().GetCompanyAddressById(ctx, companyId, id)
}

func (svc *service) CreateCompanyAddress(ctx context.Context, companyId string, address Address) (*Address, error) {
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

	addressType, err := svc.database.GetAddressTypeStore().GetAddressTypeById(ctx, address.TypeId)
	if addressType == nil && err == nil {
		return nil, ErrAddressTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	duplicate, err := svc.database.GetCompanyAddressStore().GetCompanyAddressByType(ctx, companyId, address.TypeId)
	if duplicate != nil {
		return nil, ErrCompanyAddressDuplicate
	}
	if err != nil {
		return nil, err
	}

	if address.IsPrimary {
		err = svc.ResetPrimaryCompanyAddress(ctx, companyId)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.GetCompanyAddressStore().CreateCompanyAddress(ctx, companyId, address)
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyAddressStore().GetCompanyAddressById(ctx, companyId, newId)
}

func (svc *service) UpdateCompanyAddress(ctx context.Context, companyId string, id string, address Address) (*Address, error) {
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

	existing, err := svc.database.GetCompanyAddressStore().GetCompanyAddressById(ctx, companyId, id)
	if existing == nil {
		return nil, ErrCompanyAddressNotFound
	}

	addressType, err := svc.database.GetAddressTypeStore().GetAddressTypeById(ctx, address.TypeId)
	if addressType == nil && err == nil {
		return nil, ErrAddressTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	duplicate, err := svc.database.GetCompanyAddressStore().GetCompanyAddressByType(ctx, companyId, address.TypeId)
	if duplicate != nil {
		return nil, ErrCompanyAddressDuplicate
	}
	if err != nil {
		return nil, err
	}

	if address.IsPrimary {
		err = svc.ResetPrimaryCompanyAddress(ctx, companyId)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.GetCompanyAddressStore().UpdateCompanyAddress(ctx, companyId, id, address)
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyAddressStore().GetCompanyAddressById(ctx, companyId, id)
}

func (svc *service) DeleteCompanyAddress(ctx context.Context, companyId string, id string) error {
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

	existing, err := svc.database.GetCompanyAddressStore().GetCompanyAddressById(ctx, companyId, id)
	if existing == nil {
		return ErrCompanyAddressNotFound
	}

	return svc.database.GetCompanyAddressStore().DeleteCompanyAddress(ctx, companyId, id)
}

func (svc *service) ResetPrimaryCompanyAddress(ctx context.Context, companyId string) error {
	var err error

	primary, err := svc.database.GetCompanyAddressStore().GetCompanyPrimaryAddress(ctx, companyId)
	if err != nil {
		return err
	}

	if primary != nil {
		primary.IsPrimary = false
		err = svc.database.GetCompanyAddressStore().UpdateCompanyAddress(ctx, companyId, primary.Id, *primary)
		if err != nil {
			return err
		}
	}

	return nil
}

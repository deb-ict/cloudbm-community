package contact

import (
	"context"
)

func (svc *service) GetContactAddresses(ctx context.Context, contactId string, pageIndex int, pageSize int) (*AddressList, error) {
	var err error

	//TODO: Validate contactId

	contact, err := svc.database.GetContactStore().GetContactById(ctx, contactId)
	if contact == nil && err == nil {
		return nil, ErrContactNotFound
	}
	if err != nil {
		return nil, err
	}

	return svc.database.GetContactAddressStore().GetContactAddresses(ctx, contactId, pageIndex, pageSize)
}

func (svc *service) GetContactAddressById(ctx context.Context, contactId string, id string) (*Address, error) {
	var err error

	//TODO: Validate contactId
	//TODO: Validate id

	contact, err := svc.database.GetContactStore().GetContactById(ctx, contactId)
	if contact == nil && err == nil {
		return nil, ErrContactNotFound
	}
	if err != nil {
		return nil, err
	}

	return svc.database.GetContactAddressStore().GetContactAddressById(ctx, contactId, id)
}

func (svc *service) CreateContactAddress(ctx context.Context, contactId string, address Address) (*Address, error) {
	var err error

	//TODO: Validate contactId
	//TODO: Validate id

	contact, err := svc.database.GetContactStore().GetContactById(ctx, contactId)
	if contact == nil && err == nil {
		return nil, ErrContactNotFound
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

	duplicate, err := svc.database.GetContactAddressStore().GetContactAddressByType(ctx, contactId, address.TypeId)
	if duplicate != nil {
		return nil, ErrContactAddressDuplicate
	}
	if err != nil && err != ErrContactAddressNotFound {
		return nil, err
	}

	if address.IsPrimary {
		err = svc.ResetPrimaryContactAddress(ctx, contactId)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.GetContactAddressStore().CreateContactAddress(ctx, contactId, address)
	if err != nil {
		return nil, err
	}

	return svc.database.GetContactAddressStore().GetContactAddressById(ctx, contactId, newId)
}

func (svc *service) UpdateContactAddress(ctx context.Context, contactId string, id string, address Address) (*Address, error) {
	var err error

	//TODO: Validate contactId
	//TODO: Validate id

	contact, err := svc.database.GetContactStore().GetContactById(ctx, contactId)
	if contact == nil && err == nil {
		return nil, ErrContactNotFound
	}
	if err != nil {
		return nil, err
	}

	existing, err := svc.database.GetContactAddressStore().GetContactAddressById(ctx, contactId, id)
	if existing == nil && err == nil {
		return nil, ErrContactAddressNotFound
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

	duplicate, err := svc.database.GetContactAddressStore().GetContactAddressByType(ctx, contactId, address.TypeId)
	if duplicate != nil {
		return nil, ErrContactAddressDuplicate
	}
	if err != nil && err != ErrContactAddressNotFound {
		return nil, err
	}

	if address.IsPrimary {
		err = svc.ResetPrimaryContactAddress(ctx, contactId)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.GetContactAddressStore().UpdateContactAddress(ctx, contactId, id, address)
	if err != nil {
		return nil, err
	}

	return svc.database.GetContactAddressStore().GetContactAddressById(ctx, contactId, id)
}

func (svc *service) DeleteContactAddress(ctx context.Context, contactId string, id string) error {
	var err error

	//TODO: Validate contactId
	//TODO: Validate id

	contact, err := svc.database.GetContactStore().GetContactById(ctx, contactId)
	if contact == nil && err == nil {
		return ErrContactNotFound
	}
	if err != nil {
		return err
	}

	existing, err := svc.database.GetContactAddressStore().GetContactAddressById(ctx, contactId, id)
	if existing == nil && err == nil {
		return ErrContactAddressNotFound
	}
	if err != nil {
		return err
	}

	return svc.database.GetContactAddressStore().DeleteContactAddress(ctx, contactId, id)
}

func (svc *service) ResetPrimaryContactAddress(ctx context.Context, contactId string) error {
	var err error

	primary, err := svc.database.GetContactAddressStore().GetContactPrimaryAddress(ctx, contactId)
	if err != nil && err != ErrContactAddressNotFound {
		return err
	}

	if primary != nil {
		primary.IsPrimary = false
		err = svc.database.GetContactAddressStore().UpdateContactAddress(ctx, contactId, primary.Id, *primary)
		if err != nil {
			return err
		}
	}

	return nil
}

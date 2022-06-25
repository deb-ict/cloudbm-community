package contact

import (
	"context"
)

func (svc *service) GetAddressTypes(ctx context.Context, pageIndex int, pageSize int) (*AddressTypeList, error) {
	return svc.database.GetAddressTypeStore().GetAddressTypes(ctx, pageIndex, pageSize)
}

func (svc *service) GetAddressTypeById(ctx context.Context, id string) (*AddressType, error) {
	//TODO: Validate id

	return svc.database.GetAddressTypeStore().GetAddressTypeById(ctx, id)
}

func (svc *service) CreateAddressType(ctx context.Context, addressType AddressType) (*AddressType, error) {
	var err error

	//TODO: Validate model

	duplicate, err := svc.database.GetAddressTypeStore().GetAddressTypeByName(ctx, addressType.Name)
	if duplicate != nil {
		return nil, ErrAddressTypeDuplicate
	}
	if err != nil {
		return nil, err
	}

	if addressType.IsDefault {
		err = svc.ResetDefaultAddressType(ctx)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.GetAddressTypeStore().CreateAddressType(ctx, addressType)
	if err != nil {
		return nil, err
	}

	return svc.database.GetAddressTypeStore().GetAddressTypeById(ctx, newId)
}

func (svc *service) UpdateAddressType(ctx context.Context, id string, addressType AddressType) (*AddressType, error) {
	var err error

	//TODO: Validate id
	//TODO: Validate model

	existing, err := svc.database.GetAddressTypeStore().GetAddressTypeById(ctx, id)
	if existing == nil && err != nil {
		return nil, ErrAddressTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	if existing.IsSystem {
		return nil, ErrAddressTypeReadOnly
	}

	if addressType.IsDefault {
		err = svc.ResetDefaultAddressType(ctx)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.GetAddressTypeStore().UpdateAddressType(ctx, id, addressType)
	if err != nil {
		return nil, err
	}

	return svc.database.GetAddressTypeStore().GetAddressTypeById(ctx, id)
}

func (svc *service) DeleteAddressType(ctx context.Context, id string) error {
	var err error

	//TODO: Validate id

	existing, err := svc.database.GetAddressTypeStore().GetAddressTypeById(ctx, id)
	if existing == nil && err != nil {
		return ErrAddressTypeNotFound
	}
	if err != nil {
		return err
	}

	if existing.IsSystem {
		return ErrAddressTypeReadOnly
	}

	return svc.database.GetAddressTypeStore().DeleteAddressType(ctx, id)
}

func (svc *service) ResetDefaultAddressType(ctx context.Context) error {
	var err error

	defaultType, err := svc.database.GetAddressTypeStore().GetDefaultAddressType(ctx)
	if err != nil {
		return err
	}

	if defaultType != nil {
		defaultType.IsDefault = false
		err = svc.database.GetAddressTypeStore().UpdateAddressType(ctx, defaultType.Id, *defaultType)
		if err != nil {
			return err
		}
	}

	return nil
}

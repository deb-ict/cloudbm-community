package contact

import (
	"context"
)

func (svc *service) GetPhoneTypes(ctx context.Context, pageIndex int, pageSize int) (*PhoneTypeList, error) {
	return svc.database.GetPhoneTypeStore().GetPhoneTypes(ctx, pageIndex, pageSize)
}

func (svc *service) GetPhoneTypeById(ctx context.Context, id string) (*PhoneType, error) {
	//TODO: Validate id

	return svc.database.GetPhoneTypeStore().GetPhoneTypeById(ctx, id)
}

func (svc *service) CreatePhoneType(ctx context.Context, phoneType PhoneType) (*PhoneType, error) {
	var err error

	//TODO: Validate model

	duplicate, err := svc.database.GetPhoneTypeStore().GetPhoneTypeByName(ctx, phoneType.Name)
	if duplicate != nil {
		return nil, ErrPhoneTypeDuplicate
	}
	if err != nil && err != ErrPhoneTypeNotFound {
		return nil, err
	}

	if phoneType.IsDefault {
		err = svc.ResetDefaultPhoneType(ctx)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.GetPhoneTypeStore().CreatePhoneType(ctx, phoneType)
	if err != nil {
		return nil, err
	}

	return svc.database.GetPhoneTypeStore().GetPhoneTypeById(ctx, newId)
}

func (svc *service) UpdatePhoneType(ctx context.Context, id string, phoneType PhoneType) (*PhoneType, error) {
	var err error

	//TODO: Validate id
	//TODO: Validate model

	existing, err := svc.database.GetPhoneTypeStore().GetPhoneTypeById(ctx, id)
	if existing == nil && err == nil {
		return nil, ErrPhoneTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	if existing.IsSystem {
		return nil, ErrPhoneTypeReadOnly
	}

	if phoneType.IsDefault && !existing.IsDefault {
		err = svc.ResetDefaultPhoneType(ctx)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.GetPhoneTypeStore().UpdatePhoneType(ctx, id, phoneType)
	if err != nil {
		return nil, err
	}

	return svc.database.GetPhoneTypeStore().GetPhoneTypeById(ctx, id)
}

func (svc *service) DeletePhoneType(ctx context.Context, id string) error {
	var err error

	//TODO: Validate id

	existing, err := svc.database.GetPhoneTypeStore().GetPhoneTypeById(ctx, id)
	if existing == nil && err == nil {
		return ErrPhoneTypeNotFound
	}
	if err != nil {
		return err
	}

	if existing.IsSystem {
		return ErrPhoneTypeReadOnly
	}

	return svc.database.GetPhoneTypeStore().DeletePhoneType(ctx, id)
}

func (svc *service) SetDefaultPhoneType(ctx context.Context, id string) error {
	var err error

	//TODO: Validate id
	//TODO: Validate model

	existing, err := svc.database.GetPhoneTypeStore().GetPhoneTypeById(ctx, id)
	if existing == nil && err == nil {
		return ErrPhoneTypeNotFound
	}
	if err != nil {
		return err
	}

	if !existing.IsDefault {
		err = svc.ResetDefaultPhoneType(ctx)
		if err != nil {
			return err
		}

		existing.IsDefault = true
		err = svc.database.GetPhoneTypeStore().UpdatePhoneType(ctx, id, *existing)
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *service) ResetDefaultPhoneType(ctx context.Context) error {
	var err error

	defaultType, err := svc.database.GetPhoneTypeStore().GetDefaultPhoneType(ctx)
	if err != nil && err != ErrPhoneTypeNotFound {
		return err
	}

	if defaultType != nil {
		defaultType.IsDefault = false
		err = svc.database.GetPhoneTypeStore().UpdatePhoneType(ctx, defaultType.Id, *defaultType)
		if err != nil {
			return err
		}
	}

	return nil
}

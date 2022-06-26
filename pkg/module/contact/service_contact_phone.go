package contact

import (
	"context"
)

func (svc *service) GetContactPhones(ctx context.Context, contactId string, pageIndex int, pageSize int) (*PhoneList, error) {
	var err error

	//TODO: Validate contactId

	contact, err := svc.database.GetContactStore().GetContactById(ctx, contactId)
	if contact == nil && err == nil {
		return nil, ErrContactNotFound
	}
	if err != nil {
		return nil, err
	}

	return svc.database.GetContactPhoneStore().GetContactPhones(ctx, contactId, pageIndex, pageSize)
}

func (svc *service) GetContactPhoneById(ctx context.Context, contactId string, id string) (*Phone, error) {
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

	return svc.database.GetContactPhoneStore().GetContactPhoneById(ctx, contactId, id)
}

func (svc *service) CreateContactPhone(ctx context.Context, contactId string, phone Phone) (*Phone, error) {
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

	phoneType, err := svc.database.GetPhoneTypeStore().GetPhoneTypeById(ctx, phone.TypeId)
	if phoneType == nil && err == nil {
		return nil, ErrPhoneTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	duplicate, err := svc.database.GetContactPhoneStore().GetContactPhoneByType(ctx, contactId, phone.TypeId)
	if duplicate != nil {
		return nil, ErrContactPhoneDuplicate
	}
	if err != nil && err != ErrContactPhoneNotFound {
		return nil, err
	}

	if phone.IsPrimary {
		err = svc.ResetPrimaryContactPhone(ctx, contactId)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.GetContactPhoneStore().CreateContactPhone(ctx, contactId, phone)
	if err != nil {
		return nil, err
	}

	return svc.database.GetContactPhoneStore().GetContactPhoneById(ctx, contactId, newId)
}

func (svc *service) UpdateContactPhone(ctx context.Context, contactId string, id string, phone Phone) (*Phone, error) {
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

	existing, err := svc.database.GetContactPhoneStore().GetContactPhoneById(ctx, contactId, id)
	if existing == nil && err == nil {
		return nil, ErrContactPhoneNotFound
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

	duplicate, err := svc.database.GetContactPhoneStore().GetContactPhoneByType(ctx, contactId, phone.TypeId)
	if duplicate != nil {
		return nil, ErrContactPhoneDuplicate
	}
	if err != nil && err != ErrContactPhoneNotFound {
		return nil, err
	}

	if phone.IsPrimary {
		err = svc.ResetPrimaryContactPhone(ctx, contactId)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.GetContactPhoneStore().UpdateContactPhone(ctx, contactId, id, phone)
	if err != nil {
		return nil, err
	}

	return svc.database.GetContactPhoneStore().GetContactPhoneById(ctx, contactId, id)
}

func (svc *service) DeleteContactPhone(ctx context.Context, contactId string, id string) error {
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

	existing, err := svc.database.GetContactPhoneStore().GetContactPhoneById(ctx, contactId, id)
	if existing == nil && err == nil {
		return ErrContactPhoneNotFound
	}
	if err != nil {
		return err
	}

	return svc.database.GetContactPhoneStore().DeleteContactPhone(ctx, contactId, id)
}

func (svc *service) ResetPrimaryContactPhone(ctx context.Context, contactId string) error {
	var err error

	primary, err := svc.database.GetContactPhoneStore().GetContactPrimaryPhone(ctx, contactId)
	if err != nil && err != ErrContactPhoneNotFound {
		return err
	}

	if primary != nil {
		primary.IsPrimary = false
		err = svc.database.GetContactPhoneStore().UpdateContactPhone(ctx, contactId, primary.Id, *primary)
		if err != nil {
			return err
		}
	}

	return nil
}

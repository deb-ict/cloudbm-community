package contact

import (
	"context"
)

func (svc *service) GetContacts(ctx context.Context, pageIndex int, pageSize int) (*ContactList, error) {
	return svc.database.GetContactStore().GetContacts(ctx, pageIndex, pageSize)
}

func (svc *service) GetContactById(ctx context.Context, id string) (*Contact, error) {
	//TODO: Validate id

	return svc.database.GetContactStore().GetContactById(ctx, id)
}

func (svc *service) CreateContact(ctx context.Context, contact Contact) (*Contact, error) {
	var err error

	//TODO: Validate model

	duplicate, err := svc.database.GetContactStore().GetContactByName(ctx, contact.FamilyName, contact.GivenName)
	if duplicate != nil {
		return nil, ErrContactDuplicateName
	}
	if err != nil && err != ErrContactNotFound {
		return nil, err
	}

	newId, err := svc.database.GetContactStore().CreateContact(ctx, contact)
	if err != nil {
		return nil, err
	}

	return svc.database.GetContactStore().GetContactById(ctx, newId)
}

func (svc *service) UpdateContact(ctx context.Context, id string, contact Contact) (*Contact, error) {
	var err error

	//TODO: Validate id
	//TODO: Validate model

	existing, err := svc.database.GetContactStore().GetContactById(ctx, id)
	if existing == nil && err == nil {
		return nil, ErrContactNotFound
	}
	if err != nil {
		return nil, err
	}

	err = svc.database.GetContactStore().UpdateContact(ctx, id, contact)
	if err != nil {
		return nil, err
	}

	return svc.database.GetContactStore().GetContactById(ctx, id)
}

func (svc *service) DeleteContact(ctx context.Context, id string) error {
	var err error

	//TODO: Validate id

	existing, err := svc.database.GetContactStore().GetContactById(ctx, id)
	if existing == nil && err == nil {
		return ErrContactNotFound
	}
	if err != nil {
		return err
	}

	return svc.database.GetContactStore().DeleteContact(ctx, id)
}

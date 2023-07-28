package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetContactAddresses(ctx context.Context, contactId string, offset int64, limit int64, filter *model.AddressFilter, sort *core.Sort) ([]*model.Address, int64, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, 0, err
	}
	if parent == nil {
		return nil, 0, contact.ErrContactNotFound
	}

	data, count, err := svc.database.ContactAddressRepository().GetContactAddresses(ctx, parent, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetContactAddressById(ctx context.Context, contactId string, id string) (*model.Address, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	data, err := svc.database.ContactAddressRepository().GetContactAddressById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyAddressNotFound
	}

	return data, nil
}

func (svc *service) CreateContactAddress(ctx context.Context, contactId string, model *model.Address) (*model.Address, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	newId, err := svc.database.ContactAddressRepository().CreateContactAddress(ctx, parent, model)
	if err != nil {
		return nil, err
	}
	return svc.GetContactAddressById(ctx, contactId, newId)
}

func (svc *service) UpdateContactAddress(ctx context.Context, contactId string, id string, model *model.Address) (*model.Address, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	data, err := svc.database.ContactAddressRepository().GetContactAddressById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrContactAddressNotFound
	}

	err = svc.database.ContactAddressRepository().UpdateContactAddress(ctx, parent, data)
	if err != nil {
		return nil, err
	}
	return svc.GetContactAddressById(ctx, contactId, id)
}

func (svc *service) DeleteContactAddress(ctx context.Context, contactId string, id string) error {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return err
	}
	if parent == nil {
		return contact.ErrContactNotFound
	}

	data, err := svc.database.ContactAddressRepository().GetContactAddressById(ctx, parent, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrContactAddressNotFound
	}

	err = svc.database.ContactAddressRepository().DeleteContactAddress(ctx, parent, data)
	if err != nil {
		return err
	}
	return nil
}

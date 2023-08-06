package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetContactAddresses(ctx context.Context, contactId string, offset int64, limit int64, filter *model.AddressFilter, sort *core.Sort) ([]*model.Address, int64, error) {
	parent, err := svc.database.Contacts().GetContactById(ctx, contactId)
	if err != nil {
		return nil, 0, err
	}
	if parent == nil {
		return nil, 0, contact.ErrContactNotFound
	}

	data, count, err := svc.database.ContactAddresses().GetContactAddresses(ctx, parent, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetContactAddressById(ctx context.Context, contactId string, id string) (*model.Address, error) {
	parent, err := svc.database.Contacts().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	data, err := svc.database.ContactAddresses().GetContactAddressById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyAddressNotFound
	}

	return data, nil
}

func (svc *service) CreateContactAddress(ctx context.Context, contactId string, model *model.Address) (*model.Address, error) {
	parent, err := svc.database.Contacts().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	err = svc.validateContactAddress(ctx, parent, model)
	if err != nil {
		return nil, err
	}

	if model.IsDefault {
		err = svc.resetDefaultContactAddress(ctx, parent, model)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.ContactAddresses().CreateContactAddress(ctx, parent, model)
	if err != nil {
		return nil, err
	}
	return svc.GetContactAddressById(ctx, contactId, newId)
}

func (svc *service) UpdateContactAddress(ctx context.Context, contactId string, id string, model *model.Address) (*model.Address, error) {
	parent, err := svc.database.Contacts().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	data, err := svc.database.ContactAddresses().GetContactAddressById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrContactAddressNotFound
	}

	data.Type = model.Type
	data.Street = model.Street
	data.StreetNr = model.StreetNr
	data.Unit = model.Unit
	data.PostalCode = model.PostalCode
	data.City = model.City
	data.State = model.State
	data.Country = model.Country
	data.IsDefault = model.IsDefault

	err = svc.validateContactAddress(ctx, parent, data)
	if err != nil {
		return nil, err
	}

	if data.IsDefault {
		err = svc.resetDefaultContactAddress(ctx, parent, data)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.ContactAddresses().UpdateContactAddress(ctx, parent, data)
	if err != nil {
		return nil, err
	}
	return svc.GetContactAddressById(ctx, contactId, id)
}

func (svc *service) DeleteContactAddress(ctx context.Context, contactId string, id string) error {
	parent, err := svc.database.Contacts().GetContactById(ctx, contactId)
	if err != nil {
		return err
	}
	if parent == nil {
		return contact.ErrContactNotFound
	}

	data, err := svc.database.ContactAddresses().GetContactAddressById(ctx, parent, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrContactAddressNotFound
	}
	if data.IsDefault {
		return contact.ErrContactAddressIsDefault
	}

	err = svc.database.ContactAddresses().DeleteContactAddress(ctx, parent, data)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) resetDefaultContactAddress(ctx context.Context, parent *model.Contact, model *model.Address) error {
	current, err := svc.database.ContactAddresses().GetDefaultContactAddress(ctx, parent)
	if err != nil {
		return err
	}
	if current != nil && current.Id != model.Id {
		current.IsDefault = false
		err = svc.database.ContactAddresses().UpdateContactAddress(ctx, parent, current)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) validateContactAddress(ctx context.Context, parent *model.Contact, model *model.Address) error {
	modelType, err := svc.database.AddressTypes().GetAddressTypeById(ctx, model.Type.Id)
	if err != nil {
		return err
	}
	if modelType == nil {
		return contact.ErrAddressTypeNotFound
	}

	return nil
}

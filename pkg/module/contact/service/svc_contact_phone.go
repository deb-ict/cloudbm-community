package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetContactPhones(ctx context.Context, contactId string, offset int64, limit int64, filter *model.PhoneFilter, sort *core.Sort) ([]*model.Phone, int64, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, 0, err
	}
	if parent == nil {
		return nil, 0, contact.ErrContactNotFound
	}

	data, count, err := svc.database.ContactPhoneRepository().GetContactPhones(ctx, parent, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetContactPhoneById(ctx context.Context, contactId string, id string) (*model.Phone, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	data, err := svc.database.ContactPhoneRepository().GetContactPhoneById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyAddressNotFound
	}

	return data, nil
}

func (svc *service) CreateContactPhone(ctx context.Context, contactId string, model *model.Phone) (*model.Phone, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	err = svc.validateContactPhone(ctx, parent, model)
	if err != nil {
		return nil, err
	}

	if model.IsDefault {
		err = svc.resetDefaultContactPhone(ctx, parent, model)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.ContactPhoneRepository().CreateContactPhone(ctx, parent, model)
	if err != nil {
		return nil, err
	}
	return svc.GetContactPhoneById(ctx, contactId, newId)
}

func (svc *service) UpdateContactPhone(ctx context.Context, contactId string, id string, model *model.Phone) (*model.Phone, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	data, err := svc.database.ContactPhoneRepository().GetContactPhoneById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrContactPhoneNotFound
	}

	data.Type = model.Type
	data.PhoneNumber = model.PhoneNumber
	data.Extension = model.Extension
	data.IsDefault = model.IsDefault

	err = svc.validateContactPhone(ctx, parent, data)
	if err != nil {
		return nil, err
	}

	if data.IsDefault {
		err = svc.resetDefaultContactPhone(ctx, parent, data)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.ContactPhoneRepository().UpdateContactPhone(ctx, parent, data)
	if err != nil {
		return nil, err
	}
	return svc.GetContactPhoneById(ctx, contactId, id)
}

func (svc *service) DeleteContactPhone(ctx context.Context, contactId string, id string) error {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return err
	}
	if parent == nil {
		return contact.ErrContactNotFound
	}

	data, err := svc.database.ContactPhoneRepository().GetContactPhoneById(ctx, parent, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrContactPhoneNotFound
	}
	if data.IsDefault {
		return contact.ErrContactPhoneIsDefault
	}

	err = svc.database.ContactPhoneRepository().DeleteContactPhone(ctx, parent, data)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) resetDefaultContactPhone(ctx context.Context, parent *model.Contact, model *model.Phone) error {
	current, err := svc.database.ContactPhoneRepository().GetDefaultContactPhone(ctx, parent)
	if err != nil {
		return err
	}
	if current != nil && current.Id != model.Id {
		current.IsDefault = false
		err = svc.database.ContactPhoneRepository().UpdateContactPhone(ctx, parent, current)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) validateContactPhone(ctx context.Context, parent *model.Contact, model *model.Phone) error {
	modelType, err := svc.database.PhoneTypeRepository().GetPhoneTypeById(ctx, model.Type.Id)
	if err != nil {
		return err
	}
	if modelType == nil {
		return contact.ErrPhoneTypeNotFound
	}

	return nil
}

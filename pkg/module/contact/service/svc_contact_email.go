package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetContactEmails(ctx context.Context, contactId string, offset int64, limit int64, filter *model.EmailFilter, sort *core.Sort) ([]*model.Email, int64, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, 0, err
	}
	if parent == nil {
		return nil, 0, contact.ErrContactNotFound
	}

	data, count, err := svc.database.ContactEmailRepository().GetContactEmails(ctx, parent, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetContactEmailById(ctx context.Context, contactId string, id string) (*model.Email, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	data, err := svc.database.ContactEmailRepository().GetContactEmailById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyAddressNotFound
	}

	return data, nil
}

func (svc *service) CreateContactEmail(ctx context.Context, contactId string, model *model.Email) (*model.Email, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	err = svc.validateContactEmail(ctx, parent, model)
	if err != nil {
		return nil, err
	}

	if model.IsDefault {
		err = svc.resetDefaultContactEmail(ctx, parent, model)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.ContactEmailRepository().CreateContactEmail(ctx, parent, model)
	if err != nil {
		return nil, err
	}
	return svc.GetContactEmailById(ctx, contactId, newId)
}

func (svc *service) UpdateContactEmail(ctx context.Context, contactId string, id string, model *model.Email) (*model.Email, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	data, err := svc.database.ContactEmailRepository().GetContactEmailById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrContactEmailNotFound
	}

	data.Type = model.Type
	data.Email = model.Email
	data.IsDefault = model.IsDefault

	err = svc.validateContactEmail(ctx, parent, data)
	if err != nil {
		return nil, err
	}

	if data.IsDefault {
		err = svc.resetDefaultContactEmail(ctx, parent, data)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.ContactEmailRepository().UpdateContactEmail(ctx, parent, data)
	if err != nil {
		return nil, err
	}
	return svc.GetContactEmailById(ctx, contactId, id)
}

func (svc *service) DeleteContactEmail(ctx context.Context, contactId string, id string) error {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return err
	}
	if parent == nil {
		return contact.ErrContactNotFound
	}

	data, err := svc.database.ContactEmailRepository().GetContactEmailById(ctx, parent, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrContactEmailNotFound
	}
	if data.IsDefault {
		return contact.ErrContactEmailIsDefault
	}

	err = svc.database.ContactEmailRepository().DeleteContactEmail(ctx, parent, data)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) resetDefaultContactEmail(ctx context.Context, parent *model.Contact, model *model.Email) error {
	current, err := svc.database.ContactEmailRepository().GetDefaultContactEmail(ctx, parent)
	if err != nil {
		return err
	}
	if current != nil && current.Id != model.Id {
		current.IsDefault = false
		err = svc.database.ContactEmailRepository().UpdateContactEmail(ctx, parent, current)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) validateContactEmail(ctx context.Context, parent *model.Contact, model *model.Email) error {
	modelType, err := svc.database.EmailTypeRepository().GetEmailTypeById(ctx, model.Type.Id)
	if err != nil {
		return err
	}
	if modelType == nil {
		return contact.ErrEmailTypeNotFound
	}

	return nil
}

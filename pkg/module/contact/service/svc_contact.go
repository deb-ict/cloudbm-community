package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetContacts(ctx context.Context, offset int64, limit int64, filter *model.ContactFilter, sort *core.Sort) ([]*model.Contact, int64, error) {
	data, count, err := svc.database.Contacts().GetContacts(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetContactById(ctx context.Context, id string) (*model.Contact, error) {
	data, err := svc.database.Contacts().GetContactById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrContactNotFound
	}

	return data, nil
}

func (svc *service) CreateContact(ctx context.Context, model *model.Contact) (*model.Contact, error) {
	model.Id = ""

	err := svc.validateContact(ctx, model)
	if err != nil {
		return nil, err
	}

	newId, err := svc.database.Contacts().CreateContact(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetContactById(ctx, newId)
}

func (svc *service) UpdateContact(ctx context.Context, id string, model *model.Contact) (*model.Contact, error) {
	model.Id = id

	data, err := svc.GetContactById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrContactNotFound
	}
	data.UpdateModel(model)

	err = svc.validateContact(ctx, data)
	if err != nil {
		return nil, err
	}

	err = svc.database.Contacts().UpdateContact(ctx, data)
	if err != nil {
		return nil, err
	}
	return svc.GetContactById(ctx, id)
}

func (svc *service) DeleteContact(ctx context.Context, id string) error {
	data, err := svc.GetContactById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrContactNotFound
	}
	if data.IsSystem {
		return contact.ErrContactReadOnly
	}

	err = svc.database.Contacts().DeleteContact(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) validateContact(ctx context.Context, model *model.Contact) error {
	if model.Title != nil {
		contactTitle, err := svc.database.ContactTitles().GetContactTitleById(ctx, model.Title.Id)
		if err != nil {
			return err
		}
		if contactTitle == nil {
			return contact.ErrContactTitleNotFound
		}
	}

	for _, address := range model.Addresses {
		err := svc.validateContactAddress(ctx, model, address)
		if err != nil {
			return err
		}
	}
	for _, email := range model.Emails {
		err := svc.validateContactEmail(ctx, model, email)
		if err != nil {
			return err
		}
	}
	for _, phone := range model.Phones {
		err := svc.validateContactPhone(ctx, model, phone)
		if err != nil {
			return err
		}
	}
	for _, uri := range model.Uris {
		err := svc.validateContactUri(ctx, model, uri)
		if err != nil {
			return err
		}
	}

	return nil
}

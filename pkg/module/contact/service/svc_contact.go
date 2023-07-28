package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetContacts(ctx context.Context, offset int64, limit int64, filter *model.ContactFilter, sort *core.Sort) ([]*model.Contact, int64, error) {
	data, count, err := svc.database.ContactRepository().GetContacts(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetContactById(ctx context.Context, id string) (*model.Contact, error) {
	data, err := svc.database.ContactRepository().GetContactById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrContactNotFound
	}

	return data, nil
}

func (svc *service) CreateContact(ctx context.Context, model *model.Contact) (*model.Contact, error) {
	newId, err := svc.database.ContactRepository().CreateContact(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetContactById(ctx, newId)
}

func (svc *service) UpdateContact(ctx context.Context, id string, model *model.Contact) (*model.Contact, error) {
	data, err := svc.GetContactById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return contact.ErrXXXNotFound
	}

	//TODO: Set fields

	err = svc.database.ContactRepository().UpdateContact(ctx, data)
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
		return contact.ErrXXXNotFound
	}

	err = svc.database.ContactRepository().DeleteContact(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

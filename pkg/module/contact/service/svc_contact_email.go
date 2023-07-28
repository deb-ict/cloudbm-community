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
}

func (svc *service) GetContactEmailById(ctx context.Context, contactId string, id string) (*model.Email, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}
}

func (svc *service) CreateContactEmail(ctx context.Context, contactId string, model *model.Email) (*model.Email, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}
}

func (svc *service) UpdateContactEmail(ctx context.Context, contactId string, id string, model *model.Email) (*model.Email, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}
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
		return contact.ErrContactXXXNotFound
	}

	err = svc.database.ContactEmailRepository().DeleteContactEmail(ctx, parent, data)
	if err != nil {
		return err
	}
	return nil
}

package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetContactUris(ctx context.Context, contactId string, offset int64, limit int64, filter *model.UriFilter, sort *core.Sort) ([]*model.Uri, int64, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, 0, err
	}
	if parent == nil {
		return nil, 0, contact.ErrContactNotFound
	}

	data, count, err := svc.database.ContactUriRepository().GetContactUris(ctx, parent, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetContactUriById(ctx context.Context, contactId string, id string) (*model.Uri, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	data, err := svc.database.ContactUriRepository().GetContactUriById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyAddressNotFound
	}

	return data, nil
}

func (svc *service) CreateContactUri(ctx context.Context, contactId string, model *model.Uri) (*model.Uri, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	newId, err := svc.database.ContactUriRepository().CreateContactUri(ctx, parent, model)
	if err != nil {
		return nil, err
	}
	return svc.GetContactUriById(ctx, contactId, newId)
}

func (svc *service) UpdateContactUri(ctx context.Context, contactId string, id string, model *model.Uri) (*model.Uri, error) {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	data, err := svc.database.ContactUriRepository().GetContactUriById(ctx, parent, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrContactUriNotFound
	}

	err = svc.database.ContactUriRepository().UpdateContactUri(ctx, parent, data)
	if err != nil {
		return nil, err
	}
	return svc.GetContactUriById(ctx, contactId, id)
}

func (svc *service) DeleteContactUri(ctx context.Context, contactId string, id string) error {
	parent, err := svc.database.ContactRepository().GetContactById(ctx, contactId)
	if err != nil {
		return err
	}
	if parent == nil {
		return contact.ErrContactNotFound
	}

	data, err := svc.database.ContactUriRepository().GetContactUriById(ctx, parent, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrContactUriNotFound
	}

	err = svc.database.ContactUriRepository().DeleteContactUri(ctx, parent, data)
	if err != nil {
		return err
	}
	return nil
}

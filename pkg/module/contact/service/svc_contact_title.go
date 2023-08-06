package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetContactTitles(ctx context.Context, offset int64, limit int64, filter *model.ContactTitleFilter, sort *core.Sort) ([]*model.ContactTitle, int64, error) {
	data, count, err := svc.database.ContactTitleRepository().GetContactTitles(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetContactTitleById(ctx context.Context, id string) (*model.ContactTitle, error) {
	data, err := svc.database.ContactTitleRepository().GetContactTitleById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrContactTitleNotFound
	}

	return data, nil
}

func (svc *service) CreateContactTitle(ctx context.Context, model *model.ContactTitle) (*model.ContactTitle, error) {
	model.Key = svc.StringNormalizer().NormalizeString(model.Key)
	err := svc.validateContactTitleName(ctx, model)
	if err != nil {
		return nil, err
	}

	newId, err := svc.database.ContactTitleRepository().CreateContactTitle(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetContactTitleById(ctx, newId)
}

func (svc *service) UpdateContactTitle(ctx context.Context, id string, model *model.ContactTitle) (*model.ContactTitle, error) {
	data, err := svc.GetContactTitleById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrContactTitleNotFound
	}

	data.Translations = model.Translations

	err = svc.validateContactTitleName(ctx, data)
	if err != nil {
		return nil, err
	}

	err = svc.database.ContactTitleRepository().UpdateContactTitle(ctx, data)
	if err != nil {
		return nil, err
	}
	return svc.GetContactTitleById(ctx, id)
}

func (svc *service) DeleteContactTitle(ctx context.Context, id string) error {
	data, err := svc.GetContactTitleById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrContactTitleNotFound
	}
	if data.IsSystem {
		return contact.ErrContactTitleReadOnly
	}

	err = svc.database.ContactTitleRepository().DeleteContactTitle(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) validateContactTitleName(ctx context.Context, model *model.ContactTitle) error {
	if model.IsTransient() {
		existing, err := svc.database.ContactTitleRepository().GetContactTitleByKey(ctx, model.Key)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrContactTitleDuplicateKey
		}
	}

	for _, translation := range model.Translations {
		existing, err := svc.database.ContactTitleRepository().GetContactTitleByName(ctx, translation.Language, translation.Name)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrContactTitleDuplicateName
		}
	}
	return nil
}

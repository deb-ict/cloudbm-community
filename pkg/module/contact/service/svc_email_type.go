package service

import (
	"context"
	"strings"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetEmailTypes(ctx context.Context, offset int64, limit int64, filter *model.EmailTypeFilter, sort *core.Sort) ([]*model.EmailType, int64, error) {
	data, count, err := svc.database.EmailTypeRepository().GetEmailTypes(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetEmailTypeById(ctx context.Context, id string) (*model.EmailType, error) {
	data, err := svc.database.EmailTypeRepository().GetEmailTypeById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrEmailTypeNotFound
	}

	return data, nil
}

func (svc *service) CreateEmailType(ctx context.Context, model *model.EmailType) (*model.EmailType, error) {
	model.Key = strings.ToLower(model.Key)
	existing, err := svc.database.EmailTypeRepository().GetEmailTypeByKey(ctx, model.Key)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, contact.ErrEmailTypeDuplicateKey
	}

	//TODO: Check for duplicates on name

	newId, err := svc.database.EmailTypeRepository().CreateEmailType(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetEmailTypeById(ctx, newId)
}

func (svc *service) UpdateEmailType(ctx context.Context, id string, model *model.EmailType) (*model.EmailType, error) {
	data, err := svc.GetEmailTypeById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrEmailTypeNotFound
	}

	//TODO: Check for duplicates on name

	//TODO: Set fields

	err = svc.database.EmailTypeRepository().UpdateEmailType(ctx, data)
	if err != nil {
		return nil, err
	}
	return svc.GetEmailTypeById(ctx, id)
}

func (svc *service) DeleteEmailType(ctx context.Context, id string) error {
	data, err := svc.GetEmailTypeById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrEmailTypeNotFound
	}

	//TODO: Check dependencies

	err = svc.database.EmailTypeRepository().DeleteEmailType(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

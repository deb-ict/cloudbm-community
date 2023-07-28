package service

import (
	"context"
	"strings"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetUriTypes(ctx context.Context, offset int64, limit int64, filter *model.UriTypeFilter, sort *core.Sort) ([]*model.UriType, int64, error) {
	data, count, err := svc.database.UriTypeRepository().GetUriTypes(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetUriTypeById(ctx context.Context, id string) (*model.UriType, error) {
	data, err := svc.database.UriTypeRepository().GetUriTypeById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrUriTypeNotFound
	}

	return data, nil
}

func (svc *service) CreateUriType(ctx context.Context, model *model.UriType) (*model.UriType, error) {
	model.Key = strings.ToLower(model.Key)
	existing, err := svc.database.UriTypeRepository().GetUriTypeByKey(ctx, model.Key)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, contact.ErrUriTypeDuplicateKey
	}

	//TODO: Check for duplicates on name

	newId, err := svc.database.UriTypeRepository().CreateUriType(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetUriTypeById(ctx, newId)
}

func (svc *service) UpdateUriType(ctx context.Context, id string, model *model.UriType) (*model.UriType, error) {
	data, err := svc.GetUriTypeById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrUriTypeNotFound
	}

	//TODO: Check for duplicates on name

	//TODO: Set fields

	err = svc.database.UriTypeRepository().UpdateUriType(ctx, data)
	if err != nil {
		return nil, err
	}
	return svc.GetUriTypeById(ctx, id)
}

func (svc *service) DeleteUriType(ctx context.Context, id string) error {
	data, err := svc.GetUriTypeById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrUriTypeNotFound
	}

	//TODO: Check dependencies

	err = svc.database.UriTypeRepository().DeleteUriType(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

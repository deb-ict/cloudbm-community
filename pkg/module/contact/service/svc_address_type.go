package service

import (
	"context"
	"strings"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetAddressTypes(ctx context.Context, offset int64, limit int64, filter *model.AddressTypeFilter, sort *core.Sort) ([]*model.AddressType, int64, error) {
	data, count, err := svc.database.AddressTypeRepository().GetAddressTypes(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetAddressTypeById(ctx context.Context, id string) (*model.AddressType, error) {
	data, err := svc.database.AddressTypeRepository().GetAddressTypeById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrAddressTypeNotFound
	}

	return data, nil
}

func (svc *service) CreateAddressType(ctx context.Context, model *model.AddressType) (*model.AddressType, error) {
	model.Key = strings.ToLower(model.Key)
	existing, err := svc.database.AddressTypeRepository().GetAddressTypeByKey(ctx, model.Key)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, contact.ErrAddressTypeDuplicateKey
	}

	newId, err := svc.database.AddressTypeRepository().CreateAddressType(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetAddressTypeById(ctx, newId)
}

func (svc *service) UpdateAddressType(ctx context.Context, id string, model *model.AddressType) (*model.AddressType, error) {
	data, err := svc.database.AddressTypeRepository().GetAddressTypeById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrAddressTypeNotFound
	}

	//TODO: Set fields

	err = svc.database.AddressTypeRepository().UpdateAddressType(ctx, data)
	if err != nil {
		return nil, err
	}
	return svc.GetAddressTypeById(ctx, id)
}

func (svc *service) DeleteAddressType(ctx context.Context, id string) error {
	data, err := svc.database.AddressTypeRepository().GetAddressTypeById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrAddressTypeNotFound
	}

	//TODO: Check dependencies

	err = svc.database.AddressTypeRepository().DeleteAddressType(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

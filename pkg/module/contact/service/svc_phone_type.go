package service

import (
	"context"
	"strings"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetPhoneTypes(ctx context.Context, offset int64, limit int64, filter *model.PhoneTypeFilter, sort *core.Sort) ([]*model.PhoneType, int64, error) {
	data, count, err := svc.database.PhoneTypeRepository().GetPhoneTypes(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetPhoneTypeById(ctx context.Context, id string) (*model.PhoneType, error) {
	data, err := svc.database.PhoneTypeRepository().GetPhoneTypeById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrPhoneTypeNotFound
	}

	return data, nil
}

func (svc *service) CreatePhoneType(ctx context.Context, model *model.PhoneType) (*model.PhoneType, error) {
	model.Key = strings.ToLower(model.Key)
	existing, err := svc.database.PhoneTypeRepository().GetPhoneTypeByKey(ctx, model.Key)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, contact.ErrPhoneTypeDuplicateKey
	}

	newId, err := svc.database.PhoneTypeRepository().CreatePhoneType(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetPhoneTypeById(ctx, newId)
}

func (svc *service) UpdatePhoneType(ctx context.Context, id string, model *model.PhoneType) (*model.PhoneType, error) {
	data, err := svc.GetPhoneTypeById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return contact.ErrXXXNotFound
	}

	//TODO: Set fields

	err = svc.database.PhoneTypeRepository().UpdatePhoneType(ctx, data)
	if err != nil {
		return nil, err
	}
	return svc.GetPhoneTypeById(ctx, id)
}

func (svc *service) DeletePhoneType(ctx context.Context, id string) error {
	data, err := svc.GetPhoneTypeById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrXXXNotFound
	}

	//TODO: Check dependencies
	count, err := 0, nil
	if err != nil {
		return err
	}
	if count > 0 {
		return contact.ErrPhoneTypeInUse
	}

	err = svc.database.PhoneTypeRepository().DeletePhoneType(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

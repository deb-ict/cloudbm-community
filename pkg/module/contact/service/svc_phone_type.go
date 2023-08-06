package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetPhoneTypes(ctx context.Context, offset int64, limit int64, filter *model.PhoneTypeFilter, sort *core.Sort) ([]*model.PhoneType, int64, error) {
	data, count, err := svc.database.PhoneTypes().GetPhoneTypes(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetPhoneTypeById(ctx context.Context, id string) (*model.PhoneType, error) {
	data, err := svc.database.PhoneTypes().GetPhoneTypeById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrPhoneTypeNotFound
	}

	return data, nil
}

func (svc *service) CreatePhoneType(ctx context.Context, model *model.PhoneType) (*model.PhoneType, error) {
	model.Key = svc.StringNormalizer().NormalizeString(model.Key)
	err := svc.validatePhoneTypeName(ctx, model)
	if err != nil {
		return nil, err
	}

	if model.IsDefault {
		err = svc.resetDefaultPhoneType(ctx, model)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.PhoneTypes().CreatePhoneType(ctx, model)
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
		return nil, contact.ErrPhoneTypeNotFound
	}

	data.IsDefault = model.IsDefault
	data.Translations = model.Translations

	err = svc.validatePhoneTypeName(ctx, data)
	if err != nil {
		return nil, err
	}

	if data.IsDefault {
		err = svc.resetDefaultPhoneType(ctx, data)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.PhoneTypes().UpdatePhoneType(ctx, data)
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
		return contact.ErrPhoneTypeNotFound
	}
	if data.IsDefault {
		return contact.ErrPhoneTypeIsDefault
	}
	if data.IsSystem {
		return contact.ErrPhoneTypeReadOnly
	}

	err = svc.database.PhoneTypes().DeletePhoneType(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) resetDefaultPhoneType(ctx context.Context, model *model.PhoneType) error {
	current, err := svc.database.PhoneTypes().GetDefaultPhoneType(ctx)
	if err != nil {
		return err
	}
	if current != nil && current.Id != model.Id {
		current.IsDefault = false
		err = svc.database.PhoneTypes().UpdatePhoneType(ctx, current)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) validatePhoneTypeName(ctx context.Context, model *model.PhoneType) error {
	if model.IsTransient() {
		existing, err := svc.database.PhoneTypes().GetPhoneTypeByKey(ctx, model.Key)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrPhoneTypeDuplicateKey
		}
	}

	for _, translation := range model.Translations {
		existing, err := svc.database.PhoneTypes().GetPhoneTypeByName(ctx, translation.Language, translation.Name)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrPhoneTypeDuplicateName
		}
	}
	return nil
}

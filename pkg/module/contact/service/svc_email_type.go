package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetEmailTypes(ctx context.Context, offset int64, limit int64, filter *model.EmailTypeFilter, sort *core.Sort) ([]*model.EmailType, int64, error) {
	data, count, err := svc.database.EmailTypes().GetEmailTypes(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetEmailTypeById(ctx context.Context, id string) (*model.EmailType, error) {
	data, err := svc.database.EmailTypes().GetEmailTypeById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrEmailTypeNotFound
	}

	return data, nil
}

func (svc *service) CreateEmailType(ctx context.Context, model *model.EmailType) (*model.EmailType, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = ""

	err := svc.validateEmailTypeName(ctx, model)
	if err != nil {
		return nil, err
	}

	if model.IsDefault {
		err = svc.resetDefaultEmailType(ctx, model)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.EmailTypes().CreateEmailType(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetEmailTypeById(ctx, newId)
}

func (svc *service) UpdateEmailType(ctx context.Context, id string, model *model.EmailType) (*model.EmailType, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = id

	data, err := svc.GetEmailTypeById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrEmailTypeNotFound
	}
	data.UpdateModel(model)

	err = svc.validateEmailTypeName(ctx, data)
	if err != nil {
		return nil, err
	}

	if data.IsDefault {
		err = svc.resetDefaultEmailType(ctx, data)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.EmailTypes().UpdateEmailType(ctx, data)
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
	if data.IsDefault {
		return contact.ErrEmailTypeIsDefault
	}
	if data.IsSystem {
		return contact.ErrEmailTypeReadOnly
	}

	err = svc.database.EmailTypes().DeleteEmailType(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) resetDefaultEmailType(ctx context.Context, model *model.EmailType) error {
	current, err := svc.database.EmailTypes().GetDefaultEmailType(ctx)
	if err != nil {
		return err
	}
	if current != nil && current.Id != model.Id {
		current.IsDefault = false
		err = svc.database.EmailTypes().UpdateEmailType(ctx, current)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) validateEmailTypeName(ctx context.Context, model *model.EmailType) error {
	if model.IsTransient() {
		existing, err := svc.database.EmailTypes().GetEmailTypeByKey(ctx, model.Key)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrEmailTypeDuplicateKey
		}
	}

	for _, translation := range model.Translations {
		existing, err := svc.database.EmailTypes().GetEmailTypeByName(ctx, translation.Language, translation.Name)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrEmailTypeDuplicateName
		}
	}
	return nil
}

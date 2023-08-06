package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetUriTypes(ctx context.Context, offset int64, limit int64, filter *model.UriTypeFilter, sort *core.Sort) ([]*model.UriType, int64, error) {
	data, count, err := svc.database.UriTypes().GetUriTypes(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetUriTypeById(ctx context.Context, id string) (*model.UriType, error) {
	data, err := svc.database.UriTypes().GetUriTypeById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrUriTypeNotFound
	}

	return data, nil
}

func (svc *service) CreateUriType(ctx context.Context, model *model.UriType) (*model.UriType, error) {
	model.Key = svc.StringNormalizer().NormalizeString(model.Key)
	err := svc.validateUriTypeName(ctx, model)
	if err != nil {
		return nil, err
	}

	if model.IsDefault {
		err = svc.resetDefaultUriType(ctx, model)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.UriTypes().CreateUriType(ctx, model)
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

	data.IsDefault = model.IsDefault
	data.Translations = model.Translations

	err = svc.validateUriTypeName(ctx, data)
	if err != nil {
		return nil, err
	}

	if data.IsDefault {
		err = svc.resetDefaultUriType(ctx, data)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.UriTypes().UpdateUriType(ctx, data)
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
	if data.IsDefault {
		return contact.ErrUriTypeIsDefault
	}
	if data.IsSystem {
		return contact.ErrUriTypeReadOnly
	}

	err = svc.database.UriTypes().DeleteUriType(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) resetDefaultUriType(ctx context.Context, model *model.UriType) error {
	current, err := svc.database.UriTypes().GetDefaultUriType(ctx)
	if err != nil {
		return err
	}
	if current != nil && current.Id != model.Id {
		current.IsDefault = false
		err = svc.database.UriTypes().UpdateUriType(ctx, current)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) validateUriTypeName(ctx context.Context, model *model.UriType) error {
	if model.IsTransient() {
		existing, err := svc.database.UriTypes().GetUriTypeByKey(ctx, model.Key)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrUriTypeDuplicateKey
		}
	}

	for _, translation := range model.Translations {
		existing, err := svc.database.UriTypes().GetUriTypeByName(ctx, translation.Language, translation.Name)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrUriTypeDuplicateName
		}
	}
	return nil
}

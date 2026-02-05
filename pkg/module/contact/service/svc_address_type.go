package service

import (
	"context"
	"log/slog"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/logging"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetAddressTypes(ctx context.Context, offset int64, limit int64, filter *model.AddressTypeFilter, sort *core.Sort) ([]*model.AddressType, int64, error) {
	data, count, err := svc.database.AddressTypes().GetAddressTypes(ctx, offset, limit, filter, sort)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get address types from database",
			slog.Any("error", err),
		)
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetAddressTypeById(ctx context.Context, id string) (*model.AddressType, error) {
	data, err := svc.database.AddressTypes().GetAddressTypeById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get address type from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrAddressTypeNotFound
	}

	return data, nil
}

func (svc *service) CreateAddressType(ctx context.Context, model *model.AddressType) (*model.AddressType, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = ""

	err := svc.validateAddressTypeName(ctx, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate address type",
			slog.Any("error", err),
		)
		return nil, err
	}

	if model.IsDefault {
		err = svc.resetDefaultAddressType(ctx, model)
		if err != nil {
			logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to reset default address type",
				slog.Any("error", err),
			)
			return nil, err
		}
	}

	newId, err := svc.database.AddressTypes().CreateAddressType(ctx, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to create address type in database",
			slog.Any("error", err),
		)
		return nil, err
	}

	return svc.GetAddressTypeById(ctx, newId)
}

func (svc *service) UpdateAddressType(ctx context.Context, id string, model *model.AddressType) (*model.AddressType, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = id

	data, err := svc.database.AddressTypes().GetAddressTypeById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get address type from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrAddressTypeNotFound
	}
	data.UpdateModel(model)

	err = svc.validateAddressTypeName(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate address type",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}

	if data.IsDefault {
		err = svc.resetDefaultAddressType(ctx, data)
		if err != nil {
			logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to reset default address type",
				slog.String("id", id),
				slog.Any("error", err),
			)
			return nil, err
		}
	}

	err = svc.database.AddressTypes().UpdateAddressType(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to update address type in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	return svc.GetAddressTypeById(ctx, id)
}

func (svc *service) DeleteAddressType(ctx context.Context, id string) error {
	data, err := svc.database.AddressTypes().GetAddressTypeById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get address type from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	if data == nil {
		return contact.ErrAddressTypeNotFound
	}
	if data.IsDefault {
		return contact.ErrAddressTypeIsDefault
	}
	if data.IsSystem {
		return contact.ErrAddressTypeReadOnly
	}

	err = svc.database.AddressTypes().DeleteAddressType(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to delete address type in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	return nil
}

func (svc *service) resetDefaultAddressType(ctx context.Context, model *model.AddressType) error {
	current, err := svc.database.AddressTypes().GetDefaultAddressType(ctx)
	if err != nil {
		return err
	}
	if current != nil && current.Id != model.Id {
		current.IsDefault = false
		err = svc.database.AddressTypes().UpdateAddressType(ctx, current)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) validateAddressTypeName(ctx context.Context, model *model.AddressType) error {
	if model.IsTransient() {
		existing, err := svc.database.AddressTypes().GetAddressTypeByKey(ctx, model.Key)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrAddressTypeDuplicateKey
		}
	}

	for _, translation := range model.Translations {
		existing, err := svc.database.AddressTypes().GetAddressTypeByName(ctx, translation.Language, translation.Name)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrAddressTypeDuplicateName
		}
	}
	return nil
}

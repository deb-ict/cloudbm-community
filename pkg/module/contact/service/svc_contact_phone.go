package service

import (
	"context"
	"log/slog"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/logging"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetContactPhones(ctx context.Context, contactId string, offset int64, limit int64, filter *model.PhoneFilter, sort *core.Sort) ([]*model.Phone, int64, error) {
	parent, err := svc.database.Contacts().GetContactById(ctx, contactId)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact from database by id",
			slog.String("contactId", contactId),
			slog.Any("error", err),
		)
		return nil, 0, err
	}
	if parent == nil {
		return nil, 0, contact.ErrContactNotFound
	}

	data, count, err := svc.database.ContactPhones().GetContactPhones(ctx, parent, offset, limit, filter, sort)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact phones from database",
			slog.String("contactId", contactId),
			slog.Any("error", err),
		)
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetContactPhoneById(ctx context.Context, contactId string, id string) (*model.Phone, error) {
	parent, err := svc.database.Contacts().GetContactById(ctx, contactId)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact from database by id",
			slog.String("contactId", contactId),
			slog.Any("error", err),
		)
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	data, err := svc.database.ContactPhones().GetContactPhoneById(ctx, parent, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact phone from database by id",
			slog.String("contactId", contactId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrCompanyAddressNotFound
	}

	return data, nil
}

func (svc *service) CreateContactPhone(ctx context.Context, contactId string, model *model.Phone) (*model.Phone, error) {
	model.Id = ""

	parent, err := svc.database.Contacts().GetContactById(ctx, contactId)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact from database by id",
			slog.String("contactId", contactId),
			slog.Any("error", err),
		)
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	err = svc.validateContactPhone(ctx, parent, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate contact phone",
			slog.String("contactId", contactId),
			slog.Any("error", err),
		)
		return nil, err
	}

	if model.IsDefault {
		err = svc.resetDefaultContactPhone(ctx, parent, model)
		if err != nil {
			logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to reset default contact phone",
				slog.String("contactId", contactId),
				slog.Any("error", err),
			)
			return nil, err
		}
	}

	newId, err := svc.database.ContactPhones().CreateContactPhone(ctx, parent, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to create contact phone in database",
			slog.String("contactId", contactId),
			slog.Any("error", err),
		)
		return nil, err
	}
	return svc.GetContactPhoneById(ctx, contactId, newId)
}

func (svc *service) UpdateContactPhone(ctx context.Context, contactId string, id string, model *model.Phone) (*model.Phone, error) {
	model.Id = id

	parent, err := svc.database.Contacts().GetContactById(ctx, contactId)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact from database by id",
			slog.String("contactId", contactId),
			slog.Any("error", err),
		)
		return nil, err
	}
	if parent == nil {
		return nil, contact.ErrContactNotFound
	}

	data, err := svc.database.ContactPhones().GetContactPhoneById(ctx, parent, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact phone from database by id",
			slog.String("contactId", contactId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrContactPhoneNotFound
	}
	data.UpdateModel(model)

	err = svc.validateContactPhone(ctx, parent, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate contact phone",
			slog.String("contactId", contactId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}

	if data.IsDefault {
		err = svc.resetDefaultContactPhone(ctx, parent, data)
		if err != nil {
			logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to reset default contact phone",
				slog.String("contactId", contactId),
				slog.String("id", id),
				slog.Any("error", err),
			)
			return nil, err
		}
	}

	err = svc.database.ContactPhones().UpdateContactPhone(ctx, parent, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to update contact phone in database",
			slog.String("contactId", contactId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	return svc.GetContactPhoneById(ctx, contactId, id)
}

func (svc *service) DeleteContactPhone(ctx context.Context, contactId string, id string) error {
	parent, err := svc.database.Contacts().GetContactById(ctx, contactId)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact from database by id",
			slog.String("contactId", contactId),
			slog.Any("error", err),
		)
		return err
	}
	if parent == nil {
		return contact.ErrContactNotFound
	}

	data, err := svc.database.ContactPhones().GetContactPhoneById(ctx, parent, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact phone from database by id",
			slog.String("contactId", contactId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	if data == nil {
		return contact.ErrContactPhoneNotFound
	}
	if data.IsDefault {
		return contact.ErrContactPhoneIsDefault
	}

	err = svc.database.ContactPhones().DeleteContactPhone(ctx, parent, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to delete contact phone in database",
			slog.String("contactId", contactId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	return nil
}

func (svc *service) resetDefaultContactPhone(ctx context.Context, parent *model.Contact, model *model.Phone) error {
	current, err := svc.database.ContactPhones().GetDefaultContactPhone(ctx, parent)
	if err != nil {
		return err
	}
	if current != nil && current.Id != model.Id {
		current.IsDefault = false
		err = svc.database.ContactPhones().UpdateContactPhone(ctx, parent, current)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) validateContactPhone(ctx context.Context, parent *model.Contact, model *model.Phone) error {
	modelType, err := svc.database.PhoneTypes().GetPhoneTypeById(ctx, model.Type.Id)
	if err != nil {
		return err
	}
	if modelType == nil {
		return contact.ErrPhoneTypeNotFound
	}

	return nil
}

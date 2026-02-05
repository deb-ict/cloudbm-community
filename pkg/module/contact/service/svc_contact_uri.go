package service

import (
	"context"
	"log/slog"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/logging"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetContactUris(ctx context.Context, contactId string, offset int64, limit int64, filter *model.UriFilter, sort *core.Sort) ([]*model.Uri, int64, error) {
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

	data, count, err := svc.database.ContactUris().GetContactUris(ctx, parent, offset, limit, filter, sort)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact URIs from database",
			slog.String("contactId", contactId),
			slog.Any("error", err),
		)
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetContactUriById(ctx context.Context, contactId string, id string) (*model.Uri, error) {
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

	data, err := svc.database.ContactUris().GetContactUriById(ctx, parent, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact URI from database by id",
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

func (svc *service) CreateContactUri(ctx context.Context, contactId string, model *model.Uri) (*model.Uri, error) {
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

	err = svc.validateContactUri(ctx, parent, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate contact URI",
			slog.String("contactId", contactId),
			slog.Any("error", err),
		)
		return nil, err
	}

	if model.IsDefault {
		err = svc.resetDefaultContactUri(ctx, parent, model)
		if err != nil {
			logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to reset default contact URI",
				slog.String("contactId", contactId),
				slog.Any("error", err),
			)
			return nil, err
		}
	}

	newId, err := svc.database.ContactUris().CreateContactUri(ctx, parent, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to create contact URI in database",
			slog.String("contactId", contactId),
			slog.Any("error", err),
		)
		return nil, err
	}
	return svc.GetContactUriById(ctx, contactId, newId)
}

func (svc *service) UpdateContactUri(ctx context.Context, contactId string, id string, model *model.Uri) (*model.Uri, error) {
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

	data, err := svc.database.ContactUris().GetContactUriById(ctx, parent, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact URI from database by id",
			slog.String("contactId", contactId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrContactUriNotFound
	}
	data.UpdateModel(model)

	err = svc.validateContactUri(ctx, parent, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate contact URI",
			slog.String("contactId", contactId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}

	if data.IsDefault {
		err = svc.resetDefaultContactUri(ctx, parent, data)
		if err != nil {
			logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to reset default contact URI",
				slog.String("contactId", contactId),
				slog.Any("error", err),
			)
			return nil, err
		}
	}

	err = svc.database.ContactUris().UpdateContactUri(ctx, parent, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to update contact URI in database",
			slog.String("contactId", contactId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	return svc.GetContactUriById(ctx, contactId, id)
}

func (svc *service) DeleteContactUri(ctx context.Context, contactId string, id string) error {
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

	data, err := svc.database.ContactUris().GetContactUriById(ctx, parent, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact URI from database by id",
			slog.String("contactId", contactId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	if data == nil {
		return contact.ErrContactUriNotFound
	}
	if data.IsDefault {
		return contact.ErrContactUriIsDefault
	}

	err = svc.database.ContactUris().DeleteContactUri(ctx, parent, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to delete contact URI in database",
			slog.String("contactId", contactId),
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	return nil
}

func (svc *service) resetDefaultContactUri(ctx context.Context, parent *model.Contact, model *model.Uri) error {
	current, err := svc.database.ContactUris().GetDefaultContactUri(ctx, parent)
	if err != nil {
		return err
	}
	if current != nil && current.Id != model.Id {
		current.IsDefault = false
		err = svc.database.ContactUris().UpdateContactUri(ctx, parent, current)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) validateContactUri(ctx context.Context, parent *model.Contact, model *model.Uri) error {
	modelType, err := svc.database.UriTypes().GetUriTypeById(ctx, model.Type.Id)
	if err != nil {
		return err
	}
	if modelType == nil {
		return contact.ErrUriTypeNotFound
	}

	return nil
}

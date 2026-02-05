package service

import (
	"context"
	"log/slog"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/logging"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetContactTitles(ctx context.Context, offset int64, limit int64, filter *model.ContactTitleFilter, sort *core.Sort) ([]*model.ContactTitle, int64, error) {
	data, count, err := svc.database.ContactTitles().GetContactTitles(ctx, offset, limit, filter, sort)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact titles from database",
			slog.Any("error", err),
		)
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetContactTitleById(ctx context.Context, id string) (*model.ContactTitle, error) {
	data, err := svc.database.ContactTitles().GetContactTitleById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact title from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrContactTitleNotFound
	}

	return data, nil
}

func (svc *service) CreateContactTitle(ctx context.Context, model *model.ContactTitle) (*model.ContactTitle, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = ""

	err := svc.validateContactTitle(ctx, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate contact title name",
			slog.Any("error", err),
		)
		return nil, err
	}

	newId, err := svc.database.ContactTitles().CreateContactTitle(ctx, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to create contact title in database",
			slog.Any("error", err),
		)
		return nil, err
	}

	return svc.GetContactTitleById(ctx, newId)
}

func (svc *service) UpdateContactTitle(ctx context.Context, id string, model *model.ContactTitle) (*model.ContactTitle, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = id

	data, err := svc.database.ContactTitles().GetContactTitleById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact title from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrContactTitleNotFound
	}
	data.UpdateModel(model)

	err = svc.validateContactTitle(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to validate contact title name",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}

	err = svc.database.ContactTitles().UpdateContactTitle(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to update contact title in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	return svc.GetContactTitleById(ctx, id)
}

func (svc *service) DeleteContactTitle(ctx context.Context, id string) error {
	data, err := svc.database.ContactTitles().GetContactTitleById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get contact title from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	if data == nil {
		return contact.ErrContactTitleNotFound
	}
	if data.IsSystem {
		return contact.ErrContactTitleReadOnly
	}

	err = svc.database.ContactTitles().DeleteContactTitle(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to delete contact title in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	return nil
}

func (svc *service) validateContactTitle(ctx context.Context, model *model.ContactTitle) error {
	if model.IsTransient() {
		existing, err := svc.database.ContactTitles().GetContactTitleByKey(ctx, model.Key)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrContactTitleDuplicateKey
		}
	}

	for _, translation := range model.Translations {
		existing, err := svc.database.ContactTitles().GetContactTitleByName(ctx, translation.Language, translation.Name)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrContactTitleDuplicateName
		}
	}
	return nil
}

package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetJobTitles(ctx context.Context, offset int64, limit int64, filter *model.JobTitleFilter, sort *core.Sort) ([]*model.JobTitle, int64, error) {
	data, count, err := svc.database.JobTitles().GetJobTitles(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetJobTitleById(ctx context.Context, id string) (*model.JobTitle, error) {
	data, err := svc.database.JobTitles().GetJobTitleById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrJobTitleNotFound
	}

	return data, nil
}

func (svc *service) CreateJobTitle(ctx context.Context, model *model.JobTitle) (*model.JobTitle, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = ""

	err := svc.validateJobTitleName(ctx, model)
	if err != nil {
		return nil, err
	}

	newId, err := svc.database.JobTitles().CreateJobTitle(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetJobTitleById(ctx, newId)
}

func (svc *service) UpdateJobTitle(ctx context.Context, id string, model *model.JobTitle) (*model.JobTitle, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = id

	data, err := svc.GetJobTitleById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrJobTitleNotFound
	}
	data.UpdateModel(model)

	err = svc.validateJobTitleName(ctx, data)
	if err != nil {
		return nil, err
	}

	err = svc.database.JobTitles().UpdateJobTitle(ctx, data)
	if err != nil {
		return nil, err
	}
	return svc.GetJobTitleById(ctx, id)
}

func (svc *service) DeleteJobTitle(ctx context.Context, id string) error {
	data, err := svc.GetJobTitleById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrJobTitleNotFound
	}
	if data.IsSystem {
		return contact.ErrJobTitleReadOnly
	}

	err = svc.database.JobTitles().DeleteJobTitle(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) validateJobTitleName(ctx context.Context, model *model.JobTitle) error {
	if model.IsTransient() {
		existing, err := svc.database.JobTitles().GetJobTitleByKey(ctx, model.Key)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrJobTitleDuplicateKey
		}
	}

	for _, translation := range model.Translations {
		existing, err := svc.database.JobTitles().GetJobTitleByName(ctx, translation.Language, translation.Name)
		if err != nil {
			return err
		}
		if existing != nil {
			return contact.ErrJobTitleDuplicateName
		}
	}
	return nil
}

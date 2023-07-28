package service

import (
	"context"
	"strings"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetJobTitles(ctx context.Context, offset int64, limit int64, filter *model.JobTitleFilter, sort *core.Sort) ([]*model.JobTitle, int64, error) {
	data, count, err := svc.database.JobTitleRepository().GetJobTitles(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetJobTitleById(ctx context.Context, id string) (*model.JobTitle, error) {
	data, err := svc.database.JobTitleRepository().GetJobTitleById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrJobTitleNotFound
	}

	return data, nil
}

func (svc *service) CreateJobTitle(ctx context.Context, model *model.JobTitle) (*model.JobTitle, error) {
	model.Key = strings.ToLower(model.Key)
	existing, err := svc.database.JobTitleRepository().GetJobTitleByKey(ctx, model.Key)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, contact.ErrJobTitleDuplicateKey
	}

	newId, err := svc.database.JobTitleRepository().CreateJobTitle(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetJobTitleById(ctx, newId)
}

func (svc *service) UpdateJobTitle(ctx context.Context, id string, model *model.JobTitle) (*model.JobTitle, error) {
	data, err := svc.GetJobTitleById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return contact.ErrXXXNotFound
	}

	//TODO: Set fields

	err = svc.database.JobTitleRepository().UpdateJobTitle(ctx, data)
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
		return contact.ErrXXXNotFound
	}

	//TODO: Check dependencies
	count, err := 0, nil
	if err != nil {
		return err
	}
	if count > 0 {
		return contact.ErrJobTitleInUse
	}

	err = svc.database.JobTitleRepository().DeleteJobTitle(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

package service

import (
	"context"
	"strings"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

func (svc *service) GetIndustries(ctx context.Context, offset int64, limit int64, filter *model.IndustryFilter, sort *core.Sort) ([]*model.Industry, int64, error) {
	data, count, err := svc.database.IndustryRepository().GetIndustries(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetIndustryById(ctx context.Context, id string) (*model.Industry, error) {
	data, err := svc.database.IndustryRepository().GetIndustryById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrIndustryNotFound
	}

	return data, nil
}

func (svc *service) CreateIndustry(ctx context.Context, model *model.Industry) (*model.Industry, error) {
	model.Key = strings.ToLower(model.Key)
	existing, err := svc.database.IndustryRepository().GetIndustryByKey(ctx, model.Key)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, contact.ErrIndustryDuplicateKey
	}

	//TODO: Check for duplicates on name

	newId, err := svc.database.IndustryRepository().CreateIndustry(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetIndustryById(ctx, newId)
}

func (svc *service) UpdateIndustry(ctx context.Context, id string, model *model.Industry) (*model.Industry, error) {
	data, err := svc.GetIndustryById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, contact.ErrIndustryNotFound
	}

	//TODO: Check for duplicates on name

	//TODO: Set fields

	err = svc.database.IndustryRepository().UpdateIndustry(ctx, data)
	if err != nil {
		return nil, err
	}
	return svc.GetIndustryById(ctx, id)
}

func (svc *service) DeleteIndustry(ctx context.Context, id string) error {
	data, err := svc.GetIndustryById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return contact.ErrIndustryNotFound
	}

	//TODO: Check dependencies

	err = svc.database.IndustryRepository().DeleteIndustry(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

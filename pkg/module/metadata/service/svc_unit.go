package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/metadata"
	"github.com/deb-ict/cloudbm-community/pkg/module/metadata/model"
)

func (svc *service) GetUnits(ctx context.Context, offset int64, limit int64, filter *model.UnitFilter, sort *core.Sort) ([]*model.Unit, int64, error) {
	data, count, err := svc.database.Units().GetUnits(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}
	return data, count, nil
}

func (svc *service) GetUnitById(ctx context.Context, id string) (*model.Unit, error) {
	data, err := svc.database.Units().GetUnitById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, metadata.ErrUnitNotFound
	}
	return data, nil
}

func (svc *service) CreateUnit(ctx context.Context, model *model.Unit) (*model.Unit, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = ""

	err := svc.validateUnit(ctx, model)
	if err != nil {
		return nil, err
	}

	newId, err := svc.database.Units().CreateUnit(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetUnitById(ctx, newId)
}

func (svc *service) UpdateUnit(ctx context.Context, id string, model *model.Unit) (*model.Unit, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = id

	data, err := svc.database.Units().GetUnitById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, metadata.ErrUnitNotFound
	}
	data.UpdateModel(model)

	err = svc.validateUnit(ctx, data)
	if err != nil {
		return nil, err
	}

	err = svc.database.Units().UpdateUnit(ctx, data)
	if err != nil {
		return nil, err
	}

	return svc.GetUnitById(ctx, id)
}

func (svc *service) DeleteUnit(ctx context.Context, id string) error {
	data, err := svc.database.Units().GetUnitById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return metadata.ErrUnitNotFound
	}

	err = svc.database.Units().DeleteUnit(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (svc *service) validateUnit(ctx context.Context, model *model.Unit) error {
	if model.IsTransient() {
		existing, err := svc.database.Units().GetUnitByKey(ctx, model.Key)
		if err != nil {
			return err
		}
		if existing != nil {
			return metadata.ErrUnitDuplicateKey
		}
	}

	for _, translation := range model.Translations {
		existing, err := svc.database.Units().GetUnitByName(ctx, translation.Language, translation.Name)
		if err != nil {
			return err
		}
		if existing != nil {
			return metadata.ErrUnitDuplicateName
		}
	}

	return nil
}

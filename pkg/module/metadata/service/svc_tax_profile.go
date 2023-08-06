package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/metadata"
	"github.com/deb-ict/cloudbm-community/pkg/module/metadata/model"
)

func (svc *service) GetTaxProfiles(ctx context.Context, offset int64, limit int64, filter *model.TaxProfileFilter, sort *core.Sort) ([]*model.TaxProfile, int64, error) {
	data, count, err := svc.database.TaxProfiles().GetTaxProfiles(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}
	return data, count, nil
}

func (svc *service) GetTaxProfileById(ctx context.Context, id string) (*model.TaxProfile, error) {
	data, err := svc.database.TaxProfiles().GetTaxProfileById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, metadata.ErrTaxProfileNotFound
	}
	return data, nil
}

func (svc *service) CreateTaxProfile(ctx context.Context, model *model.TaxProfile) (*model.TaxProfile, error) {
	model.Key = svc.StringNormalizer().NormalizeString(model.Key)
	err := svc.validateTaxProfile(ctx, model)
	if err != nil {
		return nil, err
	}

	newId, err := svc.database.TaxProfiles().CreateTaxProfile(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetTaxProfileById(ctx, newId)
}

func (svc *service) UpdateTaxProfile(ctx context.Context, id string, model *model.TaxProfile) (*model.TaxProfile, error) {
	data, err := svc.database.TaxProfiles().GetTaxProfileById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, metadata.ErrTaxProfileNotFound
	}

	data.Translations = model.Translations
	data.Rate = model.Rate

	err = svc.validateTaxProfile(ctx, data)
	if err != nil {
		return nil, err
	}

	err = svc.database.TaxProfiles().UpdateTaxProfile(ctx, data)
	if err != nil {
		return nil, err
	}

	return svc.GetTaxProfileById(ctx, id)
}

func (svc *service) DeleteTaxProfile(ctx context.Context, id string) error {
	data, err := svc.database.TaxProfiles().GetTaxProfileById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return metadata.ErrTaxProfileNotFound
	}

	err = svc.database.TaxProfiles().DeleteTaxProfile(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (svc *service) validateTaxProfile(ctx context.Context, model *model.TaxProfile) error {
	if model.IsTransient() {
		existing, err := svc.database.TaxProfiles().GetTaxProfileByKey(ctx, model.Key)
		if err != nil {
			return err
		}
		if existing != nil {
			return metadata.ErrTaxProfileDuplicateKey
		}
	}

	for _, translation := range model.Translations {
		existing, err := svc.database.TaxProfiles().GetTaxProfileByName(ctx, translation.Language, translation.Name)
		if err != nil {
			return err
		}
		if existing != nil {
			return metadata.ErrTaxProfileDuplicateName
		}
	}

	return nil
}

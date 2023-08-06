package metadata

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/metadata/model"
)

type Database interface {
	TaxProfiles() TaxProfileRepository
}

type TaxProfileRepository interface {
	GetTaxProfiles(ctx context.Context, offset int64, limit int64, filter *model.TaxProfileFilter, sort *core.Sort) ([]*model.TaxProfile, int64, error)
	GetTaxProfileById(ctx context.Context, id string) (*model.TaxProfile, error)
	GetTaxProfileByKey(ctx context.Context, key string) (*model.TaxProfile, error)
	GetTaxProfileByName(ctx context.Context, language string, name string) (*model.TaxProfile, error)
	CreateTaxProfile(ctx context.Context, model *model.TaxProfile) (string, error)
	UpdateTaxProfile(ctx context.Context, model *model.TaxProfile) error
	DeleteTaxProfile(ctx context.Context, model *model.TaxProfile) error
}

package global

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module"
	"github.com/deb-ict/cloudbm-community/pkg/module/global/model"
)

type Service interface {
	module.Service

	GetTaxProfiles(ctx context.Context, offset int64, limit int64, filter *model.TaxProfileFilter, sort *core.Sort) ([]*model.TaxProfile, int64, error)
	GetTaxProfileById(ctx context.Context, id string) (*model.TaxProfile, error)
	CreateTaxProfile(ctx context.Context, model *model.TaxProfile) (*model.TaxProfile, error)
	UpdateTaxProfile(ctx context.Context, id string, model *model.TaxProfile) (*model.TaxProfile, error)
	DeleteTaxProfile(ctx context.Context, id string) error
}

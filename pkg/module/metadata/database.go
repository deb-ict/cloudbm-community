package metadata

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/metadata/model"
)

type Database interface {
	Units() UnitRepository
	TaxRates() TaxRateRepository
}

type UnitRepository interface {
	GetUnits(ctx context.Context, offset int64, limit int64, filter *model.UnitFilter, sort *core.Sort) ([]*model.Unit, int64, error)
	GetUnitById(ctx context.Context, id string) (*model.Unit, error)
	GetUnitByKey(ctx context.Context, key string) (*model.Unit, error)
	GetUnitByName(ctx context.Context, language string, name string) (*model.Unit, error)
	CreateUnit(ctx context.Context, model *model.Unit) (string, error)
	UpdateUnit(ctx context.Context, model *model.Unit) error
	DeleteUnit(ctx context.Context, model *model.Unit) error
}

type TaxRateRepository interface {
	GetTaxRates(ctx context.Context, offset int64, limit int64, filter *model.TaxRateFilter, sort *core.Sort) ([]*model.TaxRate, int64, error)
	GetTaxRateById(ctx context.Context, id string) (*model.TaxRate, error)
	GetTaxRateByKey(ctx context.Context, key string) (*model.TaxRate, error)
	GetTaxRateByName(ctx context.Context, language string, name string) (*model.TaxRate, error)
	CreateTaxRate(ctx context.Context, model *model.TaxRate) (string, error)
	UpdateTaxRate(ctx context.Context, model *model.TaxRate) error
	DeleteTaxRate(ctx context.Context, model *model.TaxRate) error
}

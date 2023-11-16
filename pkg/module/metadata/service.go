package metadata

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/metadata/model"
)

type Service interface {
	StringNormalizer() core.StringNormalizer
	FeatureProvider() core.FeatureProvider
	LanguageProvider() localization.LanguageProvider

	GetUnits(ctx context.Context, offset int64, limit int64, filter *model.UnitFilter, sort *core.Sort) ([]*model.Unit, int64, error)
	GetUnitById(ctx context.Context, id string) (*model.Unit, error)
	CreateUnit(ctx context.Context, model *model.Unit) (*model.Unit, error)
	UpdateUnit(ctx context.Context, id string, model *model.Unit) (*model.Unit, error)
	DeleteUnit(ctx context.Context, id string) error

	GetTaxRates(ctx context.Context, offset int64, limit int64, filter *model.TaxRateFilter, sort *core.Sort) ([]*model.TaxRate, int64, error)
	GetTaxRateById(ctx context.Context, id string) (*model.TaxRate, error)
	CreateTaxRate(ctx context.Context, model *model.TaxRate) (*model.TaxRate, error)
	UpdateTaxRate(ctx context.Context, id string, model *model.TaxRate) (*model.TaxRate, error)
	DeleteTaxRate(ctx context.Context, id string) error
}

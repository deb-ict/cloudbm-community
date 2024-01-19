package sales

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/sales/model"
)

type Service interface {
	StringNormalizer() core.StringNormalizer
	FeatureProvider() core.FeatureProvider
	LanguageProvider() localization.LanguageProvider

	GetOrders(ctx context.Context, offset int64, limit int64, filter *model.OrderFilter, sort *core.Sort) ([]*model.Order, int64, error)
	GetOrderById(ctx context.Context, id string) (*model.Order, error)
	CreateOrder(ctx context.Context, model *model.Order) (*model.Order, error)
	UpdateOrder(ctx context.Context, id string, model *model.Order) (*model.Order, error)
	DeleteOrder(ctx context.Context, id string) error
}

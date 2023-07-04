package product

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
)

type Database interface {
	ProductRepository() ProductRepository
	CategoryRepository() CategoryRepository
}

type ProductRepository interface {
	GetProducts(ctx context.Context, offset int64, limit int64, filter *model.ProductFilter, sort *core.Sort) ([]*model.Product, int64, error)
	GetProductById(ctx context.Context, id string) (*model.Product, error)
	GetProductName(ctx context.Context, language string, name string) (*model.Product, error)
	CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
	UpdateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
	DeleteProduct(ctx context.Context, product *model.Product) error
}

type CategoryRepository interface {
	GetCategories(ctx context.Context, offset int64, limit int64, filter *model.CategoryFilter, sort *core.Sort) ([]*model.Category, int64, error)
	GetCategoryById(ctx context.Context, id string) (*model.Category, error)
	GetCategoryByName(ctx context.Context, language string, name string) (*model.Category, error)
	CreateCategory(ctx context.Context, category *model.Category) (*model.Category, error)
	UpdateCategory(ctx context.Context, category *model.Category) (*model.Category, error)
	DeleteCategory(ctx context.Context, category *model.Category) error
}

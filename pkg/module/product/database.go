package product

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
)

type Database interface {
	Products() ProductRepository
	Categories() CategoryRepository
}

type ProductRepository interface {
	GetProducts(ctx context.Context, offset int64, limit int64, filter *model.ProductFilter, sort *core.Sort) ([]*model.Product, int64, error)
	GetProductById(ctx context.Context, id string) (*model.Product, error)
	GetProductByName(ctx context.Context, language string, name string) (*model.Product, error)
	GetProductBySlug(ctx context.Context, language string, slug string) (*model.Product, error)
	CreateProduct(ctx context.Context, model *model.Product) (string, error)
	UpdateProduct(ctx context.Context, model *model.Product) error
	DeleteProduct(ctx context.Context, model *model.Product) error
}

type CategoryRepository interface {
	GetCategories(ctx context.Context, offset int64, limit int64, filter *model.CategoryFilter, sort *core.Sort) ([]*model.Category, int64, error)
	GetCategoryById(ctx context.Context, id string) (*model.Category, error)
	GetCategoryByName(ctx context.Context, language string, name string) (*model.Category, error)
	GetCategoryBySlug(ctx context.Context, language string, slug string) (*model.Category, error)
	CreateCategory(ctx context.Context, model *model.Category) (string, error)
	UpdateCategory(ctx context.Context, model *model.Category) error
	DeleteCategory(ctx context.Context, model *model.Category) error
}

package product

import (
	"context"
	"errors"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
)

var (
	ErrProductNotFound    error = errors.New("product not found")
	ErrProductNotCreated  error = errors.New("product not created")
	ErrProductNotUpdated  error = errors.New("product not updated")
	ErrProductNotDeleted  error = errors.New("product not deleted")
	ErrCategoryNotFound   error = errors.New("category not found")
	ErrCategoryNotCreated error = errors.New("category not created")
	ErrCategoryNotUpdated error = errors.New("category not updated")
	ErrCategoryNotDeleted error = errors.New("category not deleted")
)

type ServiceOption func(Service) error

type Service interface {
	GetDatabase() Database
	SetDatabase(database Database) error

	GetProducts(ctx context.Context, offset int64, limit int64, filter *model.ProductFilter, sort *core.Sort) ([]*model.Product, int64, error)
	GetProductById(ctx context.Context, id string) (*model.Product, error)
	GetProductByName(ctx context.Context, language string, name string) (*model.Product, error)
	GetProductBySlug(ctx context.Context, language string, slug string) (*model.Product, error)
	CreateProduct(ctx context.Context, model *model.Product) (*model.Product, error)
	UpdateProduct(ctx context.Context, id string, model *model.Product) (*model.Product, error)
	DeleteProduct(ctx context.Context, id string) error

	GetCategories(ctx context.Context, offset int64, limit int64, filter *model.CategoryFilter, sort *core.Sort) ([]*model.Category, int64, error)
	GetCategoryById(ctx context.Context, id string) (*model.Category, error)
	GetCategoryByName(ctx context.Context, language string, name string) (*model.Category, error)
	GetCategoryBySlug(ctx context.Context, language string, slug string) (*model.Category, error)
	CreateCategory(ctx context.Context, model *model.Category) (*model.Category, error)
	UpdateCategory(ctx context.Context, id string, model *model.Category) (*model.Category, error)
	DeleteCategory(ctx context.Context, id string) error
}

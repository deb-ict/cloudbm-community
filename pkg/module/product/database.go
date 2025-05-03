package product

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
)

type Database interface {
	Attributes() AttributeRepository
	AttributeValues() AttributeValueRepository
	Categories() CategoryRepository
	Products() ProductRepository
}

type AttributeRepository interface {
	GetAttributes(ctx context.Context, offset int64, limit int64, filter *model.AttributeFilter, sort *core.Sort) ([]*model.Attribute, int64, error)
	GetAttributeById(ctx context.Context, id string) (*model.Attribute, error)
	GetAttributeByName(ctx context.Context, language string, name string) (*model.Attribute, error)
	GetAttributeBySlug(ctx context.Context, language string, slug string) (*model.Attribute, error)
	CreateAttribute(ctx context.Context, model *model.Attribute) (string, error)
	UpdateAttribute(ctx context.Context, model *model.Attribute) error
	DeleteAttribute(ctx context.Context, model *model.Attribute) error
}

type AttributeValueRepository interface {
	GetAttributeValues(ctx context.Context, parent *model.Attribute, offset int64, limit int64, filter *model.AttributeValueFilter, sort *core.Sort) ([]*model.AttributeValue, int64, error)
	GetAttributeValueById(ctx context.Context, parent *model.Attribute, id string) (*model.AttributeValue, error)
	GetAttributeValueByName(ctx context.Context, parent *model.Attribute, language string, name string) (*model.AttributeValue, error)
	GetAttributeValueBySlug(ctx context.Context, parent *model.Attribute, language string, slug string) (*model.AttributeValue, error)
	CreateAttributeValue(ctx context.Context, parent *model.Attribute, model *model.AttributeValue) (string, error)
	UpdateAttributeValue(ctx context.Context, parent *model.Attribute, model *model.AttributeValue) error
	DeleteAttributeValue(ctx context.Context, parent *model.Attribute, model *model.AttributeValue) error
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

type ProductRepository interface {
	GetProducts(ctx context.Context, offset int64, limit int64, filter *model.ProductFilter, sort *core.Sort) ([]*model.Product, int64, error)
	GetProductById(ctx context.Context, id string) (*model.Product, error)
	GetProductByName(ctx context.Context, language string, name string) (*model.Product, error)
	GetProductBySlug(ctx context.Context, language string, slug string) (*model.Product, error)
	CreateProduct(ctx context.Context, model *model.Product) (string, error)
	UpdateProduct(ctx context.Context, model *model.Product) error
	DeleteProduct(ctx context.Context, model *model.Product) error
}

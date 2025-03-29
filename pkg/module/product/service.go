package product

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
)

type Service interface {
	StringNormalizer() core.StringNormalizer
	FeatureProvider() core.FeatureProvider
	LanguageProvider() localization.LanguageProvider

	GetAttributes(ctx context.Context, offset int64, limit int64, filter *model.AttributeFilter, sort *core.Sort) ([]*model.Attribute, int64, error)
	GetAttributeById(ctx context.Context, id string) (*model.Attribute, error)
	GetAttributeByName(ctx context.Context, language string, name string) (*model.Attribute, error)
	GetAttributeBySlug(ctx context.Context, language string, slug string) (*model.Attribute, error)
	CreateAttribute(ctx context.Context, model *model.Attribute) (*model.Attribute, error)
	UpdateAttribute(ctx context.Context, id string, model *model.Attribute) (*model.Attribute, error)
	DeleteAttribute(ctx context.Context, id string) error

	GetAttributeValues(ctx context.Context, attributeId string, offset int64, limit int64, filter *model.AttributeValueFilter, sort *core.Sort) ([]*model.AttributeValue, int64, error)
	GetAttributeValueById(ctx context.Context, attributeId string, id string) (*model.AttributeValue, error)
	GetAttributeValueByName(ctx context.Context, attributeId string, language string, name string) (*model.AttributeValue, error)
	GetAttributeValueBySlug(ctx context.Context, attributeId string, language string, slug string) (*model.AttributeValue, error)
	CreateAttributeValue(ctx context.Context, attributeId string, model *model.AttributeValue) (*model.AttributeValue, error)
	UpdateAttributeValue(ctx context.Context, attributeId string, id string, model *model.AttributeValue) (*model.AttributeValue, error)
	DeleteAttributeValue(ctx context.Context, attributeId string, id string) error

	GetCategories(ctx context.Context, offset int64, limit int64, filter *model.CategoryFilter, sort *core.Sort) ([]*model.Category, int64, error)
	GetCategoryById(ctx context.Context, id string) (*model.Category, error)
	GetCategoryByName(ctx context.Context, language string, name string) (*model.Category, error)
	GetCategoryBySlug(ctx context.Context, language string, slug string) (*model.Category, error)
	CreateCategory(ctx context.Context, model *model.Category) (*model.Category, error)
	UpdateCategory(ctx context.Context, id string, model *model.Category) (*model.Category, error)
	DeleteCategory(ctx context.Context, id string) error

	GetProducts(ctx context.Context, offset int64, limit int64, filter *model.ProductFilter, sort *core.Sort) ([]*model.Product, int64, error)
	GetProductById(ctx context.Context, id string) (*model.Product, error)
	GetProductByName(ctx context.Context, language string, name string) (*model.Product, error)
	GetProductBySlug(ctx context.Context, language string, slug string) (*model.Product, error)
	CreateProduct(ctx context.Context, model *model.Product) (*model.Product, error)
	UpdateProduct(ctx context.Context, id string, model *model.Product) (*model.Product, error)
	DeleteProduct(ctx context.Context, id string) error

	GetProductVariants(ctx context.Context, productId string, offset int64, limit int64, filter *model.ProductVariantFilter, sort *core.Sort) ([]*model.ProductVariant, int64, error)
	GetProductVariantById(ctx context.Context, productId string, id string) (*model.ProductVariant, error)
	GetProductVariantByName(ctx context.Context, productId string, language string, name string) (*model.ProductVariant, error)
	GetProductVariantBySlug(ctx context.Context, productId string, language string, slug string) (*model.ProductVariant, error)
	CreateProductVariant(ctx context.Context, productId string, model *model.ProductVariant) (*model.ProductVariant, error)
	UpdateProductVariant(ctx context.Context, productId string, id string, model *model.ProductVariant) (*model.ProductVariant, error)
	DeleteProductVariant(ctx context.Context, productId string, id string) error
}

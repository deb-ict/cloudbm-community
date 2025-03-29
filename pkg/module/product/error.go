package product

import "errors"

var (
	ErrAttributeNotFound           error = errors.New("attribute not found")
	ErrAttributeDuplicateName      error = errors.New("attribute with same name exists")
	ErrAttributeDuplicateSlug      error = errors.New("attribute with same slug exists")
	ErrAttributeValueNotFound      error = errors.New("attribute value not found")
	ErrAttributeValueDuplicateName error = errors.New("attribute value with same name exists")
	ErrAttributeValueDuplicateSlug error = errors.New("attribute value with same slug exists")
	ErrCategoryNotFound            error = errors.New("category not found")
	ErrCategoryDuplicateName       error = errors.New("category with same name exists")
	ErrCategoryDuplicateSlug       error = errors.New("category with same slug exists")
	ErrProductNotFound             error = errors.New("product not found")
	ErrProductDuplicateName        error = errors.New("product with same name exists")
	ErrProductDuplicateSlug        error = errors.New("product with same slug exists")
	ErrProductVariantNotFound      error = errors.New("product variant not found")
	ErrProductVariantDuplicateName error = errors.New("product variant with same name exists")
	ErrProductVariantDuplicateSlug error = errors.New("product variant with same slug exists")
)

package product

import "errors"

var (
	ErrProductNotFound       error = errors.New("product not found")
	ErrProductDuplicateName  error = errors.New("product with same name exists")
	ErrProductDuplicateSlug  error = errors.New("product with same slug exists")
	ErrCategoryNotFound      error = errors.New("category not found")
	ErrCategoryDuplicateName error = errors.New("category with same name exists")
	ErrCategoryDuplicateSlug error = errors.New("category with same slug exists")
)

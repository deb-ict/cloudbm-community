package product

import "errors"

var (
	ErrProductNotFound  error = errors.New("product not found")
	ErrCategoryNotFound error = errors.New("category not found")
)

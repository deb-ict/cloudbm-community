package gallery

import "errors"

var (
	ErrCategoryNotFound error = errors.New("category not found")
)

var (
	ErrMediaNotFound error = errors.New("media not found")
)

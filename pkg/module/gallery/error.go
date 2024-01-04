package gallery

import "errors"

var (
	ErrImageNotFound      error = errors.New("image not found")
	ErrImageDuplicateName error = errors.New("image with same name exists")
	ErrImageDuplicateSlug error = errors.New("image with same slug exists")
)

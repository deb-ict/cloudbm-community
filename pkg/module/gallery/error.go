package gallery

import "errors"

var (
	ErrImageNotFound           error = errors.New("image not found")
	ErrImageFileNotFound       error = errors.New("image file not found")
	ErrImageFormatNotSupported error = errors.New("image format not supported")
	ErrImageDuplicateName      error = errors.New("image with same name exists")
	ErrImageDuplicateSlug      error = errors.New("image with same slug exists")
)

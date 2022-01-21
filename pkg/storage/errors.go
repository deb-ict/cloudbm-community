package storage

import "errors"

var (
	ErrDbNotFound   = errors.New("not found")
	ErrDbNotCreated = errors.New("not created")
	ErrDbNotUpdated = errors.New("not updated")
	ErrDbNotDeleted = errors.New("not deleted")
)

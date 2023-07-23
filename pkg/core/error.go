package core

import "errors"

var (
	ErrInvalidId        error = errors.New("invalid id")
	ErrRecordNotCreated error = errors.New("record not created")
	ErrRecordNotChanged error = errors.New("record not changed")
	ErrRecordNotDeleted error = errors.New("record not changed")
)

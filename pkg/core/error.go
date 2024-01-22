package core

import "errors"

var (
	ErrInvalidId           error = errors.New("invalid id")
	ErrNilObject           error = errors.New("nil object")
	ErrRecordNotCreated    error = errors.New("record not created")
	ErrRecordNotChanged    error = errors.New("record not changed")
	ErrRecordNotDeleted    error = errors.New("record not changed")
	ErrTranslationNotFound error = errors.New("translation not found")
	ErrNotImplemented      error = errors.New("method not implemented")
)

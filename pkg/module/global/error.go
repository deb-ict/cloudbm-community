package global

import "errors"

var (
	ErrTaxProfileNotFound      error = errors.New("tax profile not found")
	ErrTaxProfileDuplicateKey  error = errors.New("tax profile with same key exists")
	ErrTaxProfileDuplicateName error = errors.New("tax profile with same language/name exists")
)

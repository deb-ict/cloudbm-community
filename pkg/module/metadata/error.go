package metadata

import "errors"

var (
	ErrUnitNotFound      error = errors.New("unit not found")
	ErrUnitDuplicateKey  error = errors.New("unit with same key exists")
	ErrUnitDuplicateName error = errors.New("unit with same language/name exists")
	ErrUnitDisabled      error = errors.New("unit disabled")
)

var (
	ErrTaxRateNotFound      error = errors.New("tax rate not found")
	ErrTaxRateDuplicateKey  error = errors.New("tax rate with same key exists")
	ErrTaxRateDuplicateName error = errors.New("tax rate with same language/name exists")
	ErrTaxRateDisabled      error = errors.New("tax rate disabled")
)

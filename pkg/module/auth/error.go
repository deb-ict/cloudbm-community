package auth

import "errors"

var (
	ErrUserNotFound      error = errors.New("user not found")
	ErrUserLocked        error = errors.New("user has been locked")
	ErrPasswordNotMatch  error = errors.New("password not match")
	ErrDuplicateUsername error = errors.New("user with same username exists")
	ErrDuplicateEmail    error = errors.New("user with same email exists")
	ErrInvalidToken      error = errors.New("invalid token")
	ErrTokenExpired      error = errors.New("token has expired")
)

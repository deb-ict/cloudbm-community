package session

import "errors"

var (
	ErrSessionNotFound error = errors.New("session not found")
	ErrSessionExpired  error = errors.New("session expired")
)

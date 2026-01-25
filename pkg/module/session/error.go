package session

import "errors"

var (
	ErrSessionNotFound     error = errors.New("session not found")
	ErrSessionDataNotFound error = errors.New("session data not found")
	ErrSessionExpired      error = errors.New("session expired")
)

package error

import "errors"

var (
	ErrKeyDoesNotExist = errors.New("key does not exist")
	ErrTokenInvalid    = errors.New("token invalid")
	ErrTokenExpired    = errors.New("token expired")
)

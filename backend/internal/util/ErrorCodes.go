package util

import "errors"

var (
	ErrInternal = errors.New("internal server error")
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenMissing                          = errors.New("token missing")
)

var CustomError = map[error]int{
	ErrInternal: 400,
	ErrInvalidToken: 400,
}

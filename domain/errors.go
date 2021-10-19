package domain

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrInvalidEntity       = errors.New("invalid entity")
	ErrNotFound            = errors.New("requested item not found")
	ErrConflict            = errors.New("item already exists")
	ErrBadParamInput       = errors.New("param not valid")
)

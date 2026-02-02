package domain

import "errors"

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
	ErrInvalidTitle  = errors.New("invalid title")
)

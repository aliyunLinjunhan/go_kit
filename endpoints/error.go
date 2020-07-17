package endpoints

import (
	"errors"
)

var (
	// ErrInvalidRequestType is used to
	ErrInvalidRequestType error = errors.New("InvalidRequest")
)

package bthome

import (
	"errors"
)

var (
	errInvalidSize  = errors.New("invalid size")
	errBufferFull   = errors.New("buffer full")
	errDataNotFound = errors.New("field not found")
)

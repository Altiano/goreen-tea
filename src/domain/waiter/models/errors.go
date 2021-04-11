package waiterModels

import "errors"

var (
	ErrEmptyList = errors.New("no more waiters left")
)

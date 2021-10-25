package storage

import "errors"

var(
	ErrSectorNotFound = errors.New("sector was not found by the provided parameters")
)

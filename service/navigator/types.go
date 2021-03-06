package navigator

import "errors"

var (
	ErrParseParameter = errors.New("can't parse provided parameter")
	ErrEmptyParameter = errors.New("parameter is empty")
	ErrSectorNotFound = errors.New("sector was not found by the provided parameters")
)

type Coordinates struct {
	X   float64
	Y   float64
	Z   float64
	Vel float64
}

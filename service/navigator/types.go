package navigator

import "errors"

var (
	ErrConvertingParameter = errors.New("can't convert incoming parameter")
)

type Coordinates struct {
	X   float64
	Y   float64
	Z   float64
	Vel float64
}

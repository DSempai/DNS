package domain

type SectorID int64

const (
	id SectorID = 123
)

type DroneNavigation interface {
	LocateDatabankByCoordinates(coordinates DroneCoordinates) (Location, error)
}

type Location float64

type DroneCoordinates struct {
	X   string
	Y   string
	Z   string
	Vel string
}

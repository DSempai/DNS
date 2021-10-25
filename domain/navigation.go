package domain

type SectorID int64
type Location float64

var (
	ID = SectorID(1)
)

type DroneNavigation interface {
	LocateDatabankByCoordinates(coordinates DroneCoordinates, id SectorID) (*DatabankLocation, error)
}

type DatabankLocation struct {
	Location `json:"loc"`
}

type DroneCoordinates struct {
	X   string `json:"x"`
	Y   string `json:"y"`
	Z   string `json:"z"`
	Vel string `json:"vel"`
}

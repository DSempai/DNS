package domain

type (
	SectorID int64
	Location float64
)

// DroneNavigation is our domain entity that describe all business functions.
type DroneNavigation interface {
	LocateDatabankByCoordinates(coordinates DroneCoordinates) (*DatabankLocation, error)
}

// DatabankLocation is our "response" to drone.
type DatabankLocation struct {
	Location `json:"loc"`
}

// DroneCoordinates format that we accepted as request from drones.
type DroneCoordinates struct {
	X   string `json:"x"`
	Y   string `json:"y"`
	Z   string `json:"z"`
	Vel string `json:"vel"`
}

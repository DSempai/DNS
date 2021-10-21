package api

type LocateDatabankRequest struct {
	X   string `json:"x"`
	Y   string `json:"y"`
	Z   string `json:"z"`
	Vel string `json:"vel"`
}

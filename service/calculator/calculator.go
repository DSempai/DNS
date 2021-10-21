package calculator

import (
	"github.com/sirupsen/logrus"
)

type Service struct {
	Logger *logrus.Logger
}

func Initialize(logger *logrus.Logger) Service {
	return Service{
		Logger: logger,
	}
}

func (s Service) AddCoordinates(coordinates ...float64) float64 {
	var sumOfCoordinates float64
	for _, coordinate := range coordinates {
		sumOfCoordinates += coordinate
	}
	return sumOfCoordinates
}

func (s Service) MultiplyCoordinate(coordinate float64, sectorID float64) float64 {
	return coordinate * float64(sectorID)
}

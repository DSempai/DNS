package navigator

import (
	"DNS/domain"
	"DNS/repository/storage"
	"DNS/service/calculator"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
)

type Service struct {
	Logger     *logrus.Logger
	Storage    storage.Storage
	Calculator calculator.Service
}

func Initialize(logger *logrus.Logger, s storage.Storage, c calculator.Service) Service {
	return Service{
		Logger:     logger,
		Storage:    s,
		Calculator: c,
	}
}

func (s Service) LocateDatabankByCoordinates(coordinates domain.DroneCoordinates) (domain.Location, error) {
	var location domain.Location

	coords, err := s.convertCoordinates(coordinates)
	if err != nil {
		s.Logger.Infof("Error while converting coordinates with error: %#v", err)
		return location, err
	}
	destinationPoint, err := s.locate(coords)
	if err != nil {
		s.Logger.Infof("Error while locate databank with error: %#v", err)
		return location, err
	}

	return domain.Location(destinationPoint), nil
}

func (s Service) locate(coordinates *Coordinates) (float64, error) {
	//destinationPoint := s.Calculator.MultiplyCoordinate(coordinates.X,)
	return 0, nil
}

func (s Service) convertCoordinates(coordinates domain.DroneCoordinates) (*Coordinates, error) {
	x, err := s.parseIncomingParameter(coordinates.X)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", err, "axis X")
	}
	y, err := s.parseIncomingParameter(coordinates.X)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", err, "axis Y")
	}
	z, err := s.parseIncomingParameter(coordinates.X)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", err, "axis Z")
	}
	vel, err := s.parseIncomingParameter(coordinates.X)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", err, "Vel")
	}

	return &Coordinates{
		X:   x,
		Y:   y,
		Z:   z,
		Vel: vel,
	}, nil
}

func (s Service) parseIncomingParameter(axis string) (float64, error) {
	coordinate, err := strconv.ParseFloat(axis, 64)
	if err != nil {
		return 0, ErrConvertingParameter
	}
	return coordinate, nil
}

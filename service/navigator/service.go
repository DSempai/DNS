package navigator

import (
	"DNS/domain"
	"DNS/repository/storage"
	"DNS/service/calculator"
	"errors"
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Service contains all the necessary entities for calculating the location of the databank.
type Service struct {
	Logger     *logrus.Logger
	Sectors    storage.SectorsInterface
	Calculator calculator.Service
	SectorID   domain.SectorID
}

func Initialize(
	logger *logrus.Logger,
	s storage.SectorsInterface,
	c calculator.Service,
	id int64) Service {
	return Service{
		Logger:     logger,
		Sectors:    s,
		Calculator: c,
		SectorID:   domain.SectorID(id),
	}
}

func (s Service) LocateDatabankByCoordinates(coords domain.DroneCoordinates) (*domain.DatabankLocation, error) {
	coordinates, err := s.ConvertCoordinates(coords)
	if err != nil {
		s.Logger.Infof("Converting coordinates failed. Error: %v", err)
		return nil, err
	}

	sector, err := s.Sectors.RetrieveBySectorID(int64(s.SectorID))
	if err != nil {
		if errors.Is(err, storage.ErrSectorNotFound) {
			return nil, ErrSectorNotFound
		}
		s.Logger.Infof("Retrieving sector failed. Error: %v", err)
		return nil, err
	}

	destinationPoint := s.Locate(coordinates, sector.SectorID)

	return &domain.DatabankLocation{
		Location: domain.Location(destinationPoint),
	}, nil
}

// Locate return a databank location based on given coordinates, sector id and mathematical formula:
// location = X*SectorID + Y*SectorID + Z*SectorID + Vel.
func (s Service) Locate(coords *Coordinates, sectorID int64) float64 {
	return s.Calculator.AdditionFloat64(
		s.Calculator.MultiplicationFloat64(coords.X, float64(sectorID)),
		s.Calculator.MultiplicationFloat64(coords.Y, float64(sectorID)),
		s.Calculator.MultiplicationFloat64(coords.Z, float64(sectorID)),
		coords.Vel)
}

// ConvertCoordinates return converted from string floating point numbers represented in navigator.Coordinates struct.
func (s Service) ConvertCoordinates(coordinates domain.DroneCoordinates) (*Coordinates, error) {
	x, err := s.ParseParameter(coordinates.X)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", err, "axis X")
	}
	y, err := s.ParseParameter(coordinates.Y)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", err, "axis Y")
	}
	z, err := s.ParseParameter(coordinates.Z)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", err, "axis Z")
	}
	vel, err := s.ParseParameter(coordinates.Vel)
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

// ParseParameter return parsed string with 64 bit precision.
func (s Service) ParseParameter(param string) (float64, error) {
	if param == "" {
		return 0, ErrEmptyParameter
	}
	coordinate, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0, ErrParseParameter
	}
	return coordinate, nil
}

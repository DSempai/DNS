package navigator

import (
	"DNS/domain"
	"DNS/repository/storage"
	"DNS/service/calculator"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
)

type Service struct {
	Logger     *logrus.Logger
	Sectors    storage.SectorsInterface
	Calculator calculator.Service
}

func Initialize(
	logger *logrus.Logger,
	s storage.SectorsInterface,
	c calculator.Service) Service {
	return Service{
		Logger:     logger,
		Sectors:    s,
		Calculator: c,
	}
}

func (s Service) LocateDatabankByCoordinates(coords domain.DroneCoordinates, sectorID domain.SectorID) (*domain.DatabankLocation, error) {
	coordinates, err := s.ConvertCoordinates(coords)
	if err != nil {
		s.Logger.Infof("Converting coordinates failed. Error: %v", err)
		return nil, err
	}
	sector, err := s.Sectors.RetrieveBySectorID(int64(sectorID))
	if err != nil {
		if errors.Is(err, storage.ErrSectorNotFound){
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

func (s Service) Locate(coords *Coordinates, sectorID int64) float64 {
	return s.Calculator.AdditionFloat64(
		s.Calculator.MultiplicationFloat64(coords.X, float64(sectorID)),
		s.Calculator.MultiplicationFloat64(coords.Y, float64(sectorID)),
		s.Calculator.MultiplicationFloat64(coords.Z, float64(sectorID)),
		coords.Vel)
}

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

func (s Service) ParseParameter(param string) (float64, error) {
	coordinate, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0, ErrParseParameter
	}
	return coordinate, nil
}

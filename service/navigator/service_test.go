package navigator_test

import (
	"DNS/domain"
	"DNS/logger"
	"DNS/repository/storage"
	"DNS/repository/storage/mock"
	"DNS/service/calculator"
	"DNS/service/navigator"
	"reflect"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
)

func TestService_locate(t *testing.T) {
	tests := []struct {
		name     string
		coords   *navigator.Coordinates
		sectorID int64
		want     float64
	}{
		{
			name: "success",
			coords: &navigator.Coordinates{
				X:   1.1,
				Y:   2.2,
				Z:   3.3,
				Vel: 4.4,
			},
			sectorID: 1,
			want:     11,
		},
		{
			name: "zero values",
			coords: &navigator.Coordinates{
				X:   0,
				Y:   0,
				Z:   0,
				Vel: 0,
			},
			sectorID: 1,
			want:     0,
		},
		{
			name: "one value is zero",
			coords: &navigator.Coordinates{
				X:   1.1111,
				Y:   0,
				Z:   2.2222,
				Vel: 3.3333,
			},
			sectorID: 1,
			want:     6.6666,
		},
		{
			name: "sector id is zero",
			coords: &navigator.Coordinates{
				X:   1.1111,
				Y:   2.2222,
				Z:   3.3333,
				Vel: 4.4444,
			},
			sectorID: 0,
			want:     4.4444,
		},
	}
	s := navigator.Service{
		Logger:     nil,
		Sectors:    nil,
		Calculator: calculator.Service{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.Locate(tt.coords, tt.sectorID); got != tt.want {
				t.Errorf("locate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ConvertCoordinates(t *testing.T) {
	tests := []struct {
		name        string
		coordinates domain.DroneCoordinates
		want        *navigator.Coordinates
		wantErr     bool
	}{
		{
			name: "success",
			coordinates: domain.DroneCoordinates{
				X:   "1.1",
				Y:   "2.2",
				Z:   "3.3",
				Vel: "4.4",
			},
			want: &navigator.Coordinates{
				X:   1.1,
				Y:   2.2,
				Z:   3.3,
				Vel: 4.4,
			},
			wantErr: false,
		},
		{
			name: "convertation failed",
			coordinates: domain.DroneCoordinates{
				X:   "1.1s",
				Y:   "2.2",
				Z:   "3.3",
				Vel: "4.4",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "parameter is empty",
			coordinates: domain.DroneCoordinates{
				X:   "",
				Y:   "2.2",
				Z:   "3.3",
				Vel: "4.4",
			},
			want:    nil,
			wantErr: true,
		},
	}
	s := navigator.Service{
		Logger:     nil,
		Sectors:    nil,
		Calculator: calculator.Service{},
	}
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ConvertCoordinates(tt.coordinates)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertCoordinates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertCoordinates() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_LocateDatabankByCoordinates(t *testing.T) {
	tests := []struct {
		name     string
		sectors  storage.SectorsInterface
		coords   domain.DroneCoordinates
		sectorID domain.SectorID
		want     *domain.DatabankLocation
		wantErr  bool
	}{
		{
			name: "success",
			sectors: mock.Sectors{
				Values: storage.Sector{
					ID:        1,
					SectorID:  1,
					CreatedAt: time.Now(),
					Active:    true,
				},
				Err: nil,
			},
			coords: domain.DroneCoordinates{
				X:   "1.1",
				Y:   "2.2",
				Z:   "3.3",
				Vel: "4.4",
			},
			sectorID: domain.SectorID(1),
			want: func() *domain.DatabankLocation {
				return &domain.DatabankLocation{Location: 11}
			}(),
			wantErr: false,
		},
		{
			name: "sector not found",
			sectors: mock.Sectors{
				Values: storage.Sector{},
				Err:    pgx.ErrNoRows,
			},
			coords: domain.DroneCoordinates{
				X:   "1.1",
				Y:   "2.2",
				Z:   "3.3",
				Vel: "4.4",
			},
			sectorID: domain.SectorID(1),
			want: func() *domain.DatabankLocation {
				return nil
			}(),
			wantErr: true,
		},
		{
			name: "error parse parameter",
			sectors: mock.Sectors{
				Values: storage.Sector{},
				Err:    pgx.ErrNoRows,
			},
			coords: domain.DroneCoordinates{
				X:   "0",
				Y:   "error",
				Z:   "",
				Vel: "",
			},
			sectorID: domain.SectorID(1),
			want: func() *domain.DatabankLocation {
				return nil
			}(),
			wantErr: true,
		},
	}
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := navigator.Service{
				Logger:     logger.Initialize(),
				Sectors:    tt.sectors,
				Calculator: calculator.Initialize(),
				SectorID:   tt.sectorID,
			}

			got, err := s.LocateDatabankByCoordinates(tt.coords)
			if (err != nil) != tt.wantErr {
				t.Errorf("LocateDatabankByCoordinates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocateDatabankByCoordinates() got = %v, want %v", got, tt.want)
			}
		})
	}
}

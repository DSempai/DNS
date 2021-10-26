package api_test

import (
	"DNS/api/v1"
	"DNS/logger"
	"DNS/repository/storage"
	"DNS/repository/storage/mock"
	"DNS/service/calculator"
	"DNS/service/navigator"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

func Test_LocateDatabankHandler(t *testing.T) {
	tests := []struct {
		storage  mock.Sectors
		sectorID int64
		payload  string
		want     []byte
	}{
		{
			storage: mock.Sectors{
				Values: storage.Sector{
					ID:        1,
					SectorID:  1,
					CreatedAt: time.Now(),
					Active:    true,
				},
				Err: nil,
			},
			sectorID: 1,
			payload: `{
				"x": "1.1",
				"y": "2.2",
				"z": "3.3",
				"vel": "4.4"
			}`,
			want: []byte(`{"loc":11}`),
		},
		{
			storage: mock.Sectors{
				Values: storage.Sector{},
				Err:    storage.ErrSectorNotFound,
			},
			sectorID: 1,
			payload: `{
				"x": "1.1",
				"y": "2.2",
				"z": "3.3",
				"vel": "4.4"
			}`,
			want: []byte(`{"status":"failed","code":404,"description":"sector was not found by the provided parameters"}`),
		},
		{
			storage: mock.Sectors{
				Values: storage.Sector{
					ID:        1,
					SectorID:  1,
					CreatedAt: time.Now(),
					Active:    true,
				},
				Err: nil,
			},
			sectorID: 1,
			payload: `{
				"x": "0",
				"y": "0",
				"z": "0",
				"vel": "0"
			}`,
			want: []byte(`{"loc":0}`),
		},
	}
	log := logger.Initialize()
	calc := calculator.Initialize()
	for _, tt := range tests {
		dns := navigator.Initialize(log, tt.storage, calc, tt.sectorID)
		ts := httptest.NewServer(http.HandlerFunc(api.LocateDatabankHandler(log, dns)))
		req, err := http.Post(ts.URL, "application/json", strings.NewReader(tt.payload))
		if err != nil {
			t.Errorf("LocateDatabankHandler() failed. Error: %v", req)
		}
		response, err := io.ReadAll(req.Body)
		if err != nil {
			t.Errorf("LocateDatabankHandler() failed. Error: %v", req)
		}
		if !reflect.DeepEqual(response, tt.want) {
			t.Errorf("LocateDatabankHandler() = %v, want %v", response, tt.want)
		}
		ts.Close()
		if err = req.Body.Close(); err != nil {
			t.Errorf("LocateDatabankHandler() error while close body: %v", err)
		}
	}
}

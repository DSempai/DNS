package api

import (
	"DNS/domain"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func LocateDatabankHandler(logger *logrus.Logger, dns domain.DroneNavigation) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rawBody, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Infof("Error during reading request body with error: %v", err)
		}

		var requestData LocateDatabankRequest
		if err = json.Unmarshal(rawBody, &requestData); err != nil {
			logger.Infof("Error unmarshalling request body: %v", err)
		}

		_, err = dns.LocateDatabankByCoordinates(domain.DroneCoordinates{
			X:   requestData.X,
			Y:   requestData.Y,
			Z:   requestData.Z,
			Vel: requestData.Vel,
		})
		if err != nil {
			logger.Infof("Error during calulate location with error: %v", err)
			//	todo a methods for response marshaler for response too
		}

		w.WriteHeader(http.StatusOK)
	}
}

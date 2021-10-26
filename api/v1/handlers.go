package api

import (
	"DNS/domain"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func LocateDatabankHandler(logger *logrus.Logger, dns domain.DroneNavigation) func(
	w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestData domain.DroneCoordinates
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			logger.Infof("Reading request body failed. Error: %v", err)
			BadRequestResponse(w, err)
			return
		}
		defer r.Body.Close()

		response, err := dns.LocateDatabankByCoordinates(domain.DroneCoordinates{
			X:   requestData.X,
			Y:   requestData.Y,
			Z:   requestData.Z,
			Vel: requestData.Vel,
		})
		if err != nil {
			logger.Infof("Locate databank failed. Error: %v", err)
			NotFoundResponse(w, err)
			return
		}

		OKResponse(w, response)
	}
}

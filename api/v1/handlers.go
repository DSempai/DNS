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
			logger.Infof("Reading request body failed. Error: %v", err)
			BadRequestResponse(w, err)
			return
		}

		var requestData domain.DroneCoordinates
		if err = json.Unmarshal(rawBody, &requestData); err != nil {
			logger.Infof("Unmarshalling request body failed. Error: %v", err)
			BadRequestResponse(w, err)
			return
		}

		response, err := dns.LocateDatabankByCoordinates(domain.DroneCoordinates{
			X:   requestData.X,
			Y:   requestData.Y,
			Z:   requestData.Z,
			Vel: requestData.Vel}, domain.ID)
		if err != nil {
			logger.Infof("Locate databank failed. Error: %v", err)
			NotFoundResponse(w, err)
			return
		}

		OKResponse(w, response)
	}
}

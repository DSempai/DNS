package main

import (
	"DNS/api/v1"
	"DNS/config"
	"DNS/repository/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func bootstrapServer(logger *logrus.Logger, config config.StartupConfig) (*http.Server, func()) {
	s, err := storage.Initialize(logger, config.DatabaseDSN, config.DatabaseMaxConn)
	if err != nil {
		logger.Error("Error creating pg service:", err)
		logger.Error("Exiting")
		os.Exit(1)
	}

	rootRouter := mux.NewRouter()
	routerApi := rootRouter.PathPrefix("/api").Subrouter()
	routerV1 := routerApi.PathPrefix("/v1").Subrouter()

	routerV1.HandleFunc("/drone", api.DroneHandler).Methods(http.MethodGet)
	return &http.Server{
			Addr:    config.ListenAddr,
			Handler: rootRouter,
		}, func() {
			s.Close()
		}
}

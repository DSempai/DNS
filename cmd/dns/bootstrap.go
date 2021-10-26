// Copyright (C) 2218 Atlas Corporation - All Rights Reserved.
// Use of this software without a license will result in an intergalactic government investigation.
// The license can be obtained from the Galactic Government Services branch on Vogsphere.
package main

import (
	"DNS/api/v1"
	"DNS/config"
	"DNS/repository/storage"
	"DNS/service/calculator"
	"DNS/service/navigator"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// bootstrapServer initialize program services, register routes and return *http.Server instance.
func bootstrapServer(logger *logrus.Logger, config config.StartupConfigDNS) (*http.Server, func()) {
	// Initialize service that we need
	storeService, err := storage.Initialize(logger, config.DatabaseDSN, config.DatabaseMaxConn)
	if err != nil {
		logger.Error("Creating storage service fail, Error:", err)
		os.Exit(1)
	}
	calculatorService := calculator.Initialize()
	navigationService := navigator.Initialize(logger, storeService, calculatorService, config.SectorID)

	// Create routers/subrouters
	rootRouter := mux.NewRouter()
	routerAPI := rootRouter.PathPrefix("/api").Subrouter()
	routerV1 := routerAPI.PathPrefix("/v1").Subrouter()

	// Register all handlers
	handlersV1 := []Handlers{
		{
			URL:     "/locate_databank",
			Func:    api.LocateDatabankHandler(logger, navigationService),
			Methods: []string{http.MethodPost},
		},
	}

	for _, handler := range handlersV1 {
		routerV1.HandleFunc(handler.URL, handler.Func).Methods(handler.Methods...)
	}

	return &http.Server{
			Addr:    config.ListenAddr,
			Handler: rootRouter,
		}, func() {
			storeService.Close()
		}
}

// Handlers created and used to register new paths and handler functions with specified methods.
type Handlers struct {
	URL     string
	Func    func(http.ResponseWriter, *http.Request)
	Methods []string
}

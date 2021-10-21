package main

import (
	"DNS/api/v1"
	"DNS/config"
	"DNS/repository/storage"
	"DNS/service/calculator"
	"DNS/service/navigator"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type Handlers struct {
	URL     string
	Func    func(http.ResponseWriter, *http.Request)
	Methods []string
}

func bootstrapServer(logger *logrus.Logger, config config.StartupConfig) (*http.Server, func()) {
	store, err := storage.Initialize(logger, config.DatabaseDSN, config.DatabaseMaxConn)
	if err != nil {
		logger.Error("Error creating storage service:", err)
		os.Exit(1)
	}
	calculatorService := calculator.Initialize(logger)
	nabigationService := navigator.Initialize(logger, store, calculatorService)

	rootRouter := mux.NewRouter()
	routerApi := rootRouter.PathPrefix("/api").Subrouter()
	routerV1 := routerApi.PathPrefix("/v1").Subrouter()

	handlersV1 := []Handlers{
		{
			URL:     "/locate_databank",
			Func:    api.LocateDatabankHandler(logger, nabigationService),
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
			store.Close()
		}
}

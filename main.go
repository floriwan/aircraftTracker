package main

import (
	"aircraftTracker/actrack"
	"aircraftTracker/config"
	"aircraftTracker/handler"
	"aircraftTracker/observer"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	configFile = "config.yaml"
)

func main() {
	var myConfig config.Config
	err := myConfig.ReadConfig(configFile)
	if err != nil {
		fmt.Errorf("can not open config file", err)
	}

	if err := observer.Init(myConfig); err != nil {
		log.Fatal(err)
	}

	actrack.InitAircraftData(myConfig)
	go actrack.StartDataUpdater(myConfig)

	r := mux.NewRouter()
	r.HandleFunc("/", handler.Home)
	r.HandleFunc("/data/{reg}", handler.GetAircraftData).Methods("GET")
	r.HandleFunc("/reg/{reg}", handler.HandleAircraftReg)
	r.HandleFunc("/search", handler.SearchAircraft).Methods(("GET"))

	log.Fatal(http.ListenAndServe(":8080", r))

}

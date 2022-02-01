package main

import (
	"aircraftTracker/actrack"
	"aircraftTracker/config"
	"aircraftTracker/handler"
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

	actrack.InitAircraftData(myConfig)
	go actrack.StartDataUpdater(myConfig)

	r := mux.NewRouter()
	r.HandleFunc("/", handler.Home)
	r.HandleFunc("/data/{reg}", handler.GetAircraftData).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

}

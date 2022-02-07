package main

import (
	"aircraftTracker/actrack"
	"aircraftTracker/config"
	"aircraftTracker/handler"
	"aircraftTracker/observer"
	"flag"
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
	var discordMode bool

	flag.BoolVar(&discordMode, "d", false, "update discord channel")
	flag.Parse()

	err := myConfig.ReadConfig(configFile)
	if err != nil {
		fmt.Errorf("can not open config file", err)
	}

	// import aircarft registration database and start data updater
	actrack.InitAircraftData(myConfig)
	go actrack.StartDataUpdater(myConfig)

	// start observer
	if err := observer.Init(myConfig); err != nil {
		log.Fatal(err)
	}
	defer observer.Close()

	// in discord mode
	if discordMode {
		observer.Start(myConfig)
		defer observer.Dbot.Close()
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handler.Home)
	r.HandleFunc("/data/{reg}", handler.GetAircraftData).Methods("GET")
	r.HandleFunc("/reg/{reg}", handler.HandleAircraftReg)
	r.HandleFunc("/search", handler.SearchAircraft).Methods(("GET"))
	r.HandleFunc("/stats", handler.GetStatistics).Methods(("GET"))
	r.HandleFunc("/list", handler.GetObservationList).Methods(("GET"))

	log.Fatal(http.ListenAndServe(":8080", r))

}

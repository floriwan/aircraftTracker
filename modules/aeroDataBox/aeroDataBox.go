package aeroDataBox

import (
	httphelper "aircraftTracker/modules/http"
	"aircraftTracker/types/aerodatabox"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

var FlightStatusResult []aerodatabox.FlightStatus

func RequestData(reg string) (data *aerodatabox.Aircraft, err error) {

	url := fmt.Sprintf("https://aerodatabox.p.rapidapi.com/flights/reg/%v?withAircraftImage=false&withLocation=false", reg)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", viper.GetString("aerodatabox.rapidapikey"))
	req.Header.Add("X-RapidAPI-Host", viper.GetString("aerodatabox.rapidapihost"))

	b, err := httphelper.SendRequest(req)
	if err != nil {
		return nil, err
	}

	data = &aerodatabox.Aircraft{}
	if err := json.Unmarshal(b, &data); err != nil {
		log.Printf("can not unmarshal %v\n%v\n", string(b), err)
		return nil, err
	}

	return data, nil
}

/*
url := fmt.Sprintf("https://aerodatabox.p.rapidapi.com/aircrafts/reg/%v?withImage=true", reg)

	req, _ := http.NewRequest("GET", url, nil)



*/

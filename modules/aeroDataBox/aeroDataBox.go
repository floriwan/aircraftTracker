package aerodatabox

import (
	"aircraftTracker/types/aerodatabox"
)

// var rapidApiKey string
// var rapidApiHost string
var FlightStatusResult []aerodatabox.FlightStatus

// func init() {
// 	rapidApiKey = viper.GetString("aerodatabox.rapidapikey")
// 	rapidApiHost = viper.GetString("aerodatabox.rapidapihost")
// }

func RequestData(reg string) error {

	// url := fmt.Sprintf("https://aerodatabox.p.rapidapi.com/flights/reg/%v?withAircraftImage=false&withLocation=false", reg)
	// req, _ := http.NewRequest("GET", url, nil)
	// req.Header.Add("X-RapidAPI-Key", rapidApiKey)
	// req.Header.Add("X-RapidAPI-Host", rapidApiHost)

	// b, err := httphelper.SendRequest(req)
	// if err != nil {
	// 	return nil, err
	// }

	return nil
}

/*
url := fmt.Sprintf("https://aerodatabox.p.rapidapi.com/aircrafts/reg/%v?withImage=true", reg)

	req, _ := http.NewRequest("GET", url, nil)



*/

package adsbexchange

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	httphelper "aircraftTracker/modules/http"
	"aircraftTracker/types/adsbexchange"

	"github.com/spf13/viper"
)

func RequestData(reg string) (data *adsbexchange.AircraftPositionByRegistration, err error) {
	url := fmt.Sprintf("https://adsbexchange-com1.p.rapidapi.com/v2/registration/%v/", reg)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", viper.GetString("adsbexchange.rapidapikey"))
	req.Header.Add("X-RapidAPI-Host", viper.GetString("adsbexchange.rapidapihost"))

	b, err := httphelper.SendRequest(req)
	if err != nil {
		return nil, err
	}

	data = &adsbexchange.AircraftPositionByRegistration{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Printf("can not unmarshal %v\n%v\n", string(b), err)
		return nil, err
	}

	return data, nil
}

func IsDataNew(current adsbexchange.AircraftPositionByRegistration,
	old adsbexchange.AircraftPositionByRegistration) bool {
	if len(current.Ac) != len(old.Ac) {
		return false
	}

	if old.Ac[0].Flight == current.Ac[0].Flight &&
		old.Ac[0].Squawk == current.Ac[0].Squawk {

		// landing
		if old.Ac[0].Alt_geom != 0 && current.Ac[0].Alt_geom == 0 {
			return true
		}

		// take off
		if old.Ac[0].Alt_geom == 0 && current.Ac[0].Alt_geom != 0 {
			return true
		}

		// parking
		if old.Ac[0].Gs != 0 && current.Ac[0].Gs == 0 {
			return true
		}

		// taxing
		if old.Ac[0].Gs == 0 && current.Ac[0].Gs != 0 {
			return true
		}
	}

	return false
}

// return true, if flight state between both adsbexchange data fields is different
func FlightStateChanged(d1 adsbexchange.AircraftPositionByRegistration, d2 adsbexchange.AircraftPositionByRegistration) bool {
	return GetFlightState(d1) != GetFlightState(d2)
}

// return flight state of the current adsbexchange data.
// gs == 0 and alt != 0 => there is something wrong :-)
// gs != 0 and alt == 0 => plane is moving
// gs != 0 and alt =! 0 => plane is in the air
// gs == 0 and alt == 0 => plane is parking somewhere
func GetFlightState(data adsbexchange.AircraftPositionByRegistration) adsbexchange.FlightState {

	//log.Printf("flight state gs: %v, alt geom: %v", data.Ac[0].Alt_geom, data.Ac[0].Gs)
	if len(data.Ac) == 0 {
		return adsbexchange.Unknown
	}

	if data.Ac[0].Gs != 0 && data.Ac[0].Alt_geom == 0 {
		return adsbexchange.Taxing
	}

	if data.Ac[0].Gs != 0 && data.Ac[0].Alt_geom != 0 {
		return adsbexchange.Enroute
	}

	if data.Ac[0].Gs == 0 && data.Ac[0].Alt_geom == 0 {
		return adsbexchange.Parking
	}

	return adsbexchange.Unknown
}

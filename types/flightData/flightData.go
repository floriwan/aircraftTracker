package flightdata

import (
	"aircraftTracker/types/adsbexchange"
	"time"
)

type FlightData struct {
	AdsbExchangeData adsbexchange.AircraftPositionByRegistration
	OtherData        string // some other data aero data box??
	Timestamp        time.Time
}

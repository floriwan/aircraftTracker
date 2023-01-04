package aerodatabox

type FlightStatus struct {
	FlightDeparture Departure     `json:"departure"`
	FlightArrival   Arrival       `json:"arrival"`
	FlightStatus    string        `json:"status"`
	FlightAircraft  AircraftShort `json:"aircraft"`
}

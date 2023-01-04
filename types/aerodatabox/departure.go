package aerodatabox

type Departure struct {
	DepartureAirport Airport `json:"airport"`
	ScheduledTimeUtc string  `json:"scheduledTimeUtc"`
}

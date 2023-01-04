package aerodatabox

type Arrival struct {
	ArrivalAirport   Airport `json:"airport"`
	ScheduledTimeUtc string  `json:"scheduledTimeUtc"`
}

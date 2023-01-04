package adsbexchange

// AdsbExchangeData is the adsbexchange response data structure.
// If aircraft was found, the array ac will contain a single ac structure.
// If there are no information for the registration available, the array len is 0.
// ?? I have never seen more than one element in the array ??
type AircraftPositionByRegistration struct {
	Ac    []Ac `json:"ac"`
	Total int  `json:"total"`
	Ctime int  `json:"ctime"`
	Now   int  `json:"now"`
	Ptime int  `json:"ptime"`
}

package adsbexchange

//go:generate go run golang.org/x/tools/cmd/stringer -type=FlightState
type FlightState int

/*
go generate types/adsbexchange/FlightState
*/
const (
	Unknown FlightState = iota
	Parking
	Taxing
	Enroute
)

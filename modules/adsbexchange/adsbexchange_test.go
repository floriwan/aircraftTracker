package adsbexchange

import (
	"aircraftTracker/types/adsbexchange"
	"testing"
)

func TestEquals(t *testing.T) {

	current := adsbexchange.AircraftPositionByRegistration{
		Ac: []adsbexchange.Ac{{Flight: "DLH123",
			Squawk:   "1000",
			Gs:       45,
			Alt_geom: 3000}},
	}
	old := adsbexchange.AircraftPositionByRegistration{
		Ac: []adsbexchange.Ac{{Flight: "DLH123",
			Squawk:   "1000",
			Gs:       45,
			Alt_geom: 3000}},
	}

	if IsDataNew(current, old) {
		t.Errorf("data should be the same")
	}

	current.Ac[0].Gs = 0
	if !IsDataNew(current, old) {
		t.Errorf("groundspeed change, data should not be equal")
	}

	current.Ac[0].Gs = 45
	current.Ac[0].Alt_geom = 0
	if !IsDataNew(current, old) {
		t.Errorf("altitude change, data should not be equal")
	}

	current.Ac[0].Alt_geom = 3000
	old.Ac[0].Alt_geom = 0
	if !IsDataNew(current, old) {
		t.Errorf("altitude change, data should not be equal")
	}
}

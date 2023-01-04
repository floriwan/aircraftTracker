package flightDataHistory

import (
	adsbexchangetypes "aircraftTracker/types/adsbexchange"
	"testing"
)

func TestAddToHistory(t *testing.T) {

	data := adsbexchangetypes.AircraftPositionByRegistration{}

	err := AddToHistory("FLORIAN", data)
	if err != nil {
		t.Fatalf("error should not be present %v", err)
	}

	err = AddToHistory("FLORIAN", data)
	if err != nil {
		t.Fatalf("second error should not be present %v", err)
	}

	if HistorySize("FLORIAN") != 2 {
		t.Fatalf("history size should be 2 but is %v", HistorySize("FLORIAN"))
	}
}

// func TestHistorySize(t *testing.T) {

// 	data := adsbexchangetypes.AircraftPositionByRegistration{}
// 	if err := AddToHistory("FLORIAN", data); err != nil {
// 		t.Fatalf("can not add data number data to history")
// 	}

// 	maxHistorySize = 5

// 	for i := 0; i < 10; i++ {
// 		if err := AddToHistory("FLORIAN", data); err != nil {
// 			t.Fatalf("can not add data number %v to history", i)
// 		}
// 	}

// 	if HistorySize("FLORIAN") != 5 {
// 		t.Fatalf("history size should be 5 but is %v", HistorySize("FLORIAN"))
// 	}

// }

func TestFlightStateChanges(t *testing.T) {
	data := adsbexchangetypes.AircraftPositionByRegistration{}
	data.Ac = append(data.Ac, adsbexchangetypes.Ac{Gs: 0, Alt_geom: 0})

	if err := AddToHistory("FLORIAN", data); err != nil {
		t.Fatalf("can not add data number data to history")
	}

	if FlightStateChanged("FLORIAN") == true {
		t.Fatalf("not enough data in history array, return should be false but was %v", FlightStateChanged("FLORIAN"))
	}

	if err := AddToHistory("FLORIAN", data); err != nil {
		t.Fatalf("can not add data number data to history")
	}

	if FlightStateChanged("FLORIAN") == true {
		t.Fatalf("no changes in data %v", FlightStateChanged("FLORIAN"))
	}

	// change the gs
	data.Ac = []adsbexchangetypes.Ac{{Gs: 10, Alt_geom: 0}}
	if err := AddToHistory("FLORIAN", data); err != nil {
		t.Fatalf("can not add data number data to history")
	}

	if FlightStateChanged("FLORIAN") == false {
		t.Fatalf("flight data should be changed %v", FlightStateChanged("FLORIAN"))
	}

	// change the gs
	data.Ac = []adsbexchangetypes.Ac{{Gs: 0, Alt_geom: 0}}
	if err := AddToHistory("FLORIAN", data); err != nil {
		t.Fatalf("can not add data number data to history")
	}

	if FlightStateChanged("FLORIAN") == false {
		t.Fatalf("flight data should be changed %v", FlightStateChanged("FLORIAN"))
	}

	// change gs and alt
	data.Ac = []adsbexchangetypes.Ac{{Gs: 100, Alt_geom: 100}}
	if err := AddToHistory("FLORIAN", data); err != nil {
		t.Fatalf("can not add data number data to history")
	}

	if FlightStateChanged("FLORIAN") == false {
		t.Fatalf("flight data should be changed %v", FlightStateChanged("FLORIAN"))
	}

}

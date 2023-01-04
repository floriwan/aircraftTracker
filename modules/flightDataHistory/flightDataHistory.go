package flightDataHistory

import (
	"aircraftTracker/modules/adsbexchange"
	adsbexchangetypes "aircraftTracker/types/adsbexchange"
	flightdata "aircraftTracker/types/flightData"
	"fmt"
	"log"
	"time"
)

type historyEntry struct {
	flightDataEntry []flightdata.FlightData
}

var history map[string]*historyEntry

//var maxHistorySize int

func init() {
	history = make(map[string]*historyEntry)
	//maxHistorySize = viper.GetInt("aircraftdatahistory.maxhistorysize")
}

// if there is a flight state change, add data to data history array
func AddToHistory(reg string, data adsbexchangetypes.AircraftPositionByRegistration) error {

	if !isRegistrationInHistory(reg) {
		history[reg] = &historyEntry{}
	}

	// this is the first entry
	if HistorySize(reg) == 0 {
		history[reg].flightDataEntry = append(history[reg].flightDataEntry, flightdata.FlightData{
			AdsbExchangeData: data,
			OtherData:        "lalelu",
			Timestamp:        time.Now(),
		})
		return nil
	}

	// add new data element only, if flight state was changed
	last := history[reg].flightDataEntry[len(history[reg].flightDataEntry)-1].AdsbExchangeData
	log.Printf("%v flight state %v > %v ", reg, adsbexchange.GetFlightState(last), adsbexchange.GetFlightState(data))
	if !adsbexchange.FlightStateChanged(last, data) {
		//log.Printf(" -> no state change\n")
		return nil
	}

	log.Printf(" -> add new data, history size: %v\n", HistorySize(reg))
	history[reg].flightDataEntry = append(history[reg].flightDataEntry, flightdata.FlightData{
		AdsbExchangeData: data,
		OtherData:        "lalelu",
		Timestamp:        time.Now(),
	})

	// TODO check history size and remove element too old from list

	return nil

	// history[reg].flightDataEntry = append(history[reg].flightDataEntry, flightdata.FlightData{
	// 	AdsbExchangeData: data,
	// 	OtherData:        "lalelu",
	// })

	// // this is the first entry
	// if HistorySize(reg) == 0{

	// }

	// // size of 0 means unlimted history size
	// // but if the size of the history array is larger than maxHistorySize,
	// // remove the first elements of the array
	// if maxHistorySize != 0 && HistorySize(reg) > maxHistorySize {
	// 	removeIndex := HistorySize(reg) - maxHistorySize
	// 	history[reg].flightDataEntry = history[reg].flightDataEntry[removeIndex:]
	// }

	// log.Printf("%v flight state: %v, history size: %v\n", reg, adsbexchange.GetFlightState(data), HistorySize(reg))
	// return nil

}

func isRegistrationInHistory(reg string) bool {
	_, ok := history[reg]
	return ok
}

// If flight state changed between the last two data entries in
// history, return true, false otherwise.
func FlightStateChanged(reg string) bool {

	// not enough data for history changes
	if HistorySize(reg) < 2 {
		return false
	}

	last := history[reg].flightDataEntry[len(history[reg].flightDataEntry)-1].AdsbExchangeData
	prev := history[reg].flightDataEntry[len(history[reg].flightDataEntry)-2].AdsbExchangeData

	log.Printf("flight state changed: %v -> %v\n", adsbexchange.GetFlightState(prev), adsbexchange.GetFlightState(last))

	return adsbexchange.GetFlightState(last) != adsbexchange.GetFlightState(prev)
}

// Check if registration is in flight data history and return the latest flight information
func GetLastData(reg string) (data *flightdata.FlightData, err error) {
	if !isRegistrationInHistory(reg) {
		return nil, fmt.Errorf("registration %v not in flight data history", reg)
	}
	return &history[reg].flightDataEntry[0], nil
}

// return the size of the aircraft history array
func HistorySize(reg string) int {
	return len(history[reg].flightDataEntry)
}

func DumpData(reg string) {

}

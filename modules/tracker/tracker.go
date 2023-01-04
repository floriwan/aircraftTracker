package tracker

import (
	"aircraftTracker/modules/adsbexchange"
	"aircraftTracker/modules/aeroDataBox"
	"aircraftTracker/modules/aircraftDatabase"
	"aircraftTracker/modules/flightDataHistory"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type trackedAircraft struct {
	registration string
	stopChannel  chan bool
}

var errorUpdateInterval int = 3600 // 60 minutes

// send message via discord hook if flight state changes was detected
// discord webhook id and token must be set in config.json
var UpdateDiscordChannel bool = false

var trackedAircraftList map[string]trackedAircraft

func init() {
	trackedAircraftList = make(map[string]trackedAircraft)
}

func AddRegistration(reg string) error {

	if !aircraftDatabase.IsRegistrationValid(reg) {
		return fmt.Errorf("aircraft registration %v not found in registration database", reg)
	}

	_, ok := trackedAircraftList[reg]
	if ok {
		return fmt.Errorf("aircraft %v already registered in update list", reg)
	}
	go addAircraftTracker(reg)
	return nil
}

func RemoveRegistration(reg string) error {
	_, ok := trackedAircraftList[reg]
	if !ok {
		return fmt.Errorf("aircraft %v not found in registration list", reg)
	}
	removeAircraftTracker(reg)
	return nil
}

func removeAircraftTracker(reg string) {
	trackedAircraftList[reg].stopChannel <- true
	delete(trackedAircraftList, reg)
}

func addAircraftTracker(reg string) {

	stopChannel := make(chan bool)

	trackedAircraftList[reg] = trackedAircraft{
		registration: reg,
		stopChannel:  stopChannel,
	}

	updateInterval := viper.GetInt("tracker.updateinterval")
	ticker := time.NewTicker(time.Duration(updateInterval * int(time.Second)))
	defer ticker.Stop()
	lastTick := time.Now()

	for {
		select {
		case t := <-ticker.C:
			if err := updateAircraft(reg); err != nil {
				// if adsbexchange requests failed, increase update interval to 30 minutes
				log.Printf("get aircraft update error, increase ticker time to %v sec %v\n", errorUpdateInterval, err)
				ticker.Reset(time.Duration(errorUpdateInterval * int(time.Second)))
			}
			if t.Sub(lastTick) > time.Duration(errorUpdateInterval*int(time.Second)) {
				log.Printf("reset ticker back to update interval of %v sec\n", updateInterval)
				ticker.Reset(time.Duration(updateInterval * int(time.Second)))
			}
		case <-stopChannel:
			log.Printf("stopping update for aircraft %v", reg)
			return
		}
	}
}

// update aircraft data
// request adsbexchange and process data to data history
// TODO add additional data from databox
// TODO send discord message of new flight state detected
func updateAircraft(reg string) error {
	data, err := adsbexchange.RequestData(reg)
	if err != nil {
		return fmt.Errorf("adsbexchange error for aircraft %v %v", reg, err)
	}

	if len(data.Ac) == 0 {
		log.Printf("no adsb exchange data for %v found\n", reg)
		return nil
	}

	if err := flightDataHistory.AddToHistory(reg, *data); err != nil {
		return err
	}

	// this is not the initial start and the first entry
	// and no flight state change, nothing todo
	if flightDataHistory.HistorySize(reg) > 1 && !flightDataHistory.FlightStateChanged(reg) {
		return nil
	}

	dbox, err := aeroDataBox.RequestData(reg)
	if err != nil {
		log.Printf("aerodatabox request error: %v\n", err)
	}

	log.Printf("aerodatbox data: %v", dbox)

	sendDiscordMessage(reg)

	return nil
}

func sendDiscordMessage(reg string) {
	/*TODO
	data, err := flightDataHistory.GetLastData(reg)
	if err != nil {
		return
	}*/
	fmt.Printf("send message to discord channel for %v\n", reg)

}

package observer

import (
	"aircraftTracker/acdb"
	"aircraftTracker/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	flightaware "github.com/floriwan/flightaware/request"
)

type observer struct {
	Reg      string
	Interval int
}

var regList []observer
var interval int
var observerFile string
var quitChannels map[string]chan bool
var flightData map[string]flightaware.Flights

func Init(config config.Config) error {

	interval, _ = strconv.Atoi(config.ObserverInterval)
	observerFile = config.ObserverFile
	quitChannels = make(map[string]chan bool)
	flightData = make(map[string]flightaware.Flights)

	// import observation list from file
	if err := readList(); err != nil {
		return err
	}

	log.Printf("observation list imported with %v aircrafts", len(regList))
	return nil
}

func GetSize() int {
	return len(regList)
}

func stopObserver(reg string) {
	quitChannels[reg] <- true
	delete(quitChannels, reg)
}

func removeReg(slice []observer, i int) []observer {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func startObserver(reg string, interval int) {

	// create quit channel in map
	quit := make(chan bool)
	quitChannels[reg] = quit

	// time trigger channel
	updateInterval := time.NewTicker((time.Duration(interval)) * time.Minute)
	log.Printf("start aircraft '%v' observer every %v minutes\n", reg, interval)

	for {
		select {
		case <-quit:
			log.Printf("stop observer for registration '%v'", reg)
			return
		case <-updateInterval.C:
			log.Printf("update information for registration '%v'", reg)
			flights := flightaware.FlightInfo(reg, "", true)
			if len(flights.Flights) == 0 {
				log.Printf("no flight information for aircraft '%v'", reg)
				// remove information from flight map
				delete(flightData, reg)
				continue
			}

			log.Printf("flight information for '%v' %v>%v", reg, flights.Flights[0].Origin.Code, flights.Flights[0].Destination.Code)
			flightData[reg] = flights
			//default:

		}
	}

}

func readList() error {
	log.Printf("import observer list '%v'\n", observerFile)
	b, err := ioutil.ReadFile(observerFile)
	if err != nil {
		log.Printf("can not import observer list '%v' %v", observerFile, err)
		return nil
	}

	err = json.Unmarshal(b, &regList)
	if err != nil {
		return fmt.Errorf("can not unmarshal observer list %v", err)
	}

	// start new observer for all registration
	for k := range regList {
		go startObserver(regList[k].Reg, regList[k].Interval)
	}
	return nil
}

func writeList() error {

	b, err := json.Marshal(regList)
	if err != nil {
		return fmt.Errorf("unable to marshal observer list %v", err)
	}

	err = ioutil.WriteFile(observerFile, b, 0644)
	if err != nil {
		log.Printf("can not export observer list to file '%v' %v\n", observerFile, err)
	}
	return nil
}

func Remove(reg string) error {

	// search registration in registration list
	toRemove := -1
	for k, v := range regList {
		if v.Reg == reg {
			toRemove = k
			break
		}
	}

	if toRemove == -1 {
		return fmt.Errorf("registration '%v' not found in observer list", reg)
	}

	// stop observation goroutine
	log.Printf("remove registration '%v' from observer list", reg)
	stopObserver(reg)

	// remove registration from registration list and save to file
	regList = removeReg(regList, toRemove)
	writeList()

	// remove information from flight map
	delete(flightData, reg)

	return nil
}

func Add(reg string) error {

	if !acdb.IsRegValid(reg) {
		return fmt.Errorf("unknown aircraft registration '%v'", reg)
	}

	if isRegObserved(reg) {
		return nil
	}

	log.Printf("add registration '%v' to observer list", reg)
	regList = append(regList, observer{Reg: reg, Interval: interval})

	if err := writeList(); err != nil {
		return err
	}

	// start the observer
	go startObserver(reg, interval)

	return nil
}

func isRegObserved(reg string) bool {
	for k := range regList {
		if regList[k].Reg == reg {
			return true
		}
	}
	return false
}

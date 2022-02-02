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
)

type observer struct {
	Reg      string
	Interval int
}

var regList []observer
var interval int

const observerFile = "test"

func Init(config config.Config) error {

	interval, _ = strconv.Atoi(config.ObserverInterval)

	if err := readList(); err != nil {
		return err
	}

	log.Printf("observation list imported with %v aircrafts", len(regList))
	return nil
}

func startObserver(reg string, interval int) {

	updateInterval := time.NewTicker((time.Duration(interval)) * time.Minute)
	log.Printf("update aircraft position every %v minutes\n", interval)

	for {
		log.Printf("update information for registration '%v'", reg)
		<-updateInterval.C
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

	go startObserver(reg, interval)

	// TODO start new update method

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

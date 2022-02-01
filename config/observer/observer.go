package observer

import (
	"aircraftTracker/acdb"
	"aircraftTracker/config"
	"fmt"
	"log"
	"strconv"
)

type observer struct {
	reg      string
	interval int
}

var regList []observer
var interval int

func Init(config config.Config) error {
	// TODO read observer list
	interval, _ = strconv.Atoi(config.ObserverInterval)
	return nil
}

func Add(reg string) error {

	if !acdb.IsRegValid(reg) {
		return fmt.Errorf("unknown aircraft registration '%v'", reg)
	}
	// TODO reg already registered?
	log.Printf("add registration '%v' to observer list", reg)
	regList = append(regList, observer{reg: reg, interval: interval})

	// TODO start new update method

	return nil
}

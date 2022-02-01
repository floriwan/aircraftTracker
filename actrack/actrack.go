package actrack

import (
	"aircraftTracker/acdb"
	"aircraftTracker/config"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	acFilename = "ac-db"
)

func InitAircraftData(myConfig config.Config) {

	if myConfig.UpdateAircraftDataInterval < 1 || len(myConfig.AircaftData) < 1 {
		log.Fatalf("invalid configuration %+v", myConfig)
	}

	if err := updateAircraftData(myConfig.AircaftData); err != nil {
		log.Fatalf(err.Error())
	}

}

func StartDataUpdater(myConfig config.Config) {
	acDataUpdateTicker := time.NewTicker((24 * time.Duration(myConfig.UpdateAircraftDataInterval)) * time.Hour)
	log.Printf("update aircarft data every %v hours\n", 24*myConfig.UpdateAircraftDataInterval)

	for {
		updateAircraftData(myConfig.AircaftData)
		<-acDataUpdateTicker.C
	}
}

func updateAircraftData(url string) error {
	if err := downloadData(url); err != nil {
		return err
	}

	f, err := os.Open(acFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	gr, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer gr.Close()
	acdb.Import(gr)

	return nil
}

func downloadData(url string) error {

	if !updateNeeded(acFilename) {
		return nil
	}

	log.Printf("downloading aircraft data: " + url)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("can not download aircarft data %v\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Create the file
	out, err := os.Create(acFilename)
	if err != nil {
		return err
	}
	defer out.Close()

	// save to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func updateNeeded(filename string) bool {
	acdb, err := os.Stat(acFilename)
	if err != nil {
		log.Printf("unable to get file statistics for '%v'", acFilename)
		return true
	}

	today := time.Now()
	yesterday := today.Add(-24 * time.Hour)

	if yesterday.Before(acdb.ModTime()) {
		log.Printf("no aircraft data update nedded")
		return false
	}

	return true

}

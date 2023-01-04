package aircraftDatabase

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// The Aircraft holds all information about a single aircraft from adsbexchange
type Aircraft struct {
	Icao         string `json:"icao"`
	Reg          string `json:"reg"`
	Icaotype     string `json:"icaotype"`
	Year         string `json:"year"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Ownop        string `json:"ownop"`
	Faa_pia      bool   `json:"faa_pia"`
	Faa_ladd     bool   `json:"faa_ladd"`
	Short_type   string `json:"short_type"`
	Mil          bool   `json:"mil"`
}

// This map holds all aircraft information with the registration as key for faster access.
var aircrafts map[string]Aircraft

func init() {
	aircrafts = make(map[string]Aircraft)
}

// Setup the aircraft registration database.
// Load database file from adsbexchange server,
// store file on disk and extract the zip file
func Setup() error {

	databaseFilename := viper.GetString("aircraftdatabase.databasefilename")
	databaseDownload := viper.GetString("aircraftdatabase.basicaircraftdatabasedownload")

	go startUpdate(databaseDownload, databaseFilename)

	return nil
}

// Returns true if the registration is in the list of aircrafts
func IsRegistrationValid(reg string) bool {
	_, ok := aircrafts[reg]
	return ok
}

// startUpdate will initial load aircraft registration database
// and update every second day
func startUpdate(url string, filename string) error {
	ticker := time.NewTicker(time.Duration(48 * int(time.Hour))) // 2 days
	defer ticker.Stop()
	for ; true; <-ticker.C {
		if err := download(url, filename); err != nil {
			return err
		}

		b, err := readData(filename)
		if err != nil {
			return err
		}

		if err := importData(b); err != nil {
			return err
		}
	}
	return nil
}

func importData(b []byte) error {
	a := strings.Split(string(b), "}")
	for _, v := range a {
		if len(v) == 1 {
			continue
		}
		// add the spit character again to every line
		v = v + "}"
		//a := Aircraft{Year: "N/A", Manufacturer: "N/A", Model: "N/A", Ownop: "N/A"}
		a := Aircraft{}
		if err := json.Unmarshal([]byte(v), &a); err != nil {
			return fmt.Errorf("can not unmashal aircraft line %v %v", v, err)
		}
		aircrafts[a.Reg] = a
	}
	log.Printf("imported %v aircraft registrations\n", len(aircrafts))
	return nil
}

func readData(filename string) (b []byte, err error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	fz, err := gzip.NewReader(fi)
	if err != nil {
		return nil, err
	}
	defer fz.Close()

	s, err := io.ReadAll(fz)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func download(url string, filename string) error {

	if !isFileOlderThan1Day(filename) {
		return nil
	}

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	log.Printf("download aircraft database: %v\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad response status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func isFileOlderThan1Day(filename string) bool {

	info, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return true
	}

	today := time.Now()
	yesterday := today.Add(-24 * time.Hour)
	return yesterday.After(info.ModTime())

}

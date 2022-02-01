package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AircaftData                             string `yaml:"aircaftData"`
	UpdateAircraftDataInterval              int    `yaml:"updateAircraftDataInterval"`
	UpdateAircraftFlightInformationInterval int    `yaml:"updateAircraftFlightInformationInterval"`
	AircraftRegistrations                   string `yaml:"aircraftRegistrations"`
	ObserverInterval                        string `yaml:"observerInterval"`
}

func (c Config) Print() string {
	return fmt.Sprintf("%+v", c)
}

func (c *Config) ReadConfig(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return err
	}

	return nil

}

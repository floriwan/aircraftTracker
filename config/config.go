package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Config struct {
	AircaftData                             string `yaml:"aircaftData"`
	UpdateAircraftDataInterval              int    `yaml:"updateAircraftDataInterval"`
	UpdateAircraftFlightInformationInterval int    `yaml:"updateAircraftFlightInformationInterval"`
	AircraftRegistrations                   string `yaml:"aircraftRegistrations"`
	ObserverInterval                        string `yaml:"observerInterval"`
	ObserverFile                            string `yaml:"observerFile"`
	AeroApiKey                              string
	DiscordToken                            string
	DiscordBotPrefix                        string
	DiscordWebHook                          string
}

func (c Config) Print() string {
	return fmt.Sprintf("%+v", c)
}

func (c *Config) ReadConfig(filename string) error {

	c.readEnv()

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

func (c *Config) readEnv() {
	godotenv.Load()
	c.AeroApiKey = os.Getenv("AERO_API_KEY")
	c.DiscordToken = os.Getenv("DISCORD_TOKEN")
	c.DiscordBotPrefix = os.Getenv("DISCORD_BOT_PREFIX")
	c.DiscordWebHook = os.Getenv("DISCORD_WEBHOOK")
}

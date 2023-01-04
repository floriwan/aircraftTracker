package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version = "0.0.1"
var configFile string
var verbose bool

var rootCmd = &cobra.Command{
	Version: version,
	Use:     "aircraftTracker",
	Short:   "track you favorite arcrafts",
	Long:    `Aircraft tracker to track you favorite aircrafts`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "config/config.json", "config file (default is config/config.yml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

func initConfig() {
	viper.SetConfigFile(configFile)

	// read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("can not read config file %v : %v", viper.ConfigFileUsed(), err)
	}

}

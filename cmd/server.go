package cmd

import (
	aircraftdatabase "aircraftTracker/modules/aircraftDatabase"
	"aircraftTracker/modules/restApi/aircraft"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var port int
var updateDiscord bool

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start aircraft tracker",
	Long:  `Start aircraft tracker server with rest interface`,
	Run:   runServerCmd,
}

func runServerCmd(cmd *cobra.Command, args []string) {
	if err := aircraftdatabase.Setup(); err != nil {
		log.Fatalf("can not read aircraft registration file %v", err)
	}
	startRestService()
}

func startRestService() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes := router.Group("/aircraft")
	routes.POST("/add/:registration", aircraft.Add)
	routes.POST("/remove/:registration", aircraft.Remove)
	routes.POST("/status/:registration", aircraft.Status)
	router.Run(fmt.Sprintf(":%v", viper.GetInt("rest.port")))
}

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "p", 8000, "rest api port")
	serverCmd.Flags().BoolVarP(&updateDiscord, "discord", "d", false, "on flight state changes, send message to discord channel (via webhooks)")
	rootCmd.AddCommand(serverCmd)
}

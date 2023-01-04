package aircraft

import (
	"aircraftTracker/modules/tracker"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	reg := c.Param("registration")
	if err := tracker.AddRegistration(reg); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{reg: "aircraft registered"})
}

package aircraft

import (
	"aircraftTracker/modules/tracker"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Remove(c *gin.Context) {
	reg := c.Param("registration")
	if err := tracker.RemoveRegistration(reg); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{reg: "aircraft removed"})
}

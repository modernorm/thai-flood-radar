package routes

import (
	"net/http"

	"github.com/aouyuu/thai-flood-radar/services"
	"github.com/gin-gonic/gin"
)

// GetAlerts handles GET /v1/alerts
func GetAlerts(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetAlerts())
}

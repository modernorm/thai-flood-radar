package routes

import (
	"net/http"

	"github.com/aouyuu/thai-flood-radar/services"
	"github.com/gin-gonic/gin"
)

// GetCurrentRain handles GET /v1/rain/current
func GetCurrentRain(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetRainOverview())
}

// GetRainForecast handles GET /v1/rain/forecast
func GetRainForecast(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetRainForecast())
}

package main

import (
	"github.com/aouyuu/thai-flood-radar/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// HTML templates
	r.LoadHTMLGlob("templates/*")

	// Web dashboard
	r.GET("/", routes.Dashboard)

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong", "status": "ok"})
	})

	// API v1
	v1 := r.Group("/v1")
	{
		// Flood affected areas
		v1.GET("/affected", routes.GetAffectedArea)
		v1.GET("/affected/overview", routes.GetAffectedAreasOverview)

		// Rain
		v1.GET("/rain/current", routes.GetCurrentRain)
		v1.GET("/rain/forecast", routes.GetRainForecast)

		// Alerts
		v1.GET("/alerts", routes.GetAlerts)

		// Provinces
		v1.GET("/provinces", routes.GetProvinces)
	}

	r.Run(":8080")
}

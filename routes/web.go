package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dashboard serves the main web UI.
func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Thai Flood Radar",
	})
}

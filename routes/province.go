package routes

import (
	"net/http"

	"github.com/aouyuu/thai-flood-radar/data"
	"github.com/gin-gonic/gin"
)

// GetProvinces handles GET /v1/provinces
func GetProvinces(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"total":     len(data.Provinces),
		"provinces": data.Provinces,
	})
}

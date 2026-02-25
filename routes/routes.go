package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/aouyuu/thai-flood-radar/services"
	"github.com/gin-gonic/gin"
)

// GetAffectedAreasOverview handles GET /v1/affected/overview?date=<epoch>
func GetAffectedAreasOverview(c *gin.Context) {
	dateStr := c.Query("date")
	var date time.Time
	if dateStr == "" {
		date = time.Now()
	} else {
		epoch, err := strconv.ParseInt(dateStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date parameter; use epoch seconds"})
			return
		}
		date = time.Unix(epoch, 0)
	}

	overview := services.GetAffectedOverview(date)
	c.JSON(http.StatusOK, overview)
}

// GetAffectedArea handles GET /v1/affected?provinceId=<id>&fromDate=<epoch>&toDate=<epoch>
func GetAffectedArea(c *gin.Context) {
	pidStr := c.Query("provinceId")
	if pidStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "provinceId is required"})
		return
	}
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid provinceId"})
		return
	}

	fromStr := c.Query("fromDate")
	toStr := c.Query("toDate")
	from := time.Now().AddDate(0, 0, -7)
	to := time.Now()

	if fromStr != "" {
		if e, err := strconv.ParseInt(fromStr, 10, 64); err == nil {
			from = time.Unix(e, 0)
		}
	}
	if toStr != "" {
		if e, err := strconv.ParseInt(toStr, 10, 64); err == nil {
			to = time.Unix(e, 0)
		}
	}

	area, ok := services.GetAffectedProvince(pid, from, to)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "province not found"})
		return
	}
	c.JSON(http.StatusOK, area)
}

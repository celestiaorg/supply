package main

import (
	"net/http"
	"time"

	"github.com/celestiaorg/supply/supply"
	"github.com/gin-gonic/gin"
)

type response struct {
	// Available is the total supply of utia available at the given time.
	Available int64 `json:"available"`
	// Circulating is the circulating supply of utia available at the given time.
	Circulating int64 `json:"circulating"`
}

// getSupply responds with the current supply of utia.
func getSupply(c *gin.Context) {
	t := time.Now()
	c.IndentedJSON(http.StatusOK, response{
		Available:   supply.AvailableSupply(t),
		Circulating: supply.CirculatingSupply(t),
	})
}

// getSupply responds with the supply of utia at the given date.
func getSupplyByDate(c *gin.Context) {
	date := c.Param("date")
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid date parameter. Date must be in the format YYYY-MM-DD.",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, response{
		Available:   supply.AvailableSupply(t),
		Circulating: supply.CirculatingSupply(t),
	})
}

func main() {
	router := gin.Default()
	router.GET("/supply", getSupply)
	router.GET("/supply/:date", getSupplyByDate)
	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}

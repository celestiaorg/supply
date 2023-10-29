package main

import (
	"net/http"
	"time"

	"github.com/celestiaorg/supply/supply"
	"github.com/gin-gonic/gin"
)

type response struct {
	Available   int64 `json:"available"`
	Circulating int64 `json:"circulating"`
}

// getSupply responds with the list of all albums as JSON.
func getSupply(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, response{
		Available:   supply.AvailableSupply(time.Now()),
		Circulating: supply.CirculatingSupply(time.Now()),
	})
}

func main() {
	router := gin.Default()
	router.GET("/supply", getSupply)
	router.Run("localhost:8080")
}

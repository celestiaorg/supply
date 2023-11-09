package main

import (
	"fmt"
	"time"

	"github.com/celestiaorg/supply/internal"
	"github.com/gin-gonic/gin"
)

const utiaPerTia = 100_000
const apiVersion = "v0"

func getCirculatingSupply(c *gin.Context) {
	t := time.Now()
	cs := internal.CirculatingSupply(t)
	c.String(200, formatTia(cs))
}

func getTotalSupply(c *gin.Context) {
	t := time.Now()
	cs := internal.TotalSupply(t)
	c.String(200, formatTia(cs))
}

func main() {
	router := gin.Default()
	router.GET(fmt.Sprintf("/%s/circulating-supply", apiVersion), getCirculatingSupply)
	router.GET(fmt.Sprintf("/%s/total-supply", apiVersion), getTotalSupply)
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}

func formatTia(utia int64) string {
	tia := float64(utia) / utiaPerTia
	return fmt.Sprintf("%f", tia)
}

package main

import (
	"fmt"
	"time"

	"github.com/celestiaorg/supply/internal"
	"github.com/gin-gonic/gin"
)

const utiaPerTia = 100_000

func getCirculating(c *gin.Context) {
	t := time.Now()
	circulating := internal.CirculatingSupply(t)
	c.String(200, formatTia(circulating))
}

func main() {
	router := gin.Default()
	router.GET("/supply/circulating", getCirculating)
	// router.GET("/supply/total", getTotal)
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}

func formatTia(utia int64) string {
	tia := float64(utia) / utiaPerTia
	return fmt.Sprintf("%f", tia)
}

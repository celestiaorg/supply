package main

import (
	"time"

	"github.com/celestiaorg/supply/internal"
	"github.com/gin-gonic/gin"
)

const landingPage = `
The API routes listed below return values in TIA:

/v0/circulating-supply
/v0/total-supply
`

func getLandingPage(c *gin.Context) {
	c.String(200, landingPage)
}

func getCirculatingSupply(c *gin.Context) {
	t := time.Now()
	cs := internal.CirculatingSupply(t)
	c.String(200, internal.FormatTia(cs))
}

func getTotalSupply(c *gin.Context) {
	t := time.Now()
	cs := internal.TotalSupply(t)
	c.String(200, internal.FormatTia(cs))
}

func main() {
	router := gin.Default()
	router.GET("/", getLandingPage)
	router.GET("/v0/circulating-supply", getCirculatingSupply)
	router.GET("/v0/total-supply", getTotalSupply)
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}

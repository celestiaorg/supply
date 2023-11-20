package main

import (
	"net/http"
	"time"

	"github.com/celestiaorg/supply/internal"
	"github.com/gin-gonic/gin"
)

const (
	// RouteCirculatingSupply is the route for circulating supply.
	RouteCirculatingSupply = "/v0/circulating-supply"

	// RouteTotalSupply is the route for total supply.
	RouteTotalSupply = "/v0/total-supply"
)

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
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"RouteCirculatingSupply": RouteCirculatingSupply,
			"RouteTotalSupply":       RouteTotalSupply,
		})
	})
	router.GET(RouteCirculatingSupply, getCirculatingSupply)
	router.GET(RouteTotalSupply, getTotalSupply)
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}

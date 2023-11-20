package main

import (
	"html/template"
	"io"
	"net/http"
	"strings"
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
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	router.SetHTMLTemplate(t)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/index.tmpl", gin.H{
			"RouteCirculatingSupply": RouteCirculatingSupply,
			"RouteTotalSupply":       RouteTotalSupply,
		})
	})
	router.GET(RouteCirculatingSupply, getCirculatingSupply)
	router.GET(RouteTotalSupply, getTotalSupply)
	err = router.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}

// loadTemplate loads templates embedded by go-assets-builder
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

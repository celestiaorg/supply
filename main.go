package main

import (
	"fmt"
	"time"

	"github.com/celestiaorg/supply/internal"
	"github.com/gin-gonic/gin"
)

const landingPage = `
The API routes listed below return values in TIA:

/v0/circulating-supply
/v0/circulating-supply?date=2025-05-31
/v0/total-supply
/v0/total-supply?date=2025-05-31

Note on methodology: due to complexity, this supply API does not adjust the circulating supply for tokens that were retroactively locked after TGE, for example due to CIP-31. As a result, the reported circulating supply may include a minor margin of error (likely no more than around 1%), decreasing to zero as lockups elapse.
`

func getLandingPage(c *gin.Context) {
	c.String(200, landingPage)
}

func getCirculatingSupply(c *gin.Context) {
	t, err := getDate(c)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	cs := internal.CirculatingSupply(t)
	c.String(200, internal.FormatTia(cs))
}

func getTotalSupply(c *gin.Context) {
	t, err := getDate(c)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	cs := internal.TotalSupply(t)
	c.String(200, internal.FormatTia(cs))
}

// getDate returns the date from the query parameter or the current date if no
// query parameter is provided.
func getDate(c *gin.Context) (time.Time, error) {
	date := c.Query("date")
	if date == "" {
		return time.Now(), nil
	}
	t, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return time.Time{}, fmt.Errorf("please provide date as YYYY-MM-DD. Invalid date format: %w", err)
	}
	return t, nil
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

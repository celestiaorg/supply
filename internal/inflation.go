package internal

import (
	"math"
	"time"
)

const (
	day  = time.Hour * 24
	year = day * 365
)

const (
	// initialInflationRate is the inflation rate that the network starts at.
	initialInflationRate = 0.08
	// disinflationRate is the rate at which the inflation rate decreases each year.
	disinflationRate = 0.1
	// targetInflationRate is the inflation rate that the network aims to
	// stabalize at. In practice, TargetInflationRate acts as a minimum so that
	// the inflation rate doesn't decrease after reaching it.
	targetInflationRate = 0.015
)

// roundToDecimalPlaces rounds a float64 to the specified number of decimal places
func roundToDecimalPlaces(value float64, places int) float64 {
	multiplier := math.Pow(10, float64(places))
	return math.Round(value*multiplier) / multiplier
}

// inflationRate returns the inflation rate at the given time.
func inflationRate(t time.Time) float64 {
	yearsSinceGenesis := yearsSinceGenesis(t)
	inflationRate := initialInflationRate * math.Pow(float64(1-disinflationRate), float64(yearsSinceGenesis))

	// HACK: round to 6 decimal places to avoid floating-point precision issues.
	inflationRate = roundToDecimalPlaces(inflationRate, 6)
	return max(inflationRate, targetInflationRate)
}

// cumulativeInflation returns the total amount of utia that will be minted due
// to inflation from genesis up until t.
func cumulativeInflation(t time.Time) int64 {
	daysSinceGenesis := daysSinceGenesis(t)
	if daysSinceGenesis == 0 {
		return 0
	}

	return cumulativeInflation(t.Add(-day)) + dailyProvisions(t.Add(-day))
}

// totalSupply returns the total supply of utia at the start of the year at the given time.
func totalSupply(t time.Time) int64 {
	yearsSinceGenesis := yearsSinceGenesis(t)
	if yearsSinceGenesis == 0 {
		return initialTotalSupplyInUtia
	}

	return totalSupply(t.Add(-year)) + annualProvisions(t.Add(-year))
}

// annualProvisions returns the amount of utia that will be minted due to inflation for the year at the given time t.
func annualProvisions(t time.Time) int64 {
	return int64(float64(totalSupply(t)) * inflationRate(t))
}

// annualProvisions returns the amount of utia that will be minted due to inflation for the day at the given time t.
func dailyProvisions(t time.Time) int64 {
	return annualProvisions(t) / 365
}

func yearsSinceGenesis(t time.Time) int64 {
	return daysSinceGenesis(t) / 365
}

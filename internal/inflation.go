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
	// Genesis era (before CIP-29 activation in celestia-app app verrsion 4)
	genesisInflationRate    = 0.08
	genesisDisinflationRate = 0.1
	// CIP-29 era (after CIP-29 activation in celestia-app app version 4)
	cip29InflationRate    = 0.0536
	cip29DisinflationRate = 0.067
	// CIP-41 era (after CIP-41 activation in celestia-app app version 6)
	cip41InflationRate    = 0.0267
	cip41DisinflationRate = 0.067
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
	initialRate := getInitialInflationRate(t)
	disinflationRate := getDisinflationRate(t)
	inflationRate := initialRate * math.Pow(float64(1-disinflationRate), float64(yearsSinceGenesis))

	// HACK: round to 6 decimal places to avoid floating-point precision issues.
	inflationRate = roundToDecimalPlaces(inflationRate, 6)
	return max(inflationRate, targetInflationRate)
}

// cumulativeInflation returns the total amount of utia that will be minted due
// to inflation from genesis up until t.
func cumulativeInflation(t time.Time) int64 {
	if t.Before(TGE) || t.Equal(TGE) {
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

func getInitialInflationRate(t time.Time) float64 {
	if t.Before(cip29ActivationDate) {
		return genesisInflationRate
	}
	if t.Before(cip41ActivationDate) {
		return cip29InflationRate
	}

	return cip41InflationRate
}

func getDisinflationRate(t time.Time) float64 {
	if t.Before(cip29ActivationDate) {
		return genesisDisinflationRate
	}
	if t.Before(cip41ActivationDate) {
		return cip29DisinflationRate
	}

	return cip29DisinflationRate
}

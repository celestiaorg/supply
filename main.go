package supply

import (
	"time"
)

const (
	publicAllocationCirculating = 75_000_000 // depends on governance
	ecosystemCirculating        = 0          // depends on governance
	investorsTotal              = 355_689_414
	coreContributorsTotal       = 176_365_875

	// InitialInflationRate is the inflation rate that the network starts at.
	InitialInflationRate = 0.08
	// DisinflationRate is the rate at which the inflation rate decreases each year.
	DisinflationRate = 0.1
	// TargetInflationRate is the inflation rate that the network aims to
	// stabalize at. In practice, TargetInflationRate acts as a minimum so that
	// the inflation rate doesn't decrease after reaching it.
	TargetInflationRate = 0.015
)

var TGE = time.Date(2023, time.October, 31, 0, 0, 0, 0, time.UTC)

// Circulating returns the circulating supply of TIA at the given time.
func Circulating(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}

	daysSinceGenesis := int64(t.Sub(TGE).Hours() / 24)
	if daysSinceGenesis < 365 {
		return publicAllocationCirculating + ecosystemCirculating + cumulativeInflation(daysSinceGenesis)
	}
	return publicAllocationCirculating + ecosystemCirculating + cumulativeInflation(daysSinceGenesis) + investorsCirculating(t) + coreContributorsCirculating(t)
}

// func cumulativeInflation(t time.Time) int64 {
// 	return 0
// }

func investorsCirculating(t time.Time) int64 {
	dateDiff := t.Sub(TGE).Hours() / 24
	return int64(min(investorsTotal/3+investorsTotal*2/3/365*(dateDiff-365), investorsTotal))
}

func coreContributorsCirculating(t time.Time) int64 {
	dateDiff := t.Sub(TGE).Hours() / 24
	return int64(min(coreContributorsTotal/3+coreContributorsTotal*2/3/(365*2)*(dateDiff-365), coreContributorsTotal))
}

// Available returns the available supply of TIA at the given time.
func Available(t time.Time) int64 {
	// TODO
	return 0
}

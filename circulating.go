package supply

import (
	"time"
)

const (
	publicAllocation      = 75_000_000 * utiaPerTia // depends on governance
	ecosystemAllocation   = 0                       // depends on governance
	investorsTotal        = 355_689_414 * utiaPerTia
	coreContributorsTotal = 176_365_875 * utiaPerTia
)

var TGE = time.Date(2023, time.October, 31, 14, 0, 0, 0, time.UTC)
var oneYearAfterTGE = TGE.AddDate(1, 0, 0)

// circulatingSupply returns the circulating supply of utia at the given time.
func circulatingSupply(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}

	daysSinceGenesis := int64(t.Sub(TGE).Hours() / 24)
	if daysSinceGenesis < 365 {
		return publicAllocationCirculating(t) + ecosystemAllocationCirculating(t) + cumulativeInflation(daysSinceGenesis)
	}
	return publicAllocationCirculating(t) + ecosystemAllocationCirculating(t) + cumulativeInflation(daysSinceGenesis) + investorsCirculating(t) + coreContributorsCirculating(t)
}

func publicAllocationCirculating(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}
	return publicAllocation
}

func ecosystemAllocationCirculating(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}
	return ecosystemAllocation
}

func investorsCirculating(t time.Time) int64 {
	if t.Before(oneYearAfterTGE) {
		return 0
	}
	daysSinceGenesis := int64(t.Sub(TGE).Hours() / 24)
	// 1/3 of total unlocks in chunk at TGE + 1 year.
	// Remaining 2/3 of total unlocks from TGE + 1 year to TGE + 2 years.
	return min(investorsTotal/3+investorsTotal*2/3/365*(daysSinceGenesis-365), investorsTotal)
}

func coreContributorsCirculating(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}
	daysSinceGenesis := int64(t.Sub(TGE).Hours() / 24)
	return int64(min(coreContributorsTotal/3+coreContributorsTotal*2/3/(365*2)*(daysSinceGenesis-365), coreContributorsTotal))
}

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
var twoYearsAfterTGE = TGE.AddDate(2, 0, 0)
var threeYearsAfterTGE = TGE.AddDate(3, 0, 0)

// circulatingSupply returns the circulating supply of utia at the given time.
func circulatingSupply(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}

	days := daysSinceGenesis(t)
	return publicAllocationCirculating(t) +
		ecosystemAllocationCirculating(t) +
		cumulativeInflation(days) +
		investorsCirculating(t) +
		coreContributorsCirculating(t)
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
	if t.Equal(twoYearsAfterTGE) || t.After(twoYearsAfterTGE) {
		return investorsTotal
	}
	// 1/3 of investor tokens unlocks in a chunk at TGE + 1 year. The remaining
	// 2/3 of investor tokens unlocks linearly from TGE + 1 year to TGE + 2
	// years.
	days := daysSinceGenesis(t)
	return investorsTotal/3 + investorsTotal*2/3/365*(days-365)
}

func coreContributorsCirculating(t time.Time) int64 {
	if t.Before(oneYearAfterTGE) {
		return 0
	}
	if t.Equal(threeYearsAfterTGE) || t.After(threeYearsAfterTGE) {
		return coreContributorsTotal
	}
	// 1/3 of core contributor tokens unlocks in a chunk at TGE + 1 year. The
	// remaining 2/3 of tokens unlocks linearly from TGE + 1 year to TGE + 3
	// years.
	days := daysSinceGenesis(t)
	return coreContributorsTotal/3 + coreContributorsTotal*2/3/(365*2)*(days-365)
}

func daysSinceGenesis(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}
	return int64(t.Sub(TGE).Hours() / 24)
}

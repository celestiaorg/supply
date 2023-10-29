package supply

import (
	"time"
)

// CirculatingSupply returns the circulating supply of utia at the given time.
func CirculatingSupply(t time.Time) int64 {
	days := daysSinceGenesis(t)
	return cumulativeInflation(days) +
		publicAllocationCirculating(t) +
		ecosystemCirculating(t) +
		investorsCirculating(t) +
		coreContributorsCirculating(t)
}

// publicAllocationCirculating returns the circulating supply of utia at the
// given time. The public allocation genesis is fully unlocked at TGE. The
// public allocation for future iniatives is subject to on chain governance.
func publicAllocationCirculating(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}
	// TODO: account for publicAllocationFuture when on-chain governance votes to spend it.
	return publicAllocationGenesis
}

func ecosystemCirculating(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}
	// TODO(@rootulp): 1 billion total supply of TIA * .268 (total supply for ecosystem) * .25 (% of ecosystem unlocked at launch) = 67m
	// so why does the spreadsheet say 50m?
	return 50_000_000 * utiaPerTia
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

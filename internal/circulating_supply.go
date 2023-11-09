package internal

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
	// Note: account for publicAllocationFuture when on-chain governance votes to spend it.
	return publicAllocationGenesis
}

func ecosystemCirculating(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}
	if t.Equal(oneYearAfterTGE) || t.Before(oneYearAfterTGE) {
		return .25 * ecosystem
	}
	if t.Equal(fourYearsAfterTGE) || t.After(fourYearsAfterTGE) {
		return ecosystem
	}
	days := daysSinceGenesis(t)
	// 25% unlocked at launch. Remaining 75% unlocks continuously from year 1 to
	// year 4.
	return ecosystem/4 + (days-365)*ecosystem*3/4/(365*3)
}

func investorsCirculating(t time.Time) int64 {
	if t.Before(oneYearAfterTGE) {
		return 0
	}
	if t.Equal(twoYearsAfterTGE) || t.After(twoYearsAfterTGE) {
		return investorsTotal
	}
	// 33.3% of investor tokens unlocks in a chunk at TGE + 1 year. The remaining
	// 66.6% of investor tokens unlocks linearly from TGE + 1 year to TGE + 2
	// years.
	days := daysSinceGenesis(t)
	return investorsTotal/3 + investorsTotal*2/3/365*(days-365)
}

func coreContributorsCirculating(t time.Time) int64 {
	if t.Before(oneYearAfterTGE) {
		return 0
	}
	if t.Equal(threeYearsAfterTGE) || t.After(threeYearsAfterTGE) {
		return coreContributors
	}
	// 33.3% of core contributor tokens unlocks in a chunk at TGE + 1 year. The
	// remaining 66.6% of tokens unlocks linearly from TGE + 1 year to TGE + 3
	// years.
	days := daysSinceGenesis(t)
	return coreContributors/3 + coreContributors*2/3/(365*2)*(days-365)
}

func daysSinceGenesis(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}
	return int64(t.Sub(TGE).Hours() / 24)
}

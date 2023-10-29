package supply

import (
	"time"
)

// The following constants are sourced from
// https://docs.celestia.org/learn/staking-governance-supply#tia-allocation-at-genesis
const (
	// publicAllocationGenesis is the number of tokens allocated to the public
	// via genesis drop & incentivized testnet.
	publicAllocationGenesis = 75_000_000 * utiaPerTia

	// publicAllocationFuture is the number of tokens allocated to the public
	// to be deployed in future initiatives.
	publicAllocationFuture = 125_000_000 * utiaPerTia

	// publicAllocationTotal is the total number of tokens allocated to the
	// public.
	publicAllocationTotal = publicAllocationGenesis + publicAllocationFuture

	// ecosystem is the number of tokens allocated to the ecosystem at genesis.
	ecosystem = 267_944_711 * utiaPerTia

	// investorsTotal is the total number of tokens allocated to investors at
	// genesis. This includes investors in the seed, series A, and series B
	// rounds.
	investorsTotal = 355_689_414 * utiaPerTia

	// coreContributorsTotal is the total number of tokens allocated to core
	// contributors at genesis.
	coreContributorsTotal = 176_365_875 * utiaPerTia
)

var TGE = time.Date(2023, time.October, 31, 14, 0, 0, 0, time.UTC)
var oneYearAfterTGE = TGE.AddDate(1, 0, 0)
var twoYearsAfterTGE = TGE.AddDate(2, 0, 0)
var threeYearsAfterTGE = TGE.AddDate(3, 0, 0)
var fourYearsAfterTGE = TGE.AddDate(4, 0, 0)

// circulatingSupply returns the circulating supply of utia at the given time.
func circulatingSupply(t time.Time) int64 {
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

// ecosystemCirculating returns the circulating supply of utia in the ecosystem category at the given time.
// 25% unlocked at launch. Remaining 75% unlocks continuously from year 1 to year 4.
func ecosystemCirculating(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}
	// TODO(@rootulp): 1 billion total supply of TIA * .268 (total supply for ecosystem) * .25 (% of ecosystem unlocked at launch) = 67m
	// so why does the spreadsheet say 50m?
	return 50_000_000 * utiaPerTia
}

// investorsCirculating returns the circulating supply of utia in the investors
// category at the given time. 1/3 of investor tokens unlocks in a chunk at TGE
// + 1 year. The remaining 2/3 of investor tokens unlocks linearly from TGE + 1
// year to TGE + 2 years.
func investorsCirculating(t time.Time) int64 {
	if t.Before(oneYearAfterTGE) {
		return 0
	}
	if t.Equal(twoYearsAfterTGE) || t.After(twoYearsAfterTGE) {
		return investorsTotal
	}
	days := daysSinceGenesis(t)
	return investorsTotal/3 + investorsTotal*2/3/365*(days-365)
}

// coreContributorsCirculating returns the circulating supply of utia in the
// core contributors category at the given time. 1/3 of core contributor tokens
// unlocks in a chunk at TGE + 1 year. The remaining 2/3 of tokens unlocks
// linearly from TGE + 1 year to TGE + 3 years.
func coreContributorsCirculating(t time.Time) int64 {
	if t.Before(oneYearAfterTGE) {
		return 0
	}
	if t.Equal(threeYearsAfterTGE) || t.After(threeYearsAfterTGE) {
		return coreContributorsTotal
	}
	days := daysSinceGenesis(t)
	return coreContributorsTotal/3 + coreContributorsTotal*2/3/(365*2)*(days-365)
}

func daysSinceGenesis(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}
	return int64(t.Sub(TGE).Hours() / 24)
}

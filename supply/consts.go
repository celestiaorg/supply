package supply

import "time"

const (
	utiaPerTia               = 1_000_000                       // 1 tia = 1 million utia
	initialTotalSupply       = 1_000_000_000                   // 1 billion tia
	initialTotalSupplyInUtia = initialTotalSupply * utiaPerTia // 1 quadrillion utia
)

// The following constants are sourced from
// https://docs.celestia.org/learn/staking-governance-supply#tia-allocation-at-genesis
// https://docs.google.com/spreadsheets/d/1OoHEIF0oLSFcqywQIn8Vf6T6VHkza5E1bIag-zVWhiA
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

	// ecosystem is the number of tokens allocated to R&D and the ecosystem at
	// genesis.
	ecosystem = 267_944_711 * utiaPerTia

	// investorsTotal is the total number of tokens allocated to investors at
	// genesis. This includes investors in the seed, series A, and series B
	// rounds.
	investorsTotal = 355_689_414 * utiaPerTia

	// coreContributorsTotal is the total number of tokens allocated to core
	// contributors at genesis.
	coreContributorsTotal = 176_365_875 * utiaPerTia
)

var (
	beforeTGE          = time.Date(2023, time.October, 31, 13, 0, 0, 0, time.UTC)
	TGE                = time.Date(2023, time.October, 31, 14, 0, 0, 0, time.UTC)
	oneDayAfterTGE     = TGE.AddDate(0, 0, 1)
	oneYearAfterTGE    = TGE.AddDate(1, 0, 0)
	twoYearsAfterTGE   = TGE.AddDate(2, 0, 0)
	threeYearsAfterTGE = TGE.AddDate(3, 0, 0)
	fourYearsAfterTGE  = TGE.AddDate(4, 0, 0)
)

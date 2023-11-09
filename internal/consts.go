package internal

const (
	// Sourced from
	// https://github.com/celestiaorg/networks/blob/master/celestia/genesis.json
	initialTotalSupplyInTia  = 1_000_000_000                        // 1 billion tia
	initialTotalSupplyInUtia = initialTotalSupplyInTia * utiaPerTia // 1 quadrillion utia
)

// The following constants are sourced from
// https://docs.celestia.org/learn/staking-governance-supply#tia-allocation-at-genesis
// https://docs.google.com/spreadsheets/d/1OoHEIF0oLSFcqywQIn8Vf6T6VHkza5E1bIag-zVWhiA
const (
	// publicAllocationGenesis is the number of tokens allocated to the public
	// via genesis drop & incentivized testnet.
	publicAllocationGenesis = 74_057_350 * utiaPerTia

	// publicAllocationFuture is the number of tokens allocated to the public
	// to be deployed in future initiatives.
	publicAllocationFuture = 125_942_650 * utiaPerTia

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

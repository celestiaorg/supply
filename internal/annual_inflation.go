package internal

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

func inflationRate(yearsSinceGenesis int64) float64 {
	if yearsSinceGenesis == 0 {
		return initialInflationRate
	}
	return max(inflationRate(yearsSinceGenesis-1)*(1-disinflationRate), targetInflationRate)
}

func totalSupply(yearsSinceGenesis int64) int64 {
	if yearsSinceGenesis == 0 {
		return initialTotalSupplyInUtia
	}

	return totalSupply(yearsSinceGenesis-1) + annualInflation(yearsSinceGenesis-1)
}

func annualInflation(yearsSinceGenesis int64) int64 {
	return int64(inflationRate(yearsSinceGenesis) * float64(totalSupply(yearsSinceGenesis)))
}

func dailyInflationForYear(yearsSinceGenesis int64) int64 {
	return annualInflation(yearsSinceGenesis) / 365
}

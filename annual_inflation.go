package supply

const (
	utiaPerTia               = 1_000_000                       // 1 tia = 1 million utia
	initialTotalSupply       = 1_000_000_000                   // 1 billion tia
	initialTotalSupplyInUtia = initialTotalSupply * utiaPerTia // 1 quadrillion utia
	initialInflationRate     = 0.08
	inflationRateDecay       = 0.9
	minInflationRate         = 0.015
)

func inflationRate(yearsSinceGenesis int64) float64 {
	if yearsSinceGenesis == 0 {
		return initialInflationRate
	}
	return max(inflationRate(yearsSinceGenesis-1)*inflationRateDecay, minInflationRate)
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

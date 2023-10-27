package supply

const (
	initialTotalSupply   = 1_000_000_000
	initialInflationRate = 0.08
	inflationRateDecay   = 0.9
	minInflationRate     = 0.015
)

func inflationRate(yearsSinceGenesis int64) float64 {
	if yearsSinceGenesis == 0 {
		return initialInflationRate
	}
	return max(inflationRate(yearsSinceGenesis-1)*inflationRateDecay, minInflationRate)
}

func totalSupply(yearsSinceGenesis int64) int64 {
	if yearsSinceGenesis == 0 {
		return initialTotalSupply
	}

	return totalSupply(yearsSinceGenesis-1) + annualInflation(yearsSinceGenesis-1)
}

func annualInflation(yearsSinceGenesis int64) int64 {
	return int64(inflationRate(yearsSinceGenesis) * float64(totalSupply(yearsSinceGenesis)))
}

func dailyInflationForYear(yearsSinceGenesis int64) int64 {
	return annualInflation(yearsSinceGenesis) / 365
}

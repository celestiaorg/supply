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

func totalSupply(year int64) int64 {
	if year == 0 {
		return initialTotalSupply
	}

	return totalSupply(year-1) + annualInflation(year-1)
}

func annualInflation(year int64) int64 {
	return int64(inflationRate(year) * float64(totalSupply(year)))
}

func dailyInflation(year int64) int64 {
	return annualInflation(year) / 365
}

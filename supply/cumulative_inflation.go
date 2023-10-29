package supply

// cumulativeInflation returns the total amount of utia that will be minted due to inflation.
func cumulativeInflation(daysSinceGenesis int64) int64 {
	if daysSinceGenesis == 0 {
		return 0
	}

	return cumulativeInflation(daysSinceGenesis-1) + dailyInflationForDay(daysSinceGenesis-1)
}

// dailyInflationForDay returns the amount of utia that will be minted per day due to inflation.
func dailyInflationForDay(daysSinceGenesis int64) int64 {
	if daysSinceGenesis == 0 {
		return dailyInflationForYear(0)
	}
	yearsSinceGenesis := daysSinceGenesis / 365
	return dailyInflationForYear(yearsSinceGenesis)
}

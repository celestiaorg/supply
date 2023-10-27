package supply

func cumulativeInflation(daysSinceGenesis int64) int64 {
	if daysSinceGenesis == 0 {
		return 0
	}

	return cumulativeInflation(daysSinceGenesis-1) + dailyInflationForDay(daysSinceGenesis-1)
}

func dailyInflationForDay(daysSinceGenesis int64) int64 {
	if daysSinceGenesis == 0 {
		return dailyInflationForYear(0)
	}
	yearsSinceGenesis := daysSinceGenesis / 365
	return dailyInflationForYear(yearsSinceGenesis)
}

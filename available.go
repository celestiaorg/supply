package supply

import "time"

// availableSupply returns the available supply of utia at the given time.
func availableSupply(t time.Time) int64 {
	days := daysSinceGenesis(t)
	return cumulativeInflation(days) +
		publicAllocationAvailable(t) +
		ecosystemAvailable(t) +
		investorsAvailable(t) +
		coreContributorsAvailable(t)
}

func publicAllocationAvailable(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}
	return publicAllocationTotal
}

// 25% unlocked at launch. Remaining 75% unlocks continuously from year 1 to
// year 4.
func ecosystemAvailable(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}
	if t.Equal(fourYearsAfterTGE) || t.After(fourYearsAfterTGE) {
		return ecosystem
	}
	days := daysSinceGenesis(t)
	unlockedAtLaunch := ecosystem / 4
	return int64(unlockedAtLaunch) + days*ecosystem*3/4/(365*3)

}

func investorsAvailable(t time.Time) int64 {
	return investorsCirculating(t)
}

func coreContributorsAvailable(t time.Time) int64 {
	return coreContributorsCirculating(t)
}

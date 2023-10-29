package supply

import "time"

// AvailableSupply returns the available supply of utia at the given time.
func AvailableSupply(t time.Time) int64 {
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

func ecosystemAvailable(t time.Time) int64 {
	return ecosystemCirculating(t)
}

func investorsAvailable(t time.Time) int64 {
	return investorsCirculating(t)
}

func coreContributorsAvailable(t time.Time) int64 {
	return coreContributorsCirculating(t)
}

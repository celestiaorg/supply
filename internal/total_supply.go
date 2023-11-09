package internal

import "time"

// TotalSupply returns the total supply of utia at the given time.
func TotalSupply(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}
	if t.Equal(TGE) {
		return initialTotalSupplyInUtia
	}
	days := daysSinceGenesis(t)
	inflation := cumulativeInflation(days)
	return initialTotalSupplyInUtia + inflation
}

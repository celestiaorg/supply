package supply

import (
	"time"
)

const (
	publicAllocationCirculating = 75_000_000 * utiaPerTia // depends on governance
	ecosystemCirculating        = 0                       // depends on governance
	investorsTotal              = 355_689_414 * utiaPerTia
	coreContributorsTotal       = 176_365_875 * utiaPerTia
)

var TGE = time.Date(2023, time.October, 31, 14, 0, 0, 0, time.UTC)

// circulatingSupply returns the circulating supply of utia at the given time.
func circulatingSupply(t time.Time) int64 {
	if t.Before(TGE) {
		return 0
	}

	daysSinceGenesis := int64(t.Sub(TGE).Hours() / 24)
	if daysSinceGenesis < 365 {
		return publicAllocationCirculating + ecosystemCirculating + cumulativeInflation(daysSinceGenesis)
	}
	return publicAllocationCirculating + ecosystemCirculating + cumulativeInflation(daysSinceGenesis) + investorsCirculating(t) + coreContributorsCirculating(t)
}

func investorsCirculating(t time.Time) int64 {
	daysSinceGenesis := int64(t.Sub(TGE).Hours() / 24)
	return int64(min(investorsTotal/3+investorsTotal*2/3/365*(daysSinceGenesis-365), investorsTotal))
}

func coreContributorsCirculating(t time.Time) int64 {
	daysSinceGenesis := int64(t.Sub(TGE).Hours() / 24)
	return int64(min(coreContributorsTotal/3+coreContributorsTotal*2/3/(365*2)*(daysSinceGenesis-365), coreContributorsTotal))
}

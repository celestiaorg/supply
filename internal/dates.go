package internal

import "time"

var (
	oneHourBeforeTGE           = time.Date(2023, time.October, 30, 23, 0, 0, 0, time.UTC)
	TGE                        = time.Date(2023, time.October, 31, 0, 0, 0, 0, time.UTC)
	oneDayAfterTGE             = TGE.AddDate(0, 0, 1)
	oneYearAfterTGEMinusOneDay = TGE.AddDate(0, 0, 364) // October 29, 2024
	oneYearAfterTGE            = TGE.AddDate(0, 0, 365) // October 30, 2024
	oneYearAfterTGEPlusOneDay  = TGE.AddDate(0, 0, 366) // October 31, 2024

	// CIP-29 activated on Mainnet at the v4 activation height height (6680339).
	// See https://www.mintscan.io/celestia/block/6680339
	cip29ActivationDate = time.Date(2025, time.July, 28, 0, 0, 0, 0, time.UTC)

	// CIP-41 activated on Mainnet at the v6 activation height (8662012).
	// See https://www.mintscan.io/celestia/block/8662012
	cip41ActivationDate = time.Date(2025, time.November, 24, 0, 0, 0, 0, time.UTC)

	// TODO: verify these dates. The unlock dates may not be exactly N years
	// after TGE. Instead, they may be N * 365 days after.
	twoYearsAfterTGE   = TGE.AddDate(2, 0, 0)
	threeYearsAfterTGE = TGE.AddDate(3, 0, 0)
	fourYearsAfterTGE  = TGE.AddDate(4, 0, 0)
)

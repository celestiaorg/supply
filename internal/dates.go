package internal

import "time"

var (
	oneHourBeforeTGE           = time.Date(2023, time.October, 30, 23, 0, 0, 0, time.UTC)
	TGE                        = time.Date(2023, time.October, 31, 0, 0, 0, 0, time.UTC)
	oneDayAfterTGE             = TGE.AddDate(0, 0, 1)
	oneYearAfterTGEMinusOneDay = TGE.AddDate(0, 0, 364) // October 29, 2024
	oneYearAfterTGE            = TGE.AddDate(0, 0, 365) // October 30, 2024
	oneYearAfterTGEPlusOneDay  = TGE.AddDate(0, 0, 366) // October 31, 2024

	// TODO: CIP-29 hasn't activated on Mainnet yet. It will activate when v4
	// activates on Mainnet so this date needs to be updated.
	// cip29ActivationDate = time.Date(2025, time.July, 31, 0, 0, 0, 0, time.UTC)

	// TODO: verify these dates. The unlock dates may not be exactly N years
	// after TGE. Instead, they may be N * 365 days after.
	twoYearsAfterTGE   = TGE.AddDate(2, 0, 0)
	threeYearsAfterTGE = TGE.AddDate(3, 0, 0)
	fourYearsAfterTGE  = TGE.AddDate(4, 0, 0)
)

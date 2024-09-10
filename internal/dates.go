package internal

import "time"

var (
	oneHourBeforeTGE           = time.Date(2023, time.October, 31, 13, 0, 0, 0, time.UTC)
	TGE                        = time.Date(2023, time.October, 31, 14, 0, 0, 0, time.UTC)
	oneDayAfterTGE             = TGE.AddDate(0, 0, 1)
	oneYearAfterTGEMinusOneDay = TGE.AddDate(0, 0, 364) // October 29, 2024
	oneYearAfterTGE            = TGE.AddDate(0, 0, 365) // October 30, 2024

	// TODO: verify these dates. The unlock dates may not be exactly N years
	// after TGE. Instead, they may be N * 365 days after.
	twoYearsAfterTGE   = TGE.AddDate(2, 0, 0)
	threeYearsAfterTGE = TGE.AddDate(3, 0, 0)
	fourYearsAfterTGE  = TGE.AddDate(4, 0, 0)
)

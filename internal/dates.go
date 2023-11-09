package internal

import "time"

var (
	beforeTGE          = time.Date(2023, time.October, 31, 13, 0, 0, 0, time.UTC)
	TGE                = time.Date(2023, time.October, 31, 14, 0, 0, 0, time.UTC)
	oneDayAfterTGE     = TGE.AddDate(0, 0, 1)
	oneYearAfterTGE    = TGE.AddDate(1, 0, 0)
	twoYearsAfterTGE   = TGE.AddDate(2, 0, 0)
	threeYearsAfterTGE = TGE.AddDate(3, 0, 0)
	fourYearsAfterTGE  = TGE.AddDate(4, 0, 0)
)

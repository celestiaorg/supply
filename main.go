package supply

import "time"

var TGE = time.Date(2023, time.October, 31, 0, 0, 0, 0, time.UTC)

// Circulating returns the circulating supply of TIA at the given time.
func Circulating(time time.Time) int64 {
	if time.Before(TGE) {
		return 0
	}

	return 0
}

// Available returns the available supply of TIA at the given time.
func Available(time time.Time) int64 {
	return 0
}

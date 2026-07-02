package internal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestUnlockDates verifies that unlock dates are computed as N * 365 days
// after TGE, ignoring leap years. Vesting accounts in the mainnet genesis file
// have end times of October 30th (not October 31st) because they do not
// account for the leap day on February 29th, 2024.
// See https://github.com/celestiaorg/supply/issues/35
func TestUnlockDates(t *testing.T) {
	type testCase struct {
		name string
		got  time.Time
		want time.Time
	}
	testCases := []testCase{
		{"oneYearAfterTGE", oneYearAfterTGE, time.Date(2024, time.October, 30, 0, 0, 0, 0, time.UTC)},
		{"twoYearsAfterTGE", twoYearsAfterTGE, time.Date(2025, time.October, 30, 0, 0, 0, 0, time.UTC)},
		{"threeYearsAfterTGE", threeYearsAfterTGE, time.Date(2026, time.October, 30, 0, 0, 0, 0, time.UTC)},
		{"fourYearsAfterTGE", fourYearsAfterTGE, time.Date(2027, time.October, 30, 0, 0, 0, 0, time.UTC)},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, tc.got)
		})
	}
}

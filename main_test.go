package supply

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCirculating(t *testing.T) {
	type testCase struct {
		time time.Time
		want int64
	}
	testCases := []testCase{
		{time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), 0},
		{time.Date(2023, time.October, 31, 0, 0, 0, 0, time.UTC), 75_000_000}, // TODO check with Nick why this isn't 0
		{time.Date(2023, time.November, 1, 0, 0, 0, 0, time.UTC), 75_219_178},
		{time.Date(2024, time.October, 31, 0, 0, 0, 0, time.UTC), 333_375_530},
	}
	for _, tc := range testCases {
		t.Run(tc.time.String(), func(t *testing.T) {
			got := Circulating(tc.time)
			require.Equal(t, tc.want, got)
		})
	}
}

// func FuzzAvailableMustBeGreaterThanCirculating(f *testing.F) {
// 	f.Fuzz(func(t *testing.T, unixTime int64) {
// 		time := time.Unix(unixTime, 0)
// 		available := Available(time)
// 		circulating := Circulating(time)
// 		require.GreaterOrEqual(t, available, circulating)
// 	})
// }
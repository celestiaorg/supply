package supply

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_circulatingSupply(t *testing.T) {
	type testCase struct {
		time time.Time
		want int64
	}
	testCases := []testCase{
		{time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), 0},             // before TGE day
		{time.Date(2023, time.October, 31, 0, 0, 0, 0, time.UTC), 0}, // before TGE hour
		{time.Date(2023, time.October, 31, 14, 0, 0, 0, time.UTC), 75_000_000_000_000},
		{time.Date(2023, time.November, 1, 14, 0, 0, 0, time.UTC), 75_219_178_082_191},
		{time.Date(2024, time.October, 31, 14, 0, 0, 0, time.UTC), 333_375_529_851_768},
		{time.Date(2025, time.October, 31, 14, 0, 0, 0, time.UTC), 706_393_270_774_762},
	}
	for _, tc := range testCases {
		t.Run(tc.time.String(), func(t *testing.T) {
			got := circulatingSupply(tc.time)
			require.Equal(t, tc.want, got)
		})
	}
}

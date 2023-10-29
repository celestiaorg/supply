package supply

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_availableSupply(t *testing.T) {
	type testCase struct {
		time time.Time
		want int64
	}
	testCases := []testCase{
		{time.Date(2023, time.October, 31, 0, 0, 0, 0, time.UTC), 0},                    // before TGE
		{time.Date(2023, time.October, 31, 14, 0, 0, 0, time.UTC), 266_986_177_750_000}, // TGE
		{time.Date(2023, time.November, 1, 14, 0, 0, 0, time.UTC), 267_388_879_606_848},
		{time.Date(2024, time.October, 31, 14, 0, 0, 0, time.UTC), 592_531_409_126_425},
		{time.Date(2025, time.October, 31, 14, 0, 0, 0, time.UTC), 1_032_535_327_799_419},
	}
	for _, tc := range testCases {
		t.Run(tc.time.String(), func(t *testing.T) {
			got := availableSupply(tc.time)
			require.Equal(t, tc.want, got)
		})
	}
}

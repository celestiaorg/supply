package supply

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAvailableSupply(t *testing.T) {
	type testCase struct {
		time time.Time
		want int64
	}
	testCases := []testCase{
		{beforeTGE, 0},
		{TGE, 266_986_177_750_000},
		{oneDayAfterTGE, 267_205_355_832_191},
		{oneYearAfterTGE, 525_361_707_601_768},
		{twoYearsAfterTGE, 965_549_150_049_419},
		{threeYearsAfterTGE, 1_166_177_169_080_896},
		{fourYearsAfterTGE, 1_304_866_360_072_480},
	}
	for _, tc := range testCases {
		t.Run(tc.time.String(), func(t *testing.T) {
			got := AvailableSupply(tc.time)
			require.Equal(t, tc.want, got)
		})
	}
}

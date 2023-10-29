package supply

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCirculatingSupply(t *testing.T) {
	type testCase struct {
		time time.Time
		want int64
	}
	testCases := []testCase{
		{beforeTGE, 0},
		{TGE, 125_000_000_000_000},
		{oneYearAfterTGE, 383_375_529_851_768},
		{twoYearsAfterTGE, 756_393_270_774_762},
		{threeYearsAfterTGE, 890_035_112_056_239},
		{fourYearsAfterTGE, 961_921_649_072_480},
	}
	for _, tc := range testCases {
		t.Run(tc.time.String(), func(t *testing.T) {
			got := CirculatingSupply(tc.time)
			require.Equal(t, tc.want, got)
		})
	}
}

func Test_publicAllocationCirculating(t *testing.T) {
	type testCase struct {
		time time.Time
		want int64
	}
	testCases := []testCase{
		{beforeTGE, 0},
		{TGE, 75_000_000_000_000},
		{oneYearAfterTGE, 75_000_000_000_000},
	}
	for _, tc := range testCases {
		t.Run(tc.time.String(), func(t *testing.T) {
			got := publicAllocationCirculating(tc.time)
			require.Equal(t, tc.want, got)
		})
	}
}

func Test_investorsCirculating(t *testing.T) {
	type testCase struct {
		time time.Time
		want int64
	}
	testCases := []testCase{
		{beforeTGE, 0},
		{TGE, 0},
		{oneDayAfterTGE, 0},
		{oneYearAfterTGE, 119_212_799_030_136},  // one year after TGE 1/3 of total unlocks
		{twoYearsAfterTGE, 355_689_414_000_000}, // two years after TGE total finishes unlock
	}
	for _, tc := range testCases {
		t.Run(tc.time.String(), func(t *testing.T) {
			got := investorsCirculating(tc.time)
			require.Equal(t, tc.want, got)
		})
	}
}

func Test_coreContributorsCirculating(t *testing.T) {
	type testCase struct {
		time time.Time
		want int64
	}
	testCases := []testCase{
		{beforeTGE, 0},
		{TGE, 0},
		{oneDayAfterTGE, 0},
		{oneYearAfterTGE, 58_949_689_726_027},
		{twoYearsAfterTGE, 117_738_314_725_882},
		{threeYearsAfterTGE, 176_365_875_000_000}, // three years after TGE total finishes unlock
	}
	for _, tc := range testCases {
		t.Run(tc.time.String(), func(t *testing.T) {
			got := coreContributorsCirculating(tc.time)
			require.Equal(t, tc.want, got)
		})
	}
}

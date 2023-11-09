package internal

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
		{TGE, 141_043_527_750_000},
		{oneYearAfterTGE, 399_419_057_601_768},
		{twoYearsAfterTGE, 839_606_500_049_419},
		{threeYearsAfterTGE, 1_040_234_519_080_896},
		{fourYearsAfterTGE, 1_178_923_710_072_480},
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
		{TGE, publicAllocationGenesis},
		{oneYearAfterTGE, publicAllocationGenesis},
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
		{oneYearAfterTGE, 119_212_799_030_136}, // one year after TGE 1/3 of total unlocks
		{twoYearsAfterTGE, investorsTotal},     // two years after TGE total finishes unlock
		{threeYearsAfterTGE, investorsTotal},
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
		{threeYearsAfterTGE, coreContributorsTotal}, // three years after TGE total finishes unlock
	}
	for _, tc := range testCases {
		t.Run(tc.time.String(), func(t *testing.T) {
			got := coreContributorsCirculating(tc.time)
			require.Equal(t, tc.want, got)
		})
	}
}

func Test_ecosystemCirculating(t *testing.T) {
	type testCase struct {
		time time.Time
		want int64
	}
	testCases := []testCase{
		{beforeTGE, 0},
		{TGE, .25 * ecosystem},
		{oneDayAfterTGE, .25 * ecosystem},
		{oneYearAfterTGE, .25 * ecosystem},
		{twoYearsAfterTGE, 134_155_879_274_657},
		{threeYearsAfterTGE, 201_142_057_024_657},
		{fourYearsAfterTGE, ecosystem},
	}
	for _, tc := range testCases {
		t.Run(tc.time.String(), func(t *testing.T) {
			got := ecosystemCirculating(tc.time)
			require.Equal(t, tc.want, got)
		})
	}
}

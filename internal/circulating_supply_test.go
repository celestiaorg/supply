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
		{oneHourBeforeTGE, 0},
		{TGE, 141_043_527_750_000},
		{oneYearAfterTGE, 398_395_290_749_715},
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
		{oneHourBeforeTGE, 0},
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
		{oneHourBeforeTGE, 0},
		{TGE, 0},
		{oneDayAfterTGE, 0},
		{oneYearAfterTGEMinusOneDay, 0},       // one day before tokens unlock.
		{oneYearAfterTGE, investorsTotal / 3}, // tokens unlock on October 30th, 2024.
		{twoYearsAfterTGE, investorsTotal},
		{threeYearsAfterTGE, investorsTotal},
		{fourYearsAfterTGE, investorsTotal},
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
		{oneHourBeforeTGE, 0},
		{TGE, 0},
		{oneDayAfterTGE, 0},
		{oneYearAfterTGEMinusOneDay, 0},         // one day before tokens unlock.
		{oneYearAfterTGE, coreContributors / 3}, // tokens unlock on October 30th, 2024.
		{twoYearsAfterTGE, 117_738_314_725_882},
		{threeYearsAfterTGE, coreContributors},
		{fourYearsAfterTGE, coreContributors},
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
		{oneHourBeforeTGE, 0},
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

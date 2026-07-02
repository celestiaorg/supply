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
		{oneYearAfterTGEMinusOneDay, 220_824_349_667_524},
		{oneYearAfterTGE, 398_395_290_749_715},
		{oneYearAfterTGEPlusOneDay, 399_602_581_376_425},
		{twoYearsAfterTGE, 832939859033664},
		{threeYearsAfterTGE, 987338016490616},
		{fourYearsAfterTGE, 1079918280757792},
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
		{twoYearsAfterTGE, 117_577_249_999_855},
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
		{twoYearsAfterTGE, 133_972_355_500_000},
		{threeYearsAfterTGE, 200_958_533_250_000},
		{fourYearsAfterTGE, ecosystem},
	}
	for _, tc := range testCases {
		t.Run(tc.time.String(), func(t *testing.T) {
			got := ecosystemCirculating(tc.time)
			require.Equal(t, tc.want, got)
		})
	}
}

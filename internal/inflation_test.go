package internal

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_inflationRate(t *testing.T) {
	type testCase struct {
		t    time.Time
		want float64
	}
	testCases := []testCase{
		{TGE, 0.08},
		{TGE.Add(1 * year), 0.072},
		{cip29ActivationDate.Add(-1 * day), 0.072},   // Before CIP-29 activation
		{cip29ActivationDate, 0.050009},              // CIP-29 activation
		{cip29ActivationDate.Add(1 * day), 0.050009}, // After CIP-29 activation
		{TGE.Add(2 * year), 0.046658},
		{TGE.Add(3 * year), 0.043532},
		{TGE.Add(18 * year), 0.015383},
		{TGE.Add(30 * year), 0.015},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("time %v", tc.t), func(t *testing.T) {
			got := inflationRate(tc.t)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_cumulativeInflation(t *testing.T) {
	type testCase struct {
		t    time.Time
		want int64
	}
	testCases := []testCase{
		{TGE, 0},
		{TGE.Add(day), 219_178_082_191},
		{TGE.Add(2 * day), 438_356_164_382},
		{TGE.Add(3 * day), 657_534_246_573},
		{TGE.Add(364 * day), 79_780_821_917_524},
		{TGE.Add(365 * day), 79_999_999_999_715},
		{TGE.Add(366 * day), 80_213_041_095_605},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("time %v", tc.t), func(t *testing.T) {
			got := cumulativeInflation(tc.t)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_dailyProvisions(t *testing.T) {
	type testCase struct {
		t    time.Time
		want int64
	}
	testCases := []testCase{
		{TGE, 219_178_082_191},
		{TGE.Add(day), 219_178_082_191},
		{TGE.Add(364 * day), 219_178_082_191},
		{TGE.Add(365 * day), 213_041_095_890},
		{TGE.Add(366 * day), 213_041_095_890},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("time %v", tc.t), func(t *testing.T) {
			got := dailyProvisions(tc.t)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_annualProvisions(t *testing.T) {
	type testCase struct {
		t    time.Time
		want int64
	}
	testCases := []testCase{
		{TGE, 80000000000000},
		{TGE.Add(1 * year), 77760000000000},
		{TGE.Add(2 * year), 54018766080000},
		{TGE.Add(3 * year), 52751153244994},
		{TGE.Add(20 * year), 28336928866935},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("time %v", tc.t), func(t *testing.T) {
			got := int64(annualProvisions(tc.t))
			assert.Equal(t, tc.want, got)
		})
	}
}

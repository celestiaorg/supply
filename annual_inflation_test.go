package supply

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dailyInflation(t *testing.T) {
	type testCase struct {
		yearsSinceGenesis int64
		want              int64
	}

	testCases := []testCase{
		{0, 219_178_082_191},
		{1, 213_041_095_890},
		{2, 205_542_049_315},
		{3, 196_975_056_699},
		{20, 82_397_466_592},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("yearsSinceGenesis %v", tc.yearsSinceGenesis), func(t *testing.T) {
			got := int64(dailyInflationForYear(tc.yearsSinceGenesis))
			assert.Equal(t, got, tc.want)
		})
	}
}

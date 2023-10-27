package supply

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cumulativeInflation(t *testing.T) {
	type testCase struct {
		daysSinceGenesis int64
		want             float64
	}

	testCases := []testCase{
		{0, 0},
		{1, 219_178},
		{2, 438_356},
		{3, 657_534},
		{15, 3_287_671},
		{25, 5_479_452},
		{50, 10_958_904},
		{100, 21_917_808},
		{364, 79_780_822},
		{365, 80_000_000},
		{366, 80_213_041},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("daysSinceGenesis %v", tc.daysSinceGenesis), func(t *testing.T) {
			got := cumulativeInflation(tc.daysSinceGenesis)
			assert.Equal(t, tc.want, got)
		})
	}

}

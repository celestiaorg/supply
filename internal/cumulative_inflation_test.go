package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cumulativeInflation(t *testing.T) {
	type testCase struct {
		daysSinceGenesis int64
		want             int64
	}
	testCases := []testCase{
		{0, 0},
		{1, 219_178_082_191},
		{2, 438_356_164_382},
		{3, 657_534_246_573},
		{364, 79_780_821_917_524},
		{365, 79_999_999_999_715},
		{366, 80_213_041_095_605},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("daysSinceGenesis %v", tc.daysSinceGenesis), func(t *testing.T) {
			got := cumulativeInflation(tc.daysSinceGenesis)
			assert.Equal(t, tc.want, got)
		})
	}
}

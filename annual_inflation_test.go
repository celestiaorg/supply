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
		{0, 219_178},
		{1, 213_041},
		{2, 205_542},
		{3, 196_975},
		{20, 82_397},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("yearsSinceGenesis %v", tc.yearsSinceGenesis), func(t *testing.T) {
			got := dailyInflation(tc.yearsSinceGenesis)
			assert.Equal(t, got, tc.want)
		})
	}
}

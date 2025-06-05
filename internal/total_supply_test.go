package internal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTotalSupply(t *testing.T) {
	type testCase struct {
		time time.Time
		want int64
	}
	testCases := []testCase{
		{oneHourBeforeTGE, 0},
		{TGE, initialTotalSupplyInUtia},
		{oneDayAfterTGE, 1000219178082191},
		{oneYearAfterTGE, 1079999999999715},
		{twoYearsAfterTGE, 1157965542048880},
		{threeYearsAfterTGE, 1232979823056239},
		{fourYearsAfterTGE, 1304866360072480},
	}
	for _, tc := range testCases {
		got := TotalSupply(tc.time)
		assert.Equal(t, tc.want, got)
	}
}

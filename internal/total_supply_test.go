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
		{twoYearsAfterTGE, 1151643489533809},
		{threeYearsAfterTGE, 1180266844240616},
		{fourYearsAfterTGE, 1205860930757792},
	}
	for _, tc := range testCases {
		got := TotalSupply(tc.time)
		assert.Equal(t, tc.want, got)
	}
}

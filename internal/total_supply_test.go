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
		{twoYearsAfterTGE, 1151791486153206},
		{threeYearsAfterTGE, 1205521395175910},
		{fourYearsAfterTGE, 1257990046704536},
	}
	for _, tc := range testCases {
		got := TotalSupply(tc.time)
		assert.Equal(t, tc.want, got)
	}
}

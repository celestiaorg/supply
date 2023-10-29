package supply

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAvailableSupply(t *testing.T) {
	type testCase struct {
		time time.Time
		want int64
	}
	testCases := []testCase{
		{beforeTGE, 0},
		{TGE, 266_986_177_750_000},
		{oneDayAfterTGE, 267_388_879_606_848},
		{oneYearAfterTGE, 592_531_409_126_425},
		{twoYearsAfterTGE, 1_032_535_327_799_419},
	}
	for _, tc := range testCases {
		t.Run(tc.time.String(), func(t *testing.T) {
			got := AvailableSupply(tc.time)
			require.Equal(t, tc.want, got)
		})
	}
}

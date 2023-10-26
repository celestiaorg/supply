package supply

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCirculating(t *testing.T) {
	type testCase struct {
		time time.Time
		want int64
	}
	testCases := []testCase{
		{time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), 0},
	}
	for _, tc := range testCases {
		t.Run(tc.time.String(), func(t *testing.T) {
			got := Circulating(tc.time)
			require.Equal(t, tc.want, got)
		})
	}
}

package supply

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCirculatingSupply(t *testing.T) {
	type testCase struct {
		name string
		time time.Time
		want int64
	}
	testCases := []testCase{
		{"a time before genesis", time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), 0},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CirculatingSupply(tc.time)
			require.Equal(t, tc.want, got)
		})
	}
}

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
		{TGE, initialTotalSupplyInUtia},
	}

	for _, tc := range testCases {
		got := TotalSupply(tc.time)
		assert.Equal(t, tc.want, got)
	}
}

package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatTia(t *testing.T) {
	type testCase struct {
		utia int64
		want string
	}
	testCases := []testCase{
		{0, "0.000000"},
		{1, "0.000001"},
		{utiaPerTia, "1.000000"},
		{utiaPerTia + 1, "1.000001"},
		{publicAllocationGenesis, "74057350.000000"},
		{initialTotalSupplyInUtia, "1000000000.000000"}, // 1 billion TIA
	}
	for _, tc := range testCases {
		got := FormatTia(tc.utia)
		assert.Equal(t, tc.want, got)
	}
}

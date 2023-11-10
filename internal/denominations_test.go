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
		{0, "0"},
		{1, "0.000001"},
		{utiaPerTia, "1"},
		{utiaPerTia + 1, "1.000001"}, // this test case fails
	}
	for _, tc := range testCases {
		got := FormatTia(tc.utia)
		assert.Equal(t, tc.want, got)
	}
}

package internal

import (
	"math/big"
)

const (
	utiaPerTia = 1_000_000 // 1 tia = 1 million utia
)

// FormatTia converts utia to TIA
func FormatTia(utia int64) string {
	x := newBigFloat(utia)
	y := newBigFloat(utiaPerTia)
	tia := new(big.Float)
	tia.Quo(x, y)
	tiaString := new(big.Float).SetPrec(7).SetMode(big.ToZero).Set(tia).Text('f', -1)
	return tiaString
}

func newBigFloat(x int64) *big.Float {
	bigFloat := new(big.Float)
	bigFloat.SetInt64(x)
	return bigFloat
}

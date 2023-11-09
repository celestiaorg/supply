package internal

import "fmt"

const (
	utiaPerTia = 1_000_000 // 1 tia = 1 million utia
)

// FormatTia converts utia to TIA
func FormatTia(utia int64) string {
	tia := float64(utia) / utiaPerTia
	return fmt.Sprintf("%f", tia)
}

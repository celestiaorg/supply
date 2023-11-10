package internal

import (
	"fmt"
)

const (
	// utiaPerTia is the number of utia in one TIA. One tia = one million utia.
	utiaPerTia = 1_000_000
	// precision is the number of decimal places to show for display values of
	// TIA. This level of precision ensures that the smallest unit of TIA (one
	// utia) is always visible when displaying values as TIA.
	precision = 6
)

// FormatTia converts utia to TIA
func FormatTia(utia int64) string {
	result := float64(utia) / float64(utiaPerTia)
	return fmt.Sprintf("%.*f", precision, result)
}

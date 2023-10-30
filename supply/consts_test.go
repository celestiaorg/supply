package supply

import (
	"fmt"
	"testing"
)

func Test_printDates(t *testing.T) {
	fmt.Printf("TGE %v\n", TGE)
	fmt.Printf("oneYearAfterTGE %v\n", oneYearAfterTGE)
	fmt.Printf("twoYearsAfterTGE %v\n", twoYearsAfterTGE)
}

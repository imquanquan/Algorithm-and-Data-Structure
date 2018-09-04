package problem009

import (
	"math"
	"fmt"
)

func PrintMax(x int)  {
	for a := 0; a < int(math.Pow10(x)); a++ {
		fmt.Println(a)
	}
}

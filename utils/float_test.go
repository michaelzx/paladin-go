package utils

import (
	"fmt"
	"testing"
)

func TestExamples(t *testing.T) {
	x := 12.3456
	fmt.Println(FloatRoundUp(x, 1))      // 12.4
	fmt.Println(FloatRoundUp(x, 2))      // 12.35
	fmt.Println(FloatRoundUp(x, 3))      // 12.346
	fmt.Println(FloatRoundDown(x, 1))    // 12.3
	fmt.Println(FloatRoundDown(x, 2))    // 12.34
	fmt.Println(FloatRoundDown(x, 3))    // 12.345
	fmt.Println(FloatRoundNearest(x, 1)) // 12.3
	fmt.Println(FloatRoundNearest(x, 2)) // 12.35
	fmt.Println(FloatRoundNearest(x, 3)) // 12.346
}

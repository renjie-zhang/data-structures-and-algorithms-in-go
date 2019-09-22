package FibonacciFind

import (
	"fmt"
	"testing"
)

func TestFibonacciFind(t *testing.T) {
	array := [6]int{1, 2, 45, 56, 67, 89}
	index := FinbonacciFind(&array, 89)
	fmt.Printf("找到了，index = %d\n", index)
}

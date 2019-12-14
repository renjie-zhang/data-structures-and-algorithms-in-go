package QuickSort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	array := []int{34, 2, 65, 32, 9}
	fmt.Printf("原数组为%d", array)
	fmt.Println()
	QuickSort(0, len(array)-1, array)
	fmt.Printf("排序过后数组为%d", array)
	fmt.Println()
}

/*
=== RUN   TestQuickSort
原数组为[34 2 65 32 9]
排序过后数组为[2 9 32 34 65]
--- PASS: TestQuickSort (0.00s)
*/

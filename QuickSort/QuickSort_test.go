package QuickSort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	array := [5]int{34, 2, 65, 32, 9}
	fmt.Printf("原数组为%d", &array)
	fmt.Println()
	QuickSort(0, len(array)-1, &array)
	fmt.Printf("排序过后数组为%d", &array)
	fmt.Println()
}

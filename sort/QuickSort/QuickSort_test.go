package QuickSort

import (
	"fmt"
	"testing"
)

//快速排序测试
func TestQuickSort(t *testing.T) {
	array := []int{34, 2, 65, 32, 9}
	temp := QuickSort(array)
	fmt.Println(temp)
}

/*
=== RUN   TestQuickSort
原数组为[34 2 65 32 9]
排序过后数组为[2 9 32 34 65]
--- PASS: TestQuickSort (0.00s)
*/

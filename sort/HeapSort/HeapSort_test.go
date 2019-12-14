package HeapSort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	array := []int{4, 6, 8, 5, 9, 2}
	fmt.Println("排序前：", array)
	HeapSort(array)
	fmt.Println("排序后：", array)
}

/*
=== RUN   TestHeapSort
排序前： [4 6 8 5 9 2]
排序后： [9 8 5 4 2 0]
--- PASS: TestHeapSort (0.00s)
PASS
*/

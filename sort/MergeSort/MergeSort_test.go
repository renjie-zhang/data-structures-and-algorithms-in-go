package MergeSort

import (
	"fmt"
	"testing"
)

//归并排序测试
func TestMergeSort(t *testing.T) {
	arr := []int{23, 4, 90, 3, 56}
	temp := MergeSort(arr)
	fmt.Println(temp)
}

/*
=== RUN   TestMergeSort
Test passed
--- PASS: TestMergeSort (0.00s)
*/

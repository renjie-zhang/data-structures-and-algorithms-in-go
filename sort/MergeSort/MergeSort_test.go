package MergeSort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := [5]int{12, -54, 4, 2, 90}
	temp := [5]int{}
	fmt.Printf("排序前%d\n", arr)
	MergeSort(&arr, 0, len(arr)-1, &temp)
	fmt.Printf("排序后%d\n", arr)
}

package selectsort

import "testing"

//选择排序测试
func TestSelectSort(t *testing.T) {
	arr := []int{23, 4, 90, 3, 56}
	SelectSort(arr)
}

/*
=== RUN   TestSelectSort
排序前： [23 4 90 3 56]
排序之后： [3 4 23 56 90]
--- PASS: TestSelectSort (0.00s)
*/

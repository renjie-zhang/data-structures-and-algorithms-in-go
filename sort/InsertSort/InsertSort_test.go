package InsertSort

import "testing"

//插入排序测试
func TestInsertSort(t *testing.T) {
	arr := []int{23, 4, 90, 3, 56}
	InsertSort(arr)
}

/*
=== RUN   TestInsertSort
排序之前的数组[23 4 90 3 56]
排序之后的数组[3 4 23 56 90]
--- PASS: TestInsertSort (0.00s)
*/

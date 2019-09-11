package SelectSort

import "testing"

func TestSelectSort(t *testing.T) {
	arr := [5]int{23, 4, 90, 3, 56}
	SelectSort(&arr)
}

/*
=== RUN   TestSelectSort
原数组为&[23 4 90 3 56]
排序过后数组为&[90 56 23 4 3]
--- PASS: TestSelectSort (0.00s)
PASS
*/

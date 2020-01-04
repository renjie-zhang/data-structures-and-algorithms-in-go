package ShellSort

import "testing"

//希尔排序
func TestShellSort(t *testing.T) {
	arr := []int{23, 4, 90, 3, 56}
	ShellSort(arr)
}

/*
=== RUN   TestShellSort
排序前： [23 4 90 3 56]
排序之后： [3 4 23 56 90]
--- PASS: TestShellSort (0.00s)
PASS
*/

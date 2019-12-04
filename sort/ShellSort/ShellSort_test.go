package ShellSort

import "testing"

func TestShellSort(t *testing.T) {
	arr := [5]int{23, 4, 90, 3, 56}
	ShellSort(&arr)
}

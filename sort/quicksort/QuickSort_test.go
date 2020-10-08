package quicksort

import (
	"testing"
)

//快速排序测试
func TestQuickSort(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{2, 1, 2}, []int{1, 2, 2}},
		{[]int{1, 2, 4, 3, 6, 5}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{6, 2, 4, 3, 1, 5}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		QuickSort(tt.in, 0, len(tt.in)-1)
		//Equal(t, tt.expected, tt.in)
	}
}

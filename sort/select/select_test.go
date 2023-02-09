package _select

import "testing"

// 选择排序测试
func TestSort(t *testing.T) {
	arr := []int{23, 4, 90, 3, 56}
	sort(arr)
}

func BenchmarkSort(b *testing.B) {
	arr := []int{23, 4, 90, 3, 56}
	for i := 0; i < b.N; i++ {
		sort(arr)
	}
}

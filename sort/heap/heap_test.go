package heap

import (
	"fmt"
	"testing"
)

// 堆排序测试
func TestSort(t *testing.T) {
	array := []int{4, 6, 8, 5, 9, 2}
	fmt.Println("排序前：", array)
	sort(array)
	fmt.Println("排序后：", array)
}

func BenchmarkSort(b *testing.B) {
	array := []int{4, 6, 8, 5, 9, 2}
	for i := 0; i < b.N; i++ {
		sort(array)
	}
}

package quick

import (
	"fmt"
	"testing"
)

// 快速排序测试
func TestSort(t *testing.T) {
	arr := []int{23, 4, 90, 3, 56}
	quick(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{23, 4, 90, 3, 56}
		quick(arr, 0, len(arr)-1)
	}
}

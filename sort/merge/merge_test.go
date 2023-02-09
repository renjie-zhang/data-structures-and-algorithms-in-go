package merge

import (
	"fmt"
	"testing"
)

// 归并排序测试
func TestSort(t *testing.T) {
	arr := []int{23, 4, 90, 3, 56}
	temp := sort(arr)
	fmt.Println(temp)
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{23, 4, 90, 3, 56}
		sort(arr)
	}
}

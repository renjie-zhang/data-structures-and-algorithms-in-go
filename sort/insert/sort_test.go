package insert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 插入排序测试
func TestSort(t *testing.T) {
	arr := []int{23, 4, 90, 3, 56}
	result := sort(arr)
	array2 := []int{3, 4, 23, 56, 90}
	assert.Equal(t, result, array2)
}

func BenchmarkSort(b *testing.B) {
	arr := []int{23, 4, 90, 3, 56}
	for i := 0; i < b.N; i++ {
		sort(arr)
	}
}

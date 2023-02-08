package bubble

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 测试冒泡排序
func TestSort(t *testing.T) {
	array1 := []int{2, 5, 10, 26, 57}
	array2 := []int{2, 5, 10, 26, 57}
	var result = Sort(array1)
	fmt.Println("排序后", result)
	assert.Equal(t, result, array2)

}

func BenchmarkSort(b *testing.B) {
	array1 := []int{2, 5, 10, 26, 57}
	for i := 0; i < b.N; i++ {
		Sort(array1)
	}
}

/**
cpu: Intel(R) Core(TM) i5-9600K CPU @ 3.70GHz
BenchmarkSort
BenchmarkSort-6         69818994                16.75 ns/op
PASS
*/

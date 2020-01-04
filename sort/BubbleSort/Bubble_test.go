package BubbleSort

import (
	"testing"
)

//测试冒泡排序
func TestBubble(t *testing.T) {
	array1 := []int{2, 5, 10, 26, 57}
	BubbleSort(array1)
	array2 := []int{99, 45, 23, 45, 23}
	BubbleSort(array2)
}

/*
=== RUN   TestBubble
排序前： [2 5 10 26 57]
提前退出啦
排序后： [2 5 10 26 57]
排序前： [99 45 23 45 23]
排序后： [23 23 45 45 99]
--- PASS: TestBubble (0.00s)
PASS
*/

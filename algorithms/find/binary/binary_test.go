package BinaryFind

import (
	"testing"
)

func TestBinaryFind(test *testing.T) {
	var array = []int{12, 23, 34, 45, 56, 67}
	if BinaryFind(array, 0, len(array)-1, 45) == 3{
		test.Log("通过测试")
	}else {
		test.Error("未通过测试")
	}
}

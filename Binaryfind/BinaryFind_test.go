package BinaryFind

import (
	"testing"
)

func TestBinaryFind(test *testing.T) {
	var array = [6]int{12, 23, 34, 45, 56, 67}
	BinaryFind(&array, 0, len(array)-1, 45)
}

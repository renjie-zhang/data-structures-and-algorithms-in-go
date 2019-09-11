package InsertSort

import "testing"

func TestInsertSort(t *testing.T) {
	arr := [5]int{23, 4, 90, 3, 56}
	InsertSort(&arr)
}

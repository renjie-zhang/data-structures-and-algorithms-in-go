package InsertValueFind

import (
	"fmt"
	"testing"
)

func TestInsertValueFind(t *testing.T) {
	var array [100]int
	for i := 0; i < 100; i++ {
		array[i] = i + 1
	}
	value := InsertValueFind(&array, 0, len(array)-1, 58)
	fmt.Printf("findValue的下标是%d\n", value)
}

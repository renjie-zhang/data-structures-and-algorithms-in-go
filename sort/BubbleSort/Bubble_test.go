package BubbleSort

import "testing"

func TestBubble(t *testing.T) {
	array := [5]int{45, 65, 2, 3, 7}
	BubbleSort(&array)
}

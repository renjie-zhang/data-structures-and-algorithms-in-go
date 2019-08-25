package SingleLinkedList

import "testing"

func TestSingleLinkedList(t *testing.T) {
	data := &Node{data: 0}
	single := &SingleLinkedList{head: data}
	single.InsertLast(45)
	single.InsertFirst(90)
	single.Print()
}

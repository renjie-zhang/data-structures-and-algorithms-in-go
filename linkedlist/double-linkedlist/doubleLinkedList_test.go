package DoubleLinkedList

import (
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	node := Node{data: 0}
	list := DoubleLinkedList{head: &node}
	list.InsertLast(23)
	list.InsertLast(32)
	t.Log(list)
}


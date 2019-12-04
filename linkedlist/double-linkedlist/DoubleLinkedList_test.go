package DoubleLinkedList

import (
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	node := Node{data: 0}
	list := LinkedList{head: &node}
	list.InsertLast(23)
	list.InsertLast(32)
	t.Log(list)
}

/*
=== RUN   TestDoubleLinkedList
--- PASS: TestDoubleLinkedList (0.00s)
    DoubleLinkedList_test.go:12: {0xc000048480 0xc0000484c0}
PASS
*/

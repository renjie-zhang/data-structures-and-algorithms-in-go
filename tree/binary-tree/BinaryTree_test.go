package BinaryTree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	first := &Node{No: 1, Name: "a"}
	second := &Node{No: 2, Name: "b"}
	three := &Node{No: 3, Name: "c"}
	four := &Node{No: 4, Name: "d"}
	five := &Node{No: 5, Name: "e"}
	six := &Node{No: 6, Name: "f"}

	var binarytree BinaryTree

	binarytree.SetRoot(first)

	first.Left = three
	first.Right = six
	three.Left = second
	three.Right = four
	six.Left = five

	fmt.Println("前序遍历")
	binarytree.PreOrder()

	fmt.Println("中序遍历")
	binarytree.InfixOrder()

	fmt.Println("后序遍历")
	binarytree.PostOrder()

}

/*
输出：
	=== RUN   TestBinaryTree
	前序遍历
	no:1  name:a
	no:3  name:c
	no:2  name:b
	no:4  name:d
	no:6  name:f
	no:5  name:e
	中序遍历
	no:2  name:b
	no:3  name:c
	no:4  name:d
	no:1  name:a
	no:5  name:e
	no:6  name:f
	后序遍历
	no:2  name:b
	no:3  name:c
	no:4  name:d
	no:1  name:a
	no:5  name:e
	no:6  name:f
	--- PASS: TestBinaryTree (0.10s)
	PASS
*/

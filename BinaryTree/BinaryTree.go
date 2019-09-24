package BinaryTree

import "fmt"

type Node struct {
	No    int
	Name  string
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	root *Node
}

//二叉树前序遍历
func (this *BinaryTree) PreOrder() {
	if this.root != nil {
		this.root.preOrder()
	} else {
		fmt.Println("当前二叉树为空，无法进行遍历")
	}
}

//二叉树中序遍历
func (this *BinaryTree) InfixOrder() {
	if this.root != nil {
		this.root.infixOrder()
	} else {
		fmt.Println("当前二叉树为空，无法进行遍历")
	}
}

//二叉树后序遍历
func (this *BinaryTree) PostOrder() {
	if this.root != nil {
		this.root.infixOrder()
	} else {
		fmt.Println("当前二叉树为空，无法进行遍历")
	}
}

//设置根节点
func (this *BinaryTree) SetRoot(root *Node) {
	this.root = root
}

//前序遍历
func (this *Node) preOrder() {
	this.show()
	if this.Left != nil {
		this.Left.preOrder()
	}
	if this.Right != nil {
		this.Right.preOrder()
	}
}

//中序遍历
func (this *Node) infixOrder() {
	if this.Left != nil {
		this.Left.infixOrder()
	}
	this.show()
	if this.Right != nil {
		this.Right.infixOrder()
	}
}

//后序遍历
func (this *Node) postOrder() {
	if this.Left != nil {
		this.Left.postOrder()
	}
	if this.Right != nil {
		this.Right.postOrder()
	}
	this.show()
}

func (this *Node) show() {
	fmt.Printf("no:%d  name:%v", this.No, this.Name)
	fmt.Println()
}

package BinaryTree

import (
	"container/list"
	"fmt"
)

type Node struct {
	No    int
	Name  string
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	root *Node
}

type queue struct {
	List *list.List
}

func (q queue) pop() interface{}{
	if element := q.List.Front();element != nil{
		q.List.Remove(element)
		return element.Value
	}
	return nil
}

func (q queue) push(value interface{})  {
	q.List.PushBack(value)
}

// 层次遍历(使用队列)
func LevelOrder(tree *Node){
	// 作为换层的flag的一个temp
	var tempFlag *Node
	// 作为换层的flag
	change := tree
	// 现在的层数
	level := 1
	qu := queue{List: list.New()}

	fmt.Println(fmt.Sprintf("this is %d level",level))
	qu.push(tree)
	for qu.List.Len() > 0 {
		node := qu.pop().(*Node)
		// 当左子结点不为空时，将结点放入到队列中
		if node.Left != nil{
			qu.push(node.Left)
			tempFlag = node.Left
		}
		// 当左子结点不为空时，将结点放入到队列中
		if node.Right != nil{
			qu.push(node.Right)
			tempFlag = node.Right
		}
		node.show()
		// 如果当前结点与层次flag相同，那么将进行遍历树的下一层
		if change == node && (node.Left != nil|| node.Right != nil){
			change = tempFlag
			level ++
			fmt.Println(fmt.Sprintf("this is %d level",level))
		}
	}
}



// 二叉树前序遍历(递归)
func (this *BinaryTree) PreOrder() {
	if this.root != nil {
		this.root.preOrder()
	} else {
		fmt.Println("当前二叉树为空，无法进行遍历")
	}
}

// 二叉树中序遍历(递归)
func (this *BinaryTree) InfixOrder() {
	if this.root != nil {
		this.root.infixOrder()
	} else {
		fmt.Println("当前二叉树为空，无法进行遍历")
	}
}

//二叉树后序遍历(递归)
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

//前序遍历(递归)
func (this *Node) preOrder() {
	this.show()
	if this.Left != nil {
		this.Left.preOrder()
	}
	if this.Right != nil {
		this.Right.preOrder()
	}
}

//中序遍历(递归)
func (this *Node) infixOrder() {
	if this.Left != nil {
		this.Left.infixOrder()
	}
	this.show()
	if this.Right != nil {
		this.Right.infixOrder()
	}
}

//后序遍历(递归)
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


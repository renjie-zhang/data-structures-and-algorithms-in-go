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

// stack
type stack struct {
	List *list.List
}

func (s *stack) pop() interface{} {
	if element := s.List.Back(); element != nil {
		s.List.Remove(element)
		return element.Value
	}
	return nil
}
func (s *stack) push(value interface{}) {
	s.List.PushBack(value)
}

// queue
type queue struct {
	List *list.List
}

func (q queue) pop() interface{} {
	if element := q.List.Front(); element != nil {
		q.List.Remove(element)
		return element.Value
	}
	return nil
}

func (q queue) push(value interface{}) {
	q.List.PushBack(value)
}

// 非递归的先序遍历
func (node *Node) preOrder() {
	s := stack{List: list.New()}
	s.push(node)

	element := s.pop()
	for element != nil {
		node, _ := element.(*Node)
		node.show()
		if right := node.Right; right != nil {
			s.push(right)
		}
		if left := node.Left; left != nil {
			s.push(left)
		}
		element = s.pop()
	}
}


// 非递归的中序遍历
func (node *Node) infixOrder() {
	s := stack{List: list.New()}
	current := node
	for s.List.Len() > 0 || current != nil {
		if current != nil {
			s.push(current)
			current = current.Left
		} else {
			current = s.pop().(*Node)
			current.show()
			current = current.Right
		}
	}

}

// 非递归方式后序遍历
func (node *Node) postOrder() {
	s1, s2 := stack{List: list.New()}, stack{List: list.New()}
	s1.push(node)

	for s1.List.Len() > 0 {
		element := s1.pop().(*Node)
		s2.push(element)
		if element.Left != nil {
			s1.push(element.Left)
		}
		if element.Right != nil {
			s1.push(element.Right)
		}
	}
	for s2.List.Len() > 0 {
		element := s2.pop().(*Node)
		element.show()
	}
}

// 层次遍历(使用队列)
func LevelOrder(tree *Node) {
	// 作为换层的flag的一个temp
	var tempFlag *Node
	// 作为换层的flag
	change := tree
	// 现在的层数
	level := 1
	qu := queue{List: list.New()}

	fmt.Println(fmt.Sprintf("this is %d level", level))
	qu.push(tree)
	for qu.List.Len() > 0 {
		node := qu.pop().(*Node)
		// 当左子结点不为空时，将结点放入到队列中
		if node.Left != nil {
			qu.push(node.Left)
			tempFlag = node.Left
		}
		// 当左子结点不为空时，将结点放入到队列中
		if node.Right != nil {
			qu.push(node.Right)
			tempFlag = node.Right
		}
		node.show()
		// 如果当前结点与层次flag相同，那么将进行遍历树的下一层
		if change == node && (node.Left != nil || node.Right != nil) {
			change = tempFlag
			level++
			fmt.Println(fmt.Sprintf("this is %d level", level))
		}
	}
}

// 二叉树前序遍历(递归)
func (this *BinaryTree) PreOrderRecu() {
	if this.root != nil {
		this.root.preOrderRecu()
	} else {
		fmt.Println("当前二叉树为空，无法进行遍历")
	}
}

// 二叉树中序遍历(递归)
func (this *BinaryTree) InfixOrderRecu() {
	if this.root != nil {
		this.root.infixOrderRecu()
	} else {
		fmt.Println("当前二叉树为空，无法进行遍历")
	}
}

//二叉树后序遍历(递归)
func (this *BinaryTree) PostOrderRecu() {
	if this.root != nil {
		this.root.infixOrderRecu()
	} else {
		fmt.Println("当前二叉树为空，无法进行遍历")
	}
}

//设置根节点
func (this *BinaryTree) SetRoot(root *Node) {
	this.root = root
}

//前序遍历(递归)
func (this *Node) preOrderRecu() {
	this.show()
	if this.Left != nil {
		this.Left.preOrderRecu()
	}
	if this.Right != nil {
		this.Right.preOrderRecu()
	}
}

//中序遍历(递归)
func (this *Node) infixOrderRecu() {
	if this.Left != nil {
		this.Left.infixOrderRecu()
	}
	this.show()
	if this.Right != nil {
		this.Right.infixOrderRecu()
	}
}

//后序遍历(递归)
func (this *Node) postOrderRecu() {
	if this.Left != nil {
		this.Left.postOrderRecu()
	}
	if this.Right != nil {
		this.Right.postOrderRecu()
	}
	this.show()
}

func (this *Node) show() {
	fmt.Printf("no:%d  name:%v", this.No, this.Name)
	fmt.Println()
}

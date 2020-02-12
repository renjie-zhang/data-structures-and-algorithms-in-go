package binary_search_tree

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// 查找
func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}

	if n.Key < key {
		return n.Right.Search(key)
	} else if n.Key > key {
		return n.Left.Search(key)
	}
	return true
}

// 插入二叉查找树
func (n *Node) Insert(key int) {
	if n.Key < key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	} else {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	}
}

// 删除
func (n *Node) Delete(key int) *Node {
	//通过key进行查找
	if n.Key < key {
		n.Right = n.Right.Delete(key)
	} else if n.Key > key {
		n.Left = n.Left.Delete(key)
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}
		min := n.Right.Min()
		n.Key = min
		n.Right = n.Right.Delete(min)
	}
	return n
}

//查找最小值
func (n *Node) Min() int {
	if n.Left == nil {
		return n.Key
	}
	return n.Left.Min()
}

// 查找最大值
func (n *Node) Max() int {
	if n.Right == nil {
		return n.Key
	}
	return n.Right.Max()
}

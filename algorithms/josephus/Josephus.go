package Josephus

import "fmt"

type Node struct {
	data int   //节点数据
	next *Node //next节点
}

//使用单链表进行解决问题
func getPosition_1(index int, total int) {
	head := &Node{data: 1, next: nil}
	prev := head
	for i := 1; i <= total; i++ {
		prev.next = &Node{data: i, next: nil}
		prev = prev.next
	}
	prev.next = head

	first := head
	second := head
	for first != first.next {
		count := 1
		for count != index {
			second = first
			first = first.next
			count++
		}
		second.next = first.next
		fmt.Printf("delete node from cirecle linkedlist : %d\n", first.data)
		first = second.next
	}
	fmt.Printf("Last person left standing is %d\n", prev.data)
}

//使用递归进行解决，此方法防止栈溢出
func getPosition_2(index int, total int) {
	result := 1
	for i := 2; i <= total; i++ {
		result = (result+index-1)%i + 1
	}
	fmt.Printf("Last person left standing is %d\n", result)
}

package SingleLinkedList

import (
	"Data-Structures-and-Algorithms-Go/linkedlist"
	"fmt"
)

func assertListImplementtation(){
	var _ linkedlist.List = (*List)(nil)
}
//链表
type List struct{
	head *element
	tail *element
	size int
}
//节点
type element struct{
	value interface{}
	next *element
}
//链表实例化
func New(values ...interface{}) *List{
	list := &List{}
	if len(values) >0{
		list.Add(values...)
	}
	return list
}

//添加节点
func (list *List) Add(values ...interface{}){
	for _,value := range values{
		newElement := &element{value:value}
		if list.size == 0{
			//如果链表为空则添加到头节点中
			list.head = newElement
			list.tail = newElement
		}else {
			//如果不为空，则添加到尾节点的下一节点
			list.tail.next = newElement
			list.tail = newElement
		}
		list.size++
	}
}


//node节点
type Node struct {
	data int
	next *Node
}

//单链表结构体
type SingleLinkedList struct {
	head *Node
}

//在头部插入
func (this *SingleLinkedList) InsertFirst(value int) {
	data := &Node{data: value}
	if this.head != nil {
		data.next = this.head
	}
	this.head = data
}

//在最后插入
func (this *SingleLinkedList) InsertLast(value int) {
	data := &Node{data: value}
	if this.head == nil {
		this.head = data
		return
	}
	current := this.head
	for current.next != nil {
		current = current.next
	}
	current.next = data
}

//通过values进行移除
func (this *SingleLinkedList) RemoveByValue(value int) bool {
	if this.head == nil {
		return false
	}
	if this.head.data == value {
		this.head = this.head.next
		return true
	}
	current := this.head
	for current.next != nil {
		if current.next.data == value {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

//通过index进行移除
func (this *SingleLinkedList) RemoveByIndex(index int) bool {
	if index < 0 {
		return false
	}
	if index == 1 {
		this.head = this.head.next
		return true
	}
	if this.head == nil {
		return false
	}
	current := this.head
	for u := 1; u < index; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	current.next = current.next.next
	return true
}

//搜索是否有value
func (this *SingleLinkedList) SearchValue(value int) bool {
	if this.head == nil {
		return false
	}
	current := this.head
	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}

//获取头节点
func (this *SingleLinkedList) GetFirst() (int, bool) {
	if this.head == nil {
		return 0, false
	}
	return this.head.data, true
}

//获取最后一个元素
func (this *SingleLinkedList) GetLast() (int, bool) {
	if this.head == nil {
		return 0, false
	}
	current := this.head
	for current.next != nil {
		current = current.next
	}
	return current.data, true
}

//获取节点的个数
func (this *SingleLinkedList) GetCount() int {
	count := 0
	if this.head == nil {
		return count
	}
	current := this.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

//输出链表返回一个数组
func (this *SingleLinkedList) GetItems() []int {
	var items []int
	current := this.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}
func (this *SingleLinkedList) Print() {
	current := this.head
	format := ""
	for current != nil {
		format += fmt.Sprintf("%+v", current.data)
		current = current.next
		if current != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}

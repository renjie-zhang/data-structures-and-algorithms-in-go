package SingleLinkedList

import (
	"fmt"
)

type SinglyLinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	data interface{}
	next *Node
}

// Add Node
func (s *SinglyLinkedList) Add(values ...interface{}) {
	for _, value := range values {
		newNode := &Node{data: value}
		if s.size == 0 {
			//如果链表为空则添加到头节点中
			s.head = newNode
			s.tail = newNode
		} else {
			//如果不为空，则添加到尾节点的下一节点
			s.tail.next = newNode
			s.tail = newNode
		}
		s.size++
	}
}

// New SinglyLinkedList
func New(values ...interface{}) *SinglyLinkedList {
	list := &SinglyLinkedList{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// InsertFirst Insert a Node to SinglyLinkedList
func (s *SinglyLinkedList) InsertFirst(value int) {
	// new Node
	data := &Node{data: value}
	if s.head != nil {
		data.next = s.head
	}
	s.head = data
}

func (s *SinglyLinkedList) InsertLast(value int) {
	data := &Node{data: value}
	if s.head == nil {
		s.head = data
		return
	}
	current := s.head
	for current.next != nil {
		current = current.next
	}
	current.next = data
}

func (s *SinglyLinkedList) RemoveByValue(value int) bool {
	if s.head == nil {
		return false
	}
	if s.head.data == value {
		s.head = s.head.next
		return true
	}
	current := s.head
	for current.next != nil {
		if current.next.data == value {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func (s *SinglyLinkedList) RemoveByIndex(index int) bool {
	if index < 0 {
		return false
	}
	if index == 1 {
		s.head = s.head.next
		return true
	}
	if s.head == nil {
		return false
	}
	current := s.head
	for u := 1; u < index; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	current.next = current.next.next
	return true
}

func (s *SinglyLinkedList) SearchValue(value int) bool {
	if s.head == nil {
		return false
	}
	current := s.head
	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}

func (s *SinglyLinkedList) GetFirst() (interface{}, bool) {
	if s.head == nil {
		return 0, false
	}
	return s.head.data, true
}

func (s *SinglyLinkedList) GetLast() (interface{}, bool) {
	if s.head == nil {
		return 0, false
	}
	current := s.head
	for current.next != nil {
		current = current.next
	}
	return current.data, true
}

func (s *SinglyLinkedList) GetCount() int {
	count := 0
	if s.head == nil {
		return count
	}
	current := s.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (s *SinglyLinkedList) GetItems() []interface{} {
	var items []interface{}
	current := s.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}
func (s *SinglyLinkedList) Print() {
	current := s.head
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

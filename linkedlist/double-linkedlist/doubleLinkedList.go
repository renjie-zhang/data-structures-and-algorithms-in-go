package DoubleLinkedList



type DoubleLinkedList struct {
	head *Node
	tail *Node
	count int
}

type Node struct {
	data int
	next *Node
	prev *Node
}


func (d *DoubleLinkedList) InsertFirst(i int) {
	data := &Node{data: i}
	if d.head != nil {
		d.head.prev = data
		data.next = d.head
	}
	d.head = data
	d.count ++
}

func (d *DoubleLinkedList) InsertLast(i int) {
	data := &Node{data: i}
	if d.head == nil {
		d.head = data
		d.tail = data
		return
	}
	if d.tail != nil {
		d.tail.next = data
		data.prev = d.tail
	}
	d.tail = data

}

func (d *DoubleLinkedList) RemoveByValue(i int) bool {
	if d.head == nil {
		return false
	}
	if d.head.data == i {
		d.head = d.head.next
		d.head.prev = nil
		return true
	}
	if d.tail.data == i {
		d.tail = d.tail.prev
		d.tail.next = nil
		return true
	}
	current := d.head
	for current.next != nil {
		if current.next.data == i {
			if current.next.next != nil {
				current.next.next.prev = current
			}
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false

}

func (d *DoubleLinkedList) RemoveByIndex(index int) bool {
	if d.head == nil {
		return false
	}
	if index < 0 {
		return false
	}
	if index == 0 {
		d.head.prev = nil
		d.head = d.head.next
		return true
	}
	current := d.head
	for j := 1; j < index; j++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	if current.next.next != nil {
		current.next.next.prev = current
	}
	current.next = current.next.next
	return true
}

func (d *DoubleLinkedList) GetSize() int {
	var count int
	current := d.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

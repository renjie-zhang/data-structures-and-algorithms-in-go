package linkedlist

import "Data-Structures-and-Algorithms-Go/utils"

type List interface {
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(values ...interface{})
	Contains(values ...interface{})
	Sort(comparator utils.Comparator)
	Swap(index1, index2 int)
	Insert(index int, values ...interface{})
	Set(index int, value interface{})
	Empty() bool
	Size() int
	Clear()
}

package heap

// sort 堆排序
func sort(array []int) {
	length := len(array)
	for i := length/2 - 1; i > 0; i-- {
		changeToBigHeap(array, i, length)
	}
	var temp int
	for i := length - 1; i > 0; i-- {
		temp, array[i] = array[i], temp
		changeToBigHeap(array, 0, i)
	}
}

// changeToBigHeap 将对应的i非叶子节点的数调整为大顶堆
func changeToBigHeap(array []int, i int, length int) {
	temp := array[i]
	for k := 2*i + 1; k < length; k = k*2 + 1 {
		if k+1 < length && array[k] < array[k+1] {
			k++
		}
		if array[k] > temp {
			array[i] = array[k]
			i = k
		} else {
			break
		}
	}
	array[i] = temp
}

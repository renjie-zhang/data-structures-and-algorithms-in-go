package HeapSort

func HeapSort(array *[6]int) {
	for i := len(array)/2 - 1; i > 0; i-- {
		ChangeToBigHeap(array, i, len(array))
	}
	var temp int
	for i := len(array) - 1; i > 0; i-- {
		temp, array[i] = array[i], temp
		ChangeToBigHeap(array, 0, i)
	}
}

//将对应的i非叶子节点的数调整为大顶堆
func ChangeToBigHeap(array *[6]int, i int, length int) {
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

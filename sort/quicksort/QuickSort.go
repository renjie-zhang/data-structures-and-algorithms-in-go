package quicksort

//快速排序
func QuickSort(arr []int) []int {
	partition(0, len(arr)-1, arr)
	return arr
}

func partition(left int, right int, arr []int) {

	l := left
	r := right
	pivot := arr[(l+r)/2]
	for l < r {
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		//交换
		arr[l], arr[r] = arr[r], arr[l]
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		partition(left, r, arr)
	}
	//向右递归
	if right > l {
		partition(l, right, arr)
	}

}

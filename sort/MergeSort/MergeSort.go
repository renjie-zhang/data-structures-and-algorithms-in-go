package MergeSort

func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

/*func MergeSort(arr []int, left int, right int, temp []int) {
	if left < right {
		mid := (left + right) / 2
		MergeSort(arr, left, mid, temp)
		MergeSort(arr, mid+1, right, temp)
		Merge(arr, left, right, mid, temp)
	}
}

//合并的方法
func Merge(arr []int, left int, right int, mid int, temp []int) {
	//初始化
	i := left
	j := mid + 1
	t := 0
	//先把左右两边的有序序列按照规则填充到temp数组中
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[t] = arr[i]
			i += 1
			t += 1
		} else {
			temp[t] = arr[j]
			j += 1
			t += 1
		}
	}
	//将剩余的数据全部填充到temp中
	for i <= mid {
		temp[t] = arr[i]
		t += 1
		i += 1
	}
	for j <= right {
		temp[t] = arr[j]
		j += 1
		t += 1
	}

	//将所有的数据拷贝会arr中
	t = 0
	tempLeft := left
	for tempLeft <= right {
		arr[tempLeft] = temp[t]
		t += 1
		tempLeft += 1
	}
}*/

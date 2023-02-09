package _select

// 选择排序
func sort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		minValue := arr[i]
		minIndex := i
		for j := i + 1; j < length; j++ {
			if minValue > arr[j] {
				minValue = arr[j]
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}

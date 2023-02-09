package quick

func quick(arr []int, low int, high int) {
	if low < high {
		j := partition(arr, low, high)
		quick(arr, low, j-1)
		quick(arr, j+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

package shell

// 希尔排序
func sort(arr []int) {
	n := len(arr)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && (arr[j] < arr[j-h]); j -= h {
				temp := arr[j]
				arr[j] = arr[j-h]
				arr[j-h] = temp
			}
		}
		h /= 3
	}
}

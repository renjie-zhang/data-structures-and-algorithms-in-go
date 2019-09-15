package ShellSort

import "fmt"

func ShellSort(arr *[5]int) {
	fmt.Printf("排序之前的数组%d", arr)
	fmt.Println()
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
	fmt.Printf("排序之后的数组%d", arr)
	fmt.Println()
}

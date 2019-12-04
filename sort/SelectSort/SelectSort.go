package SelectSort

import "fmt"

func SelectSort(arr *[5]int) {

	fmt.Printf("原数组为%d", arr)
	fmt.Println()
	for i := 0; i < len(arr)-1; i++ {
		maxValue := arr[i]
		maxIndex := i
		for j := i + 1; j < len(arr); j++ {
			if maxValue < arr[j] {
				maxValue = arr[j]
				maxIndex = j
			}
		}
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
	}
	fmt.Printf("排序过后数组为%d", arr)
	fmt.Println()
}

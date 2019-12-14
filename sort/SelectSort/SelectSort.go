package SelectSort

import "fmt"

func SelectSort(arr []int) {

	fmt.Println("排序前：", arr)
	for i := 0; i < len(arr)-1; i++ {
		minValue := arr[i]
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if minValue > arr[j] {
				minValue = arr[j]
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	fmt.Println("排序之后：", arr)
}

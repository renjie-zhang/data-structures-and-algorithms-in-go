package selectsort

import "fmt"

//选择排序
func SelectSort(arr []int) {

	fmt.Println("排序前：", arr)
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
	fmt.Println("排序之后：", arr)
}

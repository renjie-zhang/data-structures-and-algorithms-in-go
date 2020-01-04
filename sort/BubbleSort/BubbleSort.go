package BubbleSort

import "fmt"

//冒泡排序
func BubbleSort(arr []int) []int {
	fmt.Println("排序前：", arr)
	if len(arr) == 1 {
		fmt.Println("排序后：", arr)
		return arr
	}

	isExchang := false
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isExchang = true
			}
		}
		if !isExchang {
			fmt.Println("提前退出啦")
			break
		}
	}
	fmt.Println("排序后：", arr)
	return arr
}

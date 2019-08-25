package BinaryFind

import "fmt"

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {

	if leftIndex > rightIndex {
		fmt.Printf("找不到")
		return
	}

	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		//说明我们查找的数是再右边，因该在leftIndex <-->middel -1
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了,下标为%v", middle)
	}

}

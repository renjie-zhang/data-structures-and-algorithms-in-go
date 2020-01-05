package insertsort

import "fmt"

//插入排序
func InsertSort(arr []int) {
	fmt.Printf("排序之前的数组%d", arr)
	fmt.Println()
	length := len(arr)
	for i := 1; i < length; i++ {
		insertValue := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] > insertValue {
			//将数据后移
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		//插入数据
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertValue
		}
	}
	fmt.Printf("排序之后的数组%d", arr)
	fmt.Println()

}

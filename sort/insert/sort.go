package insert

// sort 插入排序
func sort(arr []int) []int {
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
	return arr
}

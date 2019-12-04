package InsertValueFind

func InsertValueFind(array *[100]int, left int, right int, findValue int) int {
	if right < left || findValue < array[0] || findValue > array[len(array)-1] {
		return -1
	}
	//插值查找公式
	mid := left + (right-left)*(findValue-array[left])/(array[right]-array[left])

	if findValue < array[mid] {
		InsertValueFind(array, left, mid-1, findValue)
	} else if findValue > array[mid] {
		InsertValueFind(array, mid+1, right, findValue)
	} else {
		return mid
	}
	return -1

}

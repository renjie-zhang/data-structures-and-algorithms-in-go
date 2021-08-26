package BinaryFind

func BinaryFind(arr []int, leftIndex int, rightIndex int, findVal int) int {
	if rightIndex >= leftIndex{
		mid := leftIndex + (rightIndex - leftIndex) /2
		if arr[mid] == findVal{
			return mid
		} else if arr[mid] > findVal{
			return BinaryFind(arr,leftIndex,mid - 1,findVal)
		} else if arr[mid] < arr[rightIndex]{
			return BinaryFind(arr,mid+1,rightIndex,findVal)
		}
	}
	return -1
}

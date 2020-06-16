package quicksort

import (
	"reflect"
	"testing"
)

func QuickSort(arr []int,low int,high int){
	if low > high{
		return
	}
	j := partition(arr,low,high)
	QuickSort(arr,low,j-1)
	QuickSort(arr,j+1,high)
}

func partition(arr []int,low int, high int, ) int {
	i := low
	j := high
	temp := arr[low]
	for i <= j{
		for i<high && arr[i] > temp{
			i++
		}
		for j > low && arr[j] <= temp{
			j--
		}
		if i < j{
			arr[i],arr[j] = arr[j],arr[i]
		}else {
			arr[i],arr[high] = arr[high],arr[i]
		}
	}

	return i
}

func Equal(t *testing.T, expected, result interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("should be %v instead of %v", expected, result)
	}
}

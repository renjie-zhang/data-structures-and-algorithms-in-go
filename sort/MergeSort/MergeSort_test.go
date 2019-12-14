package MergeSort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 初始化 array with num
func initArray(num int) []int {
	if num < 1 {
		panic("num must bigger than 1")
	}

	arr := make([]int, num)
	middle := num / 2
	// fmt.Println("middle :", middle)
	for i, _ := range arr {
		arr[i] = i - middle
	}
	return arr
}

// 比较 sort 前后数组是否相同
func compare(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, _ := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func TestMergeSort(t *testing.T) {
	num := 500
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	src := initArray(num)
	dest := make([]int, len(src))
	perm := r.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	result := MergeSort(dest)
	if compare(src, result) {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed")
	}
}

/*
=== RUN   TestMergeSort
Test passed
--- PASS: TestMergeSort (0.00s)
*/

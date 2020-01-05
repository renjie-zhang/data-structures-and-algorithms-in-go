package eightqueue

import (
	"fmt"
	"math"
)

//皇后数量
var max int = 8
//暂存皇后放置的temp数组
var tempArray [8]int
//记录解法
var count int

func Check(n int){
	if n == max{
		printResult()
		return
	}
	//一次放入，并判断是否冲突
	for i := 0;i<max ;i++  {
		tempArray[n] = i
		if judge(n) {
			Check(n+1)
		}
	}

}

//判断是否冲突
func judge(n int) bool {
	for i := 0;i<n ;i++  {
		if tempArray[i] == tempArray[n] || math.Abs(float64(n-i)) == math.Abs(float64(tempArray[n]-tempArray[i])){
			return false
		}
	}
	return true
}


//打印最终结果
func printResult(){
	count++
	length := len(tempArray)
	fmt.Printf("这是第%d解法: ",count)
	for i := 0;i< length;i++  {
		fmt.Print(tempArray[i])
	}
	fmt.Println()

}

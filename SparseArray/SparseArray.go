/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-08-19 15:30:20
 * @LastEditTime: 2019-08-19 18:38:09
 */
package SparseArray

import "fmt"

func SparseArray() {
	//原数组init
	var oldArray [10][10]int
	oldArray[1][2] = 1
	oldArray[2][3] = 2
	oldArray[4][5] = 3
	//打印原数组
	fmt.Println("打印原数组")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d\t", oldArray[i][j])
		}
		fmt.Println()
	}

	//根据原数组创建稀疏数组
	//首先遍历出所有不为0的值
	fmt.Println("创建稀疏数值")
	var count int
	for i := 0; i < 0; i++ {
		for j := 0; j < 10; j++ {
			if oldArray[i][j] != 0 {
				count++
			}
		}
	}
	fmt.Printf("不为0的个数是：%d\n", count)
	var SparseArray [count][3]int
	SparseArray[0][0] = 10
	SparseArray[0][1] = 10
	SparseArray[0][2] = count
}

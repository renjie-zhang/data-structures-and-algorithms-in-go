/*
 * @Descripttion:稀疏数组
 * @version:v1
 * @Author: renjie.zhang
 * @Date: 2019-08-19 15:30:20
 * @LastEditTime: 2019-08-20 19:47:41
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
	count := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if oldArray[i][j] != 0 {
				count = count + 1
			}
		}
	}
	fmt.Printf("稀疏数组中不为0的个数是：%d\n", count)
	//待改进-->使用切片
	var SparseArray [4][4]int
	SparseArray[0][0] = 10
	SparseArray[0][1] = 10
	SparseArray[0][2] = count
	//循环遍历原二维数组
	var flag int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if oldArray[i][j] != 0 {
				flag++
				SparseArray[flag][0] = i
				SparseArray[flag][1] = j
				SparseArray[flag][2] = oldArray[i][j]
			}
		}
	}

	//输出稀疏数组
	fmt.Println("输出稀疏数组")
	for i := 0; i < len(SparseArray); i++ {
		fmt.Printf("%d\t%d\t%d\t\n", SparseArray[i][0], SparseArray[i][1], SparseArray[i][2])
		fmt.Println()
	}

	//稀疏数组转换为普通数组
	var newArray [10][10]int
	for a := 1; a < len(SparseArray); a++ {
		newArray[SparseArray[a][0]][SparseArray[a][1]] = SparseArray[a][2]
	}
	//输出新数组
	fmt.Println("输出新数组")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d\t", newArray[i][j])
		}
		fmt.Println()
	}
	/*
		=== RUN   TestSparseArray
		打印原数组
		0       0       0       0       0       0       0       0       0       0
		0       0       1       0       0       0       0       0       0       0
		0       0       0       2       0       0       0       0       0       0
		0       0       0       0       0       0       0       0       0       0
		0       0       0       0       0       3       0       0       0       0
		0       0       0       0       0       0       0       0       0       0
		0       0       0       0       0       0       0       0       0       0
		0       0       0       0       0       0       0       0       0       0
		0       0       0       0       0       0       0       0       0       0
		0       0       0       0       0       0       0       0       0       0
		稀疏数组中不为0的个数是：3
		输出稀疏数组
		10      10      3

		1       2       1

		2       3       2

		4       5       3

		输出新数组
		0       0       0       0       0       0       0       0       0       0
		0       0       1       0       0       0       0       0       0       0
		0       0       0       2       0       0       0       0       0       0
		0       0       0       0       0       0       0       0       0       0
		0       0       0       0       0       3       0       0       0       0
		0       0       0       0       0       0       0       0       0       0
		0       0       0       0       0       0       0       0       0       0
		0       0       0       0       0       0       0       0       0       0
		0       0       0       0       0       0       0       0       0       0
		0       0       0       0       0       0       0       0       0       0
		--- PASS: TestSparseArray (0.01s)
		PASS
	*/
}

/*
 * @Descripttion:
 * @version: 冒泡排序
 * @Author: renjie.zhang
 * @Date: 2019-07-29 22:14:32
 * @LastEditTime: 2019-12-08 10:34:32
 */
package BubbleSort

import "fmt"

/*
冒泡排序
*/
func BubbleSort(arr []int) {
	fmt.Println("排序前：", arr)
	if len(arr) == 1 {
		fmt.Println("排序后：", arr)
	}

	isExchang := false
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isExchang = true
			}
		}
		if !isExchang {
			fmt.Println("提前退出啦")
			break
		}
	}
	fmt.Println("排序后：", arr)
}

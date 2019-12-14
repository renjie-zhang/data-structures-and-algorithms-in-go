/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-08-25 15:19:31
 * @LastEditTime: 2019-12-08 10:32:40
 */
package BubbleSort

import "testing"

func TestBubble(t *testing.T) {
	array1 := []int{2, 3, 5, 6, 7}
	BubbleSort(array1)

	array2 := []int{99, 45, 23, 45, 23}
	BubbleSort(array2)
}

/*
=== RUN   TestBubble
排序前： [2 3 5 6 7]
提前退出啦
排序后： [2 3 5 6 7]
排序前： [99 45 23 45 23]
排序后： [23 23 45 45 99]
--- PASS: TestBubble (0.00s)
*/

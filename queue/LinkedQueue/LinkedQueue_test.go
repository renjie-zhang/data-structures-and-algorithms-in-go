/*
 * @Descripttion:
 * @version: linkedqueue 测试
 * @Author: renjie.zhang
 * @Date: 2019-08-22 18:32:16
 * @LastEditTime: 2019-08-22 19:34:44
 */
package LinkedQueue

import (
	"fmt"
	"testing"
)

func TestLinkedQueue(t *testing.T) {
	N1 := Node{data: 1}
	Q := Queue{rear: &N1}
	Q.EnQueue(2)
	var el []int
	el = Q.Get()
	fmt.Println(el)
}

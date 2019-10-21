/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-09-04 21:35:55
 * @LastEditTime: 2019-10-21 23:00:44
 */
package ArrayStack

import (
	"errors"
	"fmt"
)

type Stack struct {
	Top    int           //表示栈顶 默认值是-1
	MaxTop int           //表示栈最大存放的数量
	data   []interface{} //栈的最大容量
}

//切片方式实现栈
func NewArrayStack(max int) *Stack {
	return &Stack{0, max, make([]interface{}, max)}
}

//判断栈是否已满
func (this *Stack) IsFull() bool {
	if this.Top == this.MaxTop-1 {
		return true
	}
	return false
}

//判断栈是否为空
func (this *Stack) IsEmpty() bool {
	if this.Top == -1 {
		return true
	}
	return false
}

//入栈
func (this *Stack) Push(value int) (err error) {
	if this.IsFull() {
		return errors.New("error : stack is full")
	}
	this.Top++
	this.data[this.Top] = value
	return
}

//遍历栈
func (this *Stack) Show() (err error) {
	if this.IsEmpty() {
		return errors.New("error : stack is empty")
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("data[%d]=%d", i, this.data[i])
	}
	return
}

//出栈
func (this *Stack) Pull() (value interface{}, err error) {
	if this.IsEmpty() {
		return 0, errors.New("error:stack is empty，can't pull")
	}
	temp := this.data[this.Top]
	this.Top--
	return temp, nil
}

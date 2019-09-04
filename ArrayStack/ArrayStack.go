package ArrayStack

import (
	"errors"
	"fmt"
)

type Stack struct {
	Top    int     //表示栈顶 默认值是-1
	MaxTop int     //表示栈最大存放的数量
	data   [10]int //栈的最大容量
}

//入栈
func (this *Stack) Push(value int) (err error) {
	if this.Top == this.MaxTop-1 {
		return errors.New("error : stack is full")
	}
	this.Top++
	this.data[this.Top] = value
	return
}

//遍历栈
func (this *Stack) Show() (err error) {
	if this.Top == -1 {
		return errors.New("error : stack is empty")
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("data[%d]=%d", i, this.data[i])
	}
	return
}

//出栈
func (this *Stack) Pull() (value int, err error) {
	if this.Top == -1 {
		return 0, errors.New("error:stack is empty，can't pull")
	}
	temp := this.data[this.Top]
	this.Top--
	return temp, nil
}

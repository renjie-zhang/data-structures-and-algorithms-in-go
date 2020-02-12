package FibonacciFind

//TODO 修改问题
func FinbonacciFind(array []int, findValue int) int {

	low := 0
	high := len(array) - 1
	k := 0
	mid := 0
	f := fib()
	for high > f[k]-1 {
		k++
	}
	// to do 使用切片，此处测试未通过
	temp := copy(array, f)
	for i := high + 1; i < len(temp); i++ {
		temp[i] = array[high]
	}

	for low < high {
		mid = low + f[k-1] - 1
		if findValue < temp[mid] {
			high = mid - 1
			k--
		} else if findValue > temp[mid] {
			low = mid + 1
			k -= 2
		} else {
			if mid < high {
				return mid
			} else {
				return high
			}
		}
	}
	return -1
}

//生成Fibonacci数列
func fib() []int {
	var f [20]int
	f[0] = 1
	f[1] = 1
	for i := 2; i < 20; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}

/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-11-24 17:14:49
 * @LastEditTime: 2019-11-24 21:21:54
 */
package PrimeNumber

import (
	"fmt"
	"testing"
)

func TestPrimeNumber(t *testing.T) {
	/*input := bufio.NewScanner(os.Stdin)
	var number int
	var temp int
	fmt.Println("请输入一个整数：")
	for input.Scan() {
		temp2 := input.Bytes()
		fmt.Println(temp2)
	}*/
	fmt.Println("100以内的素数有：")
	fmt.Printf("%d", 2)
	for j := 3; j <= 100; j += 2 {
		if Prime(j) == 1 {
			fmt.Printf("\t%d", j)
		}
	}
	fmt.Println()
}

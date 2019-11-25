/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-11-25 20:36:35
 * @LastEditTime: 2019-11-25 20:51:47
 */
package PalindromeNumber

import "strings"

func IsPalindromeNumber(s string) bool {
	if len(s) <= 2 {
		return false
	}
	temp := strings.Split(s, "")
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		if temp[i] != temp[j] {
			return false
		}
	}
	return true
}

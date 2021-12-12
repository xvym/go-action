/**
 * @author: xv
 * @date: 2021/12/12 17:14
 * @link: https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
 * @tag: 动态规划
 */

package swordoffer

import (
	"fmt"
	"testing"
)

func TestSo010(t *testing.T) {
	fmt.Println(fib(8))
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	temp1 := 0
	temp2 := 1
	res := 0
	for i := 2; i <= n; i++ {
		res = (temp1 + temp2) % 1000000007
		temp1 = temp2
		temp2 = res
	}
	return res
}

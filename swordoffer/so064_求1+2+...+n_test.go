/**
 * @author: xv
 * @date: 2021/12/15 22:31
 */

package swordoffer

import (
	"fmt"
	"testing"
)

func TestSo(t *testing.T) {
	fmt.Println(sumNums(9))
}

func sumNums(n int) int {
	return doSum(0, n, 0)
}

func doSum(num int, target int, res int) int {
	res += num
	if num == target {
		return res
	} else {
		return doSum(num+1, target, res)
	}
}

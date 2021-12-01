/**
 * @author: xv
 * @date: 2021/12/1 17:42
 * @link: https://leetcode-cn.com/problems/number-of-1-bits/
 * @tag: 位运算
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc191(t *testing.T) {
	fmt.Println(hammingWeight(35))
}

func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		num &= num - 1
		res++
	}
	return res
}
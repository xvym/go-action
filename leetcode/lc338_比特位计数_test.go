/**
 * @author: xv
 * @date: 2021/12/6 14:33
 * @link: https://leetcode-cn.com/problems/counting-bits/
 * @tag: 位运算
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc338(t *testing.T) {
	fmt.Println(countBits(5))
}

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i/2] + i%2
	}
	return res
}

//func countBits(n int) []int {
// res := make([]int, 0)
// for i := 0; i <= n; i++ {
//  count := countBitsOfNum(i)
//  res = append(res, count)
// }
// return res
//}
//
//func countBitsOfNum(num int) int {
// res := 0
// for num != 0 {
//  num &= num - 1
//  res++
// }
// return res
//}
/**
 * @author: xv
 * @date: 2021/12/1 10:14
 * @link: https://leetcode-cn.com/problems/climbing-stairs/
 * @tag: 动态规划
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc070(t *testing.T) {
	fmt.Println(climbStairs(1))
}

///**
// * 常规dp
// */
//func climbStairs(n int) int {
// if n <= 2 {
//  return n
// }
// dp := make([]int, n)
// dp[1] = 1
// dp[2] = 2
// for i := 2; i < n; i++ {
//  dp[i] = dp[i-1] + dp[i-2]
// }
// return dp[n-1]
//}

/**
 * dp状态压缩
 */
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	res := 0
	temp1 := 1
	temp2 := 2
	for i := 2; i < n; i++ {
		res = temp1 + temp2
		temp1 = temp2
		temp2 = res
	}
	return res
}
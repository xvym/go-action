/**
 * @author: xv
 * @date: 2021/12/1 13:00
 * @link: https://leetcode-cn.com/problems/triangle/
 * @tag: 动态规划
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc120(t *testing.T) {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}

///**
// * 常规dp
// */
//func minimumTotal(triangle [][]int) int {
// dp := make([][]int, len(triangle))
// for i := range dp {
//  dp[i] = make([]int, i+1)
// }
// res := 2147483647
// dp[0][0] = triangle[0][0]
// for i := 1; i < len(triangle); i++ {
//  for j := 0; j <= i; j++ {
//   if j != 0 && j != i {
//    dp[i][j] = triangle[i][j] + minInt(dp[i-1][j-1], dp[i-1][j])
//   } else if j == 0 {
//    dp[i][j] = triangle[i][j] + dp[i-1][0]
//   } else if j == i {
//    dp[i][j] = triangle[i][j] + dp[i-1][j-1]
//   }
//  }
// }
// for i := range dp[len(triangle)-1] {
//  res = minInt(res, dp[len(triangle)-1][i])
// }
// return res
//}

/**
 * 状态压缩，O(n)空间复杂度
 */
func minimumTotal(triangle [][]int) int {
	return 0
}

func minInt(var1 int, var2 int) int {
	if var1 < var2 {
		return var1
	} else {
		return var2
	}
}
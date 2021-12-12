/**
 * @author: xv
 * @date: 2021/12/12 18:42
 * @link: https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/
 * @tag: 动态规划
 */

package swordoffer

import (
	"testing"
)

func TestSo047(t *testing.T) {
}

func maxValue(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				if dp[i-1][j] < dp[i][j-1] {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				} else {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				}
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

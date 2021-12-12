/**
 * @author: xv
 * @date: 2021/12/12 18:09
 * @link: https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
 * @tag: 动态规划
 */

package swordoffer

import (
	"fmt"
	"testing"
)

func TestSo042(t *testing.T) {
	fmt.Println()
}

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	// dp数组含义: 以nums[i]为结尾的子数组的最大和
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}
	res := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
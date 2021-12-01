/**
 * @author: xv
 * @date: 2021/12/1 10:46
 * @link: https://leetcode-cn.com/problems/house-robber/
 * @tag: 动态规划
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc198(t *testing.T) {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

/**
 * 常规dp
 */
func rob(nums []int) int {
	length := len(nums)
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, length)
	dp[0] = nums[0]
	dp[1] = maxInt(nums[0], nums[1])
	for i := 2; i < length; i++ {
		dp[i] = maxInt(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[length-1]
}
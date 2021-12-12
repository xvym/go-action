/**
 * @author: xv
 * @date: 2021/12/12 22:44
 * @link: https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
 * @tag: 动态规划
 */

package swordoffer

import (
	"fmt"
	"testing"
)

func TestSo046(t *testing.T) {
	fmt.Println(translateNum(12258))
}

func translateNum(num int) int {
	nums := make([]int, 0)
	for num != 0 {
		temp := num % 10
		nums = append([]int{temp}, nums...)
		num /= 10
	}
	if len(nums) <= 1 {
		return 1
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	if (nums[1]+nums[0]*10) >= 10 && (nums[1]+nums[0]*10) <= 25 {
		dp[1] = 2
	} else {
		dp[1] = 1
	}
	for i := 2; i < len(nums); i++ {
		if (nums[i]+nums[i-1]*10) >= 10 && (nums[i]+nums[i-1]*10) <= 25 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(nums)-1]
}

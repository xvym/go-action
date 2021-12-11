/**
 * @author: xv
 * @date: 2021/12/8 22:27
 * @link: https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
 * @tag: 数组、二分法
 */

package swordoffer

import (
	"fmt"
	"testing"
)

func TestSo053II(t *testing.T) {
	fmt.Println(missingNumber([]int{0, 1, 3}))
}

func missingNumber(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (end-start)/2 + start
		if nums[mid] > mid {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}

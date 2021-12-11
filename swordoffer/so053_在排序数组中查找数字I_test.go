/**
 * @author: xv
 * @date: 2021/12/8 21:57
 * @link: https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
 * @tag: 二分法、数组
 */

package swordoffer

import (
	"fmt"
	"testing"
)

func TestSo053(t *testing.T) {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
}

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	// 找结束位置
	for start <= end {
		mid := (end-start)/2 + start
		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	right := start
	start = 0
	end = len(nums) - 1
	for start <= end {
		mid := (end-start)/2 + start
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	left := start
	return right - left
}

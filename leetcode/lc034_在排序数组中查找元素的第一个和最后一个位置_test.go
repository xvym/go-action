/**
 * @author: xv
 * @date: 2021/12/5 15:56
 * @link: https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
 * @tag: 二分法
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc034(t *testing.T) {
	fmt.Println(searchRange([]int{1}, 1))
}

func searchRange(nums []int, target int) []int {
	start := 0
	end := len(nums) - 1
	res := make([]int, 2)
	res[0] = -1
	res[1] = -1
	mid := -1
	// 找开始位置
	for start <= end {
		mid = (end-start)/2 + start
		if nums[mid] >= target {
			if nums[mid] == target && (mid == 0 || nums[mid-1] != target) {
				res[0] = mid
				break
			}
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		}
	}

	start = 0
	end = len(nums) - 1
	// 找结束位置
	for start <= end {
		mid = (end-start)/2 + start
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] <= target {
			if nums[mid] == target && (mid == len(nums) - 1 || nums[mid+1] != target) {
				res[1] = mid
				break
			}
			start = mid + 1
		}
	}
	return res
}

/**
 * @author: xv
 * @date: 2021/11/17 17:22
 * @link: https://leetcode-cn.com/problems/binary-search/
 * @tag: 二分法
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc704(t *testing.T) {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func search(nums []int, target int) int {
	var (
		start int
		end   int
	)

	start = 0
	end = len(nums) - 1

	for start <= end {
		var mid int
		mid = start + (end-start)/2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

/**
 * @author: xv
 * @date: 2021/11/18 0:09
 * @link: https://leetcode-cn.com/problems/search-insert-position/
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc035(t *testing.T) {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (end-start)/2 + start
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] == target {
			return mid
		}
	}
	return start
}

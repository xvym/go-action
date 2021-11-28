/**
 * @author: xv
 * @date: 2021/11/18 11:32
 * @link: https://leetcode-cn.com/problems/rotate-array/
 * @tag: 双指针
 */
package leetcode

import (
	"testing"
)

func TestLc189(t *testing.T) {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

func rotate(nums []int, k int) {
	length := len(nums)
	k %= length
	reserveNums(nums, 0, length-1)
	reserveNums(nums, 0, k - 1)
	reserveNums(nums, k, length-1)
}

func reserveNums(nums []int, start int, end int) {
	for start < end {
		tmp := nums[end]
		nums[end] = nums[start]
		nums[start] = tmp
		start++
		end--
	}
}

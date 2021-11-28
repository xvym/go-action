/**
 * @author: xv
 * @date: 2021/11/18 10:14
 * @link: https://leetcode-cn.com/problems/squares-of-a-sorted-array/submissions/
 * @tag: 双指针
 */

package leetcode

import "testing"

func TestLc977(t *testing.T) {
	sortedSquares([]int{})
}

func sortedSquares(nums []int) []int {
	var length = len(nums)
	var res = make([]int, length)
	i := length - 1
	left := 0
	right := length - 1
	for left <= right {
		leftVal := nums[left] * nums[left]
		rightVal := nums[right] * nums[right]
		if leftVal <= rightVal {
			res[i] = rightVal
			right--
		} else {
			res[i] = leftVal
			left++
		}
		i--
	}
	return res
}

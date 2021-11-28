/**
 * @author: xv
 * @date: 2021/11/24 14:49
 * @link: https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
 * @tag: 双指针
 */

package leetcode

import "testing"

func TestLc167(t *testing.T) {
	twoSum([]int{2, 7, 11, 5}, 9)
}

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	res := make([]int, 2)
	for left < right {
		if numbers[left]+numbers[right] < target {
			left++
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			res[0] = left + 1
			res[1] = right + 1
			break
		}
	}
	return res
}
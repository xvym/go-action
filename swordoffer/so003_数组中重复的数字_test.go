/**
 * @author: xv
 * @date: 2021/12/8 21:56
 * @link: https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
 * @tag: 数组
 */

package swordoffer

import "testing"

func TestSo003(t *testing.T) {
	findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3})
}

func findRepeatNumber(nums []int) int {
	length := len(nums)
	res := -1
	for i := range nums {
		index := nums[i]
		if index >= length {
			index -= length
		}
		if nums[index] >= length {
			res = index
			break
		} else {
			nums[index] += length
		}
	}
	return res
}

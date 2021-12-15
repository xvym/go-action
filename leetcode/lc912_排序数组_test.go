/**
 * @author: xv
 * @date: 2021/12/15 22:57
 * @link: https://leetcode-cn.com/problems/sort-an-array/
 * @tag: 排序
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc912(t *testing.T) {
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start int, end int) {
	if start < end {
		pos := partition(nums, start, end)
		quickSort(nums, start, pos-1)
		quickSort(nums, pos+1, end)
	}
}

func partition(nums []int, start int, end int) int {
	i := start
	j := end
	pivot := nums[start]
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[start] = nums[start], nums[i]
	return i
}

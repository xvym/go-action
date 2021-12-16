/**
 * @author: xv
 * @date: 2021/12/15 22:57
 * @link: https://leetcode-cn.com/problems/sort-an-array/
 * @tag: 排序
 */

package leetcode

import (
	"fmt"
	"math/rand"
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

func partition(nums []int, left int, right int) int {
	pivotIndex := rand.Intn(right-left+1) + left
	nums[pivotIndex], nums[left] = nums[left], nums[pivotIndex]
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}

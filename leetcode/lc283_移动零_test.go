/**
 * @author: xv
 * @date: 2021/11/24 14:33
 * @link: https://leetcode-cn.com/problems/move-zeroes/
 * @tag: 双指针
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc283(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	for i := range nums {
		fmt.Println(nums[i])
	}

}

func moveZeroes(nums []int) {
	curIndex := 0
	for i := range nums {
		if nums[i] != 0 {
			if i != curIndex {
				nums[curIndex] = nums[i]
				nums[i] = 0
			}
			curIndex++
		}
	}
}
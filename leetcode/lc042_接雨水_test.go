/**
 * @author: xv
 * @date: 2021/12/2 17:10
 * @link: https://leetcode-cn.com/problems/trapping-rain-water/
 * @tag: 单调栈、动态规划、双指针
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc042(t *testing.T) {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}

func trap(height []int) int {
	stack := make([]int, 0)
	stack = append(stack, 0)
	res := 0
	for i := range height {
		cur := 0
		for len(stack) > 0 && height[i] >= height[stack[len(stack)-1]] {
			topIndex := stack[len(stack)-1]
			stack = append(stack[:len(stack)-1], stack[len(stack)-1+1:]...)
			if len(stack) > 0 {
				secondTopIndex := stack[len(stack)-1]
				var high int
				if height[i] > height[secondTopIndex] {
					high = height[secondTopIndex]
				} else {
					high = height[i]
				}
				cur += (high - height[topIndex]) * (i - secondTopIndex - 1)
			}
		}
		res += cur
		stack = append(stack, i)
	}
	return res
}

/**
 * @author: xv
 * @date: 2021/11/24 16:56
 * @link: https://leetcode-cn.com/problems/max-chunks-to-make-sorted/
 * @tag: 单调栈
 */

package leetcode

import (
	"fmt"
	"testing"
)

func TestLc769(t *testing.T) {
	fmt.Println(maxChunksToSorted([]int{4, 3, 2, 1, 0}))
}

func maxChunksToSorted(arr []int) int {
	res := 0
	max := 0
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
		if max == i {
			res++
		}
	}
	return res
}
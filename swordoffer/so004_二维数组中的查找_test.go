/**
 * @author: xv
 * @date: 2021/12/8 22:42
 * @link: https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
 * @tag: 数组
 */

package swordoffer

import (
	"fmt"
	"testing"
)

func TestSo004(t *testing.T) {
	fmt.Println(findNumberIn2DArray([][]int{{5}}, -10))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) <= 0 {
		return false
	} else if len(matrix) <= 0 {
		return false
	}

	for i := len(matrix) - 1; i >= 0; i-- {
		for j := 0; j < len(matrix[0]); j++ {
			for i >= 0 && matrix[i][j] > target {
				i--
			}
			if i < 0 {
				return false
			}
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

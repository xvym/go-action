/**
 * @author: xv
 * @date: 2021/11/30 16:22
 * @link: https://leetcode-cn.com/problems/combinations/
 * @tag: 回溯
 */

package leetcode

import (
	"testing"
)

func TestLc077(t *testing.T) {
	combine(4, 2)
}

func combine(n int, k int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	return backtrackIn077(n, 1, k, res, path)
}

func backtrackIn077(n int, num int, k int, res [][]int, path []int) [][]int {
	if len(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		res = append(res, temp)
	} else {
		for i := num; i <= n; i++ {
			path = append(path, i)
			res = backtrackIn077(n, i+1, k, res, path)
			path = append(path[0 : len(path)-1])
		}
	}
	return res
}

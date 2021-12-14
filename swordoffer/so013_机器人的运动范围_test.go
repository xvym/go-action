/**
 * @author: xv
 * @date: 2021/12/15 0:21
 * @link: https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
 * @tag: dfs
 */

package swordoffer

import (
	"fmt"
	"testing"
)

func TestSo013(t *testing.T) {
	fmt.Println(movingCount(3, 1, 0))
}

func movingCount(m int, n int, k int) int {
	res := 0
	used := make([][]int, m)
	for i := range used {
		used[i] = make([]int, n)
	}
	res = backtrackInSo013(0, 0, m, n, k, res, used)
	return res
}

func backtrackInSo013(i int, j int, m int, n int, k int, res int, used [][]int) int {
	if i >= 0 && i < m && j >= 0 && j < n && getSum(i)+getSum(j) <= k && used[i][j] != 1 {
		res++
		used[i][j] = 1
		res = backtrackInSo013(i+1, j, m, n, k, res, used)
		res = backtrackInSo013(i-1, j, m, n, k, res, used)
		res = backtrackInSo013(i, j+1, m, n, k, res, used)
		res = backtrackInSo013(i, j-1, m, n, k, res, used)
	}
	return res
}

func getSum(num int) int {
	sum := 0
	for num != 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

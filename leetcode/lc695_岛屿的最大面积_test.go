/**
 * @author: xv
 * @date: 2021/11/29 11:04
 * @link: https://leetcode-cn.com/problems/max-area-of-island/
 * @tag: dfs
 */

package leetcode

import "testing"

var tmp int

func TestLc695(t *testing.T) {
	maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}})
}

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				countArea(grid, i, j)
				if tmp > res {
					res = tmp
				}
				tmp = 0
			}
		}
	}
	return res
}

func countArea(grid [][]int, i int, j int) {
	if i < len(grid) && j < len(grid[0]) && i >= 0 && j >= 0 && grid[i][j] == 1 {
		tmp++
		grid[i][j] = 0
		countArea(grid, i+1, j)
		countArea(grid, i-1, j)
		countArea(grid, i, j+1)
		countArea(grid, i, j-1)
	}
}
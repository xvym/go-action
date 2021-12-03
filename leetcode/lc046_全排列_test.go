/**
 * @author: xv
 * @date: 2021/11/30 23:41
 * @link: https://leetcode-cn.com/problems/permutations/
 * @tag: 回溯
 */

package leetcode

import "testing"

func TestLc046(t *testing.T) {

}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, len(nums))
	visit := make([]int, len(nums))
	for i := range visit {
		visit[i] = 0
	}
	res = backtrackIn046(0, res, path, nums, visit)
	return res
}

func backtrackIn046(index int, res [][]int, path []int, nums []int, visit []int) [][]int {
	if index == len(nums) {
		temp := make([]int, index)
		copy(temp, path)
		res = append(res, temp)
	}
	for i := 0; i < len(nums); i++ {
		if visit[i] != 1 {
			visit[i] = 1
			path[index] = nums[i]
			backtrackIn046(index + 1, res, path, nums, visit)
			visit[i] = 0
		}
	}
	return res
}

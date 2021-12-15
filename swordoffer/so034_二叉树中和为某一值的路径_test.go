/**
 * @author: xv
 * @date: 2021/12/15 11:08
 * @link: https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
 * @tag: dfs
 */

package swordoffer

import (
	"fmt"
	. "go-action/common"
	"testing"
)

func TestSo034(t *testing.T) {
	root := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				4,
				nil,
				nil,
			},
			nil},
		&TreeNode{3,
			//&TreeNode{15, nil, nil},
			nil,
			&TreeNode{5, nil, nil},
		},
	}
	fmt.Println(pathSum(root, 7))
}

func pathSum(root *TreeNode, target int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	return backtrackInSo034(root, target, 0, path, res)
}

func backtrackInSo034(root *TreeNode, target int, curSum int, path []int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	path = append(path, root.Val)
	if root.Val+curSum == target {
		if root.Left == nil && root.Right == nil {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
	}
	res = backtrackInSo034(root.Left, target, curSum+root.Val, path, res)
	res = backtrackInSo034(root.Right, target, curSum+root.Val, path, res)
	return res
}
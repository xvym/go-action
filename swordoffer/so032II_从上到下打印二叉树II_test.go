/**
 * @author: xv
 * @date: 2021/12/11 11:29
 * @link: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
 * @tag: 二叉树、bfs
 */

package swordoffer

import (
	"fmt"
	. "go-action/common"
	"testing"
)

func TestSo032II(t *testing.T) {
	root := &TreeNode{
		3,
		&TreeNode{
			9,
			nil,
			nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil},
		},
	}
	fmt.Println(levelOrderII(root))
}

func levelOrderII(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	res := make([][]int, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		temp := make([]int, 0)
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = append(queue[:0], queue[1:]...)
		}
		res = append(res, temp)
	}
	return res
}

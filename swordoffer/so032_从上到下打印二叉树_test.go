/**
 * @author: xv
 * @date: 2021/12/8 23:48
 * @link: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
 * @tag: bfs、二叉树
 */

package swordoffer

import (
	"fmt"
	. "go-action/common"
	"testing"
)

func TestSo032(t *testing.T) {
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
	fmt.Println(levelOrder(root))
}

func levelOrder(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	res := make([]int, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		node := queue[0]
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		queue = append(queue[:0], queue[1:]...)
	}
	return res
}

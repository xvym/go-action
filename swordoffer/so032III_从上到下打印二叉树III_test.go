/**
 * @author: xv
 * @date: 2021/12/11 11:51
 * @link: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
 * @tag: 二叉树、bfs
 */

package swordoffer

import (
	"fmt"
	. "go-action/common"
	"testing"
)

func TestSo032III(t *testing.T) {
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
	fmt.Println(levelOrderIII(root))
}

func levelOrderIII(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	res := make([][]int, 0)
	if root != nil {
		queue = append(queue, root)
	}
	direct := 0
	for len(queue) > 0 {
		temp := make([]int, 0)
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			if direct == 0 {
				// 正序追加
				temp = append(temp, node.Val)
			} else {
				// 逆序追加
				temp = append([]int{node.Val}, temp...)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = append(queue[:0], queue[1:]...)
		}
		if direct == 0 {
			direct = 1
		} else {
			direct = 0
		}
		res = append(res, temp)
	}
	return res
}

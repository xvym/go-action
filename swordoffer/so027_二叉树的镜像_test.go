/**
 * @author: xv
 * @date: 2021/12/12 15:55
 * @link: https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
 * @tag: 二叉树、递归
 */

package swordoffer

import (
	. "go-action/common"
	"testing"
)

func TestSo027(t *testing.T) {

}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else {
		temp := root.Left
		root.Left = root.Right
		root.Right = temp
		mirrorTree(root.Left)
		mirrorTree(root.Right)
		return root
	}
}

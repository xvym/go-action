/**
 * @author: xv
 * @date: 2021/12/12 16:27
 * @link: https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
 * @tag: 二叉树、递归
 */

package swordoffer

import (
	. "go-action/common"
	"testing"
)

func TestSo028(t *testing.T) {

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		return doIsSymmetric(root.Left, root.Right)
	}
}

func doIsSymmetric(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) || node1.Val != node2.Val {
		return false
	} else {
		return doIsSymmetric(node1.Left, node2.Right) && doIsSymmetric(node1.Right, node2.Left)
	}
}

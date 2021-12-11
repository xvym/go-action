/**
 * @author: xv
 * @date: 2021/12/11 17:47
 * @link: https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
 * @tag: 树、递归
 */

package swordoffer

import (
	. "go-action/common"
	"testing"
)

func TestSo026(t *testing.T) {

}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	} else {
		return doIsSubStructure(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
	}
}

func doIsSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	} else if A == nil || A.Val != B.Val {
		return false
	} else {
		return doIsSubStructure(A.Left, B.Left) && doIsSubStructure(A.Right, B.Right)
	}
}

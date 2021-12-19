/**
 * @author: xv
 * @date: 2021/12/18 0:56
 */

package leetcode

import (
	. "go-action/common"
	"testing"
)

func TestLc235(t *testing.T) {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return doLowestCommonAncestor(root, p, q, nil)
}

func doLowestCommonAncestor(root, p, q, res *TreeNode) *TreeNode {
	if root != nil {
		res = lowestCommonAncestor(root.Left, p, q)
		if res == nil {
			res = lowestCommonAncestor(root.Right, p, q)
		}
		if res == nil {
			if doFindTarget(root, p) && doFindTarget(root, q) {
				res = root
			}
		}
	}
	return res
}

func doFindTarget(root, target *TreeNode) bool {
	if root == nil {
		return false
	} else if root == target {
		return true
	} else {
		return doFindTarget(root.Left, target) || doFindTarget(root.Right, target)
	}
}

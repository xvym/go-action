/**
 * @author: xv
 * @date: 2021/11/29 22:10
 * @link: https://leetcode-cn.com/problems/merge-two-binary-trees/
 * @tag: 二叉树、dfs
 */

package leetcode

import "testing"

func TestLc617(t *testing.T) {

}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 != nil && root2 != nil {
		root1.Val += root2.Val
		left := mergeTrees(root1.Left, root2.Left)
		right := mergeTrees(root1.Right, root2.Right)
		root1.Left = left
		root1.Right = right
		return root1
	} else if root1 != nil {
		return root1
	} else if root2 != nil {
		return root2
	}
	return nil
}

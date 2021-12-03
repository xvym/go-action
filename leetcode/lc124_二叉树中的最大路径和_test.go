/**
 * @author: xv
 * @date: 2021/12/3 10:46
 * @link: https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
 * @tag: dfs、递归、二叉树
 */

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func main() {
	maxPathSum(nil)
}
//func TestLc124(t *testing.T) {
//
//}

func maxPathSum(root *TreeNode) int {
	var res int
	_, res = backtrackInLc124(root, -2147483648)
	return res
}

func backtrackInLc124(root *TreeNode, res int) (int, int) {
	curVal := root.Val
	curPathVal := root.Val
	leftPathVal := 0
	rightPathVal := 0
	if root.Left != nil {
		leftPathVal, res = backtrackInLc124(root.Left, res)
	}
	if root.Right != nil {
		rightPathVal, res = backtrackInLc124(root.Right, res)
	}
	if leftPathVal > 0 {
		curVal += leftPathVal
	}
	if rightPathVal > 0 {
		curVal += rightPathVal
	}

	if curVal > res {
		res = curVal
	}

	if leftPathVal > 0 || rightPathVal > 0 {
		if leftPathVal >= rightPathVal {
			curPathVal += leftPathVal
		} else {
			curPathVal += rightPathVal
		}
	}
	return curPathVal, res
}
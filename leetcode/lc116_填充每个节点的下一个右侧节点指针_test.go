/**
 * @author: xv
 * @date: 2021/11/29 22:09
 * @link: https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/submissions/
 * @tag: 二叉树、dfs
 */

package leetcode

import "testing"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
func TestLc116(t *testing.T) {

}

func connect(root *Node) *Node {
	if root != nil {
		if root.Left != nil {
			root.Left.Next = root.Right
			if root.Next != nil {
				root.Right.Next = root.Next.Left
			}
			connect(root.Left)
			connect(root.Right)
		}
	}
	return root
}

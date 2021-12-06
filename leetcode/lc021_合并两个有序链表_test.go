/**
 * @author: xv
 * @date: 2021/12/2 10:55
 * @link: https://leetcode-cn.com/problems/merge-two-sorted-lists/
 * @tag: 链表
 */

package leetcode

import (
	. "go-action/common"
	"testing"
)

func TestLc021(t *testing.T) {
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}
	head2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	mergeTwoLists(head1, head2)
}

/**
 * 迭代解法
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	for list1 != nil {
		cur.Next = list1
		cur = cur.Next
		list1 = list1.Next
	}
	for list2 != nil {
		cur.Next = list2
		cur = cur.Next
		list2 = list2.Next
	}
	return dummy.Next
}

/**
 * 递归解法
 */

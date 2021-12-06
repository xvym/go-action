/**
 * @author: xv
 * @date: 2021/12/2 9:53
 * @link: https://leetcode-cn.com/problems/merge-k-sorted-lists/
 * @tag: 链表
 */

package leetcode

import (
	. "go-action/common"
	"testing"
)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func TestLc023(t *testing.T) {
	head1 := ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}
	head2 := ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	head3 := ListNode{Val: 2, Next: &ListNode{Val: 6, Next: nil}}
	lists := make([]*ListNode, 3)
	lists[0] = &head1
	lists[1] = &head2
	lists[2] = &head3
	mergeKLists(lists)
}

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	for i := 0; i < len(lists); i++ {
		head = mergeTwoListsInLc023(head, lists[i])
	}
	return head
}

func mergeTwoListsInLc023(list1 *ListNode, list2 *ListNode) *ListNode {
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
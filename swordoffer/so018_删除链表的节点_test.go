/**
 * @author: xv
 * @date: 2021/12/13 23:38
 */

package swordoffer

import (
	. "go-action/common"
	"testing"
)

func TestSo018(t *testing.T) {

}

func deleteNode(head *ListNode, val int) *ListNode {
	cur := head
	dummy := &ListNode{-1, head}
	dummy.Next = head
	pre := dummy
	for cur != nil {
		if cur.Val == val {
			if cur.Next != nil {
				cur.Val = cur.Next.Val
				cur.Next = cur.Next.Next
			} else {
				pre.Next = nil
			}
		}
		cur = cur.Next
		pre = pre.Next
	}
	return head
}

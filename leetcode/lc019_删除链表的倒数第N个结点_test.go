/**
 * @author: xv
 * @date: 2021/11/24 15:07
 * @link: https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
 * @tag: 链表、双指针、快慢指针
 */

package leetcode

import (
	. "go-action/common"
	"testing"
)

func TestLc019(t *testing.T) {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	dummy := &ListNode{-1, slow}
	pre := dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
		pre = pre.Next
	}
	pre.Next = slow.Next
	return dummy.Next
}

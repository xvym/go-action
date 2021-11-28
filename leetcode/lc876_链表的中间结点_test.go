/**
 * @author: xv
 * @date: 2021/11/24 14:58
 * @link: https://leetcode-cn.com/problems/middle-of-the-linked-list/
 * @tag: 链表、双指针、快慢指针
 */

package leetcode

import "testing"

func TestLc876(t *testing.T) {

}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
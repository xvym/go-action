/**
 * @author: xv
 * @date: 2021/12/6 11:36
 * @link: https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
 * @tag: 链表
 */

package swordoffer

import (
	. "go-action/common"
	"testing"
)

func TestSo024(t *testing.T) {
	head := BuildDefaultListNode(5)
	reverseList(head)
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
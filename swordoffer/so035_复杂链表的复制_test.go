/**
 * @author: xv
 * @date: 2021/12/6 14:12
 * @link: https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/
 * @tag: 链表
 */

package swordoffer

import (
	. "go-action/common"
	"testing"
)

func TestSo035(t *testing.T) {
	head := BuildDefaultNode(5)
	copyRandomList(head)
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	myMap := make(map[*Node]*Node)
	newHead := &Node{Val: head.Val}
	cur := head
	newCur := newHead
	myMap[cur] = newCur
	for cur.Next != nil {
		next := &Node{Val: cur.Next.Val}
		newCur.Next = next
		cur = cur.Next
		newCur = newCur.Next
		myMap[cur] = next
	}
	cur = head
	newCur = newHead
	for cur != nil {
		if cur.Random != nil {
			newCur.Random = myMap[cur.Random]
		}
		cur = cur.Next
		newCur = newCur.Next
	}
	return newHead
}